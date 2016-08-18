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
	"net/http"

	"gopkg.in/spacemonkeygo/monkit.v2"
)

var (
	httpclient = monkit.ScopeNamed("http.client")
	httpserver = monkit.ScopeNamed("http.server")
)

// Client is an interface that matches an http.Client
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

func traceRequest(s *monkit.Span, cl Client, req *http.Request) (
	resp *http.Response, err error) {
	s.Annotate("http.uri", req.URL.String())
	RequestFromSpan(s).SetHeader(req.Header)
	resp, err = cl.Do(req)
	if err != nil {
		return resp, err
	}
	s.Annotate("http.responsecode", fmt.Sprint(resp.StatusCode))
	return resp, nil
}

type responseWriterObserver struct {
	w  http.ResponseWriter
	sc *int
}

func (w *responseWriterObserver) WriteHeader(status_code int) {
	w.sc = &status_code
	w.w.WriteHeader(status_code)
}

func (w *responseWriterObserver) Write(p []byte) (n int, err error) {
	if w.sc == nil {
		sc := 200
		w.sc = &sc
	}
	return w.w.Write(p)
}

func (w *responseWriterObserver) Header() http.Header {
	return w.w.Header()
}

func (w *responseWriterObserver) StatusCode() int {
	if w.sc == nil {
		return 200
	}
	return *w.sc
}

//go:generate sh -c "m4 -D_STDLIB_IMPORT_='\"context\"' -D_OTHER_IMPORT_= -D_BUILD_TAG_='// +build go1.7' httpctxgen.go.m4 > httpctx17.go"
//go:generate sh -c "m4 -D_STDLIB_IMPORT_= -D_OTHER_IMPORT_='\"golang.org/x/net/context\"' -D_BUILD_TAG_='// +build !go1.7' httpctxgen.go.m4 > httpxctx.go"
//go:generate gofmt -w -s httpctx17.go httpxctx.go
