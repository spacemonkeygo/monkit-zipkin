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
	"strconv"

	"gopkg.in/spacemonkeygo/monkit.v2"
)

// Request is a structure representing an incoming RPC request. Every field
// is optional.
type Request struct {
	TraceId  *int64
	SpanId   *int64
	ParentId *int64
	Sampled  *bool
	Flags    *int64
}

// HeaderGetter is an interface that http.Header matches for RequestFromHeader
type HeaderGetter interface {
	Get(string) string
}

// HeaderSetter is an interface that http.Header matches for Request.SetHeader
type HeaderSetter interface {
	Set(string, string)
}

// RequestFromHeader will create a Request object given an http.Header or
// anything that matches the HeaderGetter interface.
func RequestFromHeader(header HeaderGetter) (rv Request) {
	trace_id, err := fromHeader(header.Get("X-B3-TraceId"))
	if err == nil {
		rv.TraceId = &trace_id
	}
	span_id, err := fromHeader(header.Get("X-B3-SpanId"))
	if err == nil {
		rv.SpanId = &span_id
	}
	parent_id, err := fromHeader(header.Get("X-B3-ParentSpanId"))
	if err == nil {
		rv.ParentId = &parent_id
	}
	sampled, err := strconv.ParseBool(header.Get("X-B3-Sampled"))
	if err == nil {
		rv.Sampled = &sampled
	}
	flags, err := fromHeader(header.Get("X-B3-Flags"))
	if err == nil {
		rv.Flags = &flags
	}
	return rv
}

func ref(v int64) *int64 {
	return &v
}

func RequestFromSpan(s *monkit.Span) Request {
	trace := s.Trace()

	sampled, ok := trace.Get(sampleKey).(bool)
	if !ok {
		sampled = false
	}

	if !sampled {
		return Request{Sampled: &sampled}
	}
	flags, ok := trace.Get(flagsKey).(int64)
	if !ok {
		flags = 0
	}
	parent_id, _ := getParentId(s)
	return Request{
		TraceId:  ref(trace.Id()),
		SpanId:   ref(s.Id()),
		Sampled:  &sampled,
		Flags:    &flags,
		ParentId: parent_id,
	}
}

// SetHeader will take a Request and fill out an http.Header, or anything that
// matches the HeaderSetter interface.
func (r Request) SetHeader(header HeaderSetter) {
	if r.TraceId != nil {
		header.Set("X-B3-TraceId", toHeader(*r.TraceId))
	}
	if r.SpanId != nil {
		header.Set("X-B3-SpanId", toHeader(*r.SpanId))
	}
	if r.ParentId != nil {
		header.Set("X-B3-ParentSpanId", toHeader(*r.ParentId))
	}
	if r.Sampled != nil {
		header.Set("X-B3-Sampled", strconv.FormatBool(*r.Sampled))
	}
	if r.Flags != nil {
		header.Set("X-B3-Flags", toHeader(*r.Flags))
	}
}

func (zipreq Request) Trace() (trace *monkit.Trace, spanId int64) {
	if zipreq.TraceId != nil {
		trace = monkit.NewTrace(*zipreq.TraceId)
	} else {
		trace = monkit.NewTrace(monkit.NewId())
	}
	if zipreq.SpanId != nil {
		spanId = *zipreq.SpanId
	} else {
		spanId = monkit.NewId()
	}

	if zipreq.ParentId != nil {
		trace.Set(remoteParentKey, *zipreq.ParentId)
	}
	if zipreq.Sampled != nil {
		trace.Set(sampleKey, *zipreq.Sampled)
	}
	if zipreq.Flags != nil {
		trace.Set(flagsKey, *zipreq.Flags)
	}
	return trace, spanId
}

// fromHeader reads a signed int64 that has been formatted as a hex uint64
func fromHeader(s string) (int64, error) {
	v, err := strconv.ParseUint(s, 16, 64)
	return int64(v), err
}

// toHeader writes a signed int64 as hex uint64
func toHeader(i int64) string {
	return strconv.FormatUint(uint64(i), 16)
}
