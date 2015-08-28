// Copyright (C) 2014 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zipkin

import (
	"fmt"
	"time"

	"github.com/spacemonkeygo/errors"
	"gopkg.in/spacemonkeygo/monitor-zipkin.v2/gen-go/zipkin"
	"gopkg.in/spacemonkeygo/monitor.v2"
)

type traceKey int

const (
	sampleKey       traceKey = 0
	flagsKey        traceKey = 1
	remoteParentKey traceKey = 2
)

type Options struct {
	Fraction  float64
	Debug     bool
	LocalHost *zipkin.Endpoint

	collector TraceCollector
}

func RegisterZipkin(reg *monitor.Registry, collector TraceCollector,
	opts Options) {
	opts.collector = collector
	reg.ObserveTraces(func(t *monitor.Trace) {
		sampled, exists := t.Get(sampleKey).(bool)
		if !exists {
			sampled = rng.Float64() < opts.Fraction
			t.Set(sampleKey, sampled)
		}
		if sampled {
			flags, ok := t.Get(flagsKey).(int64)
			if !ok {
				flags = 0
			}
			if opts.Debug {
				flags = flags | 1
			}
			t.Set(flagsKey, flags)
			t.ObserveSpans(opts.observeSpan)
		}
	})
}

func getParentId(s *monitor.Span) (parent_id *int64, server bool) {
	parent := s.Parent()
	if parent != nil {
		parent_id := parent.Id()
		return &parent_id, false
	}
	if parent_id, ok := s.Trace().Get(remoteParentKey).(int64); ok {
		return &parent_id, true
	}
	return nil, false
}

func (opts Options) observeSpan(s *monitor.Span, err error, panicked bool,
	finish time.Time) {
	parent_id, server := getParentId(s)
	zs := &zipkin.Span{
		TraceId:  s.Trace().Id(),
		Name:     s.Func().FullName(),
		Id:       s.Id(),
		ParentId: parent_id,
	}

	start_name := zipkin.CLIENT_SEND
	end_name := zipkin.CLIENT_RECV
	if server {
		start_name = zipkin.SERVER_RECV
		end_name = zipkin.SERVER_SEND
	}

	annotations := s.Annotations()
	args := s.Args()
	zs.Annotations = make([]*zipkin.Annotation, 0,
		3+len(annotations)+len(args))

	zs.Annotations = append(zs.Annotations, &zipkin.Annotation{
		Timestamp: s.Start().UnixNano() / 1000,
		Value:     start_name,
		Host:      opts.LocalHost})
	for arg_idx, arg := range args {
		zs.BinaryAnnotations = []*zipkin.BinaryAnnotation{{
			Key:            fmt.Sprintf("arg_%d", arg_idx),
			Value:          []byte(arg),
			AnnotationType: zipkin.AnnotationType_STRING,
			Host:           opts.LocalHost}}
	}
	for _, annotation := range annotations {
		zs.BinaryAnnotations = []*zipkin.BinaryAnnotation{{
			Key:            annotation.Name,
			Value:          []byte(annotation.Value),
			AnnotationType: zipkin.AnnotationType_STRING,
			Host:           opts.LocalHost}}
	}
	if panicked {
		zs.Annotations = append(zs.Annotations, &zipkin.Annotation{
			Timestamp: finish.UnixNano() / 1000,
			Value:     "failed",
			Host:      opts.LocalHost})
		zs.Annotations = append(zs.Annotations, &zipkin.Annotation{
			Timestamp: finish.UnixNano() / 1000,
			Value:     "panic",
			Host:      opts.LocalHost})
	} else if err != nil {
		zs.Annotations = append(zs.Annotations, &zipkin.Annotation{
			Timestamp: finish.UnixNano() / 1000,
			Value:     "failed",
			Host:      opts.LocalHost})
		zs.BinaryAnnotations = []*zipkin.BinaryAnnotation{{
			Key:            "error",
			Value:          []byte(errors.GetClass(err).String()),
			AnnotationType: zipkin.AnnotationType_STRING,
			Host:           opts.LocalHost}}
	}
	zs.Annotations = append(zs.Annotations, &zipkin.Annotation{
		Timestamp: finish.UnixNano() / 1000,
		Value:     end_name,
		Host:      opts.LocalHost})
	opts.collector.Collect(zs)
}
