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

	"golang.org/x/net/context"
	"gopkg.in/spacemonkeygo/monitor.v2"
)

var (
	monhttp = monitor.PackageNamed("http")
)

// client stuff -----

// Client is an interface that matches an http.Client
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// TraceRequest will perform an HTTP request, creating a new Span for the HTTP
// request and sending the Span in the HTTP request headers.
// Compare to http.Client.Do.
func TraceRequest(ctx context.Context, cl Client, req *http.Request) (
	resp *http.Response, err error) {
	defer monhttp.TaskNamed(req.Method)(&ctx)(&err)
	s := monitor.SpanFromCtx(ctx)
	s.Annotate("http.uri", req.URL.String())
	RequestFromSpan(s).SetHeader(req.Header)
	resp, err = cl.Do(req)
	if err != nil {
		return resp, err
	}
	s.Annotate("http.responsecode", fmt.Sprint(resp.StatusCode))
	return resp, nil
}

// server stuff -----

// TraceHandler wraps a ContextHTTPHandler with a Span pulled from incoming
// requests, possibly starting new Traces if necessary.
func TraceHandler(c ContextHTTPHandler) ContextHTTPHandler {
	return ContextHTTPHandlerFunc(func(
		ctx context.Context, w http.ResponseWriter, r *http.Request) {
		trace, spanId := RequestFromHeader(r.Header).Trace()
		defer monhttp.FuncNamed(r.Method).RemoteTrace(&ctx, spanId, trace)(nil)
		s := monitor.SpanFromCtx(ctx)
		s.Annotate("http.uri", r.RequestURI)
		wrapped := &responseWriterObserver{w: w}
		c.ServeHTTP(ctx, wrapped, r)
		s.Annotate("http.responsecode", fmt.Sprint(wrapped.StatusCode()))
	})
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

// ContextHTTPHandler is like http.Handler, but expects a Context object
// as the first parameter.
type ContextHTTPHandler interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

// ContextHTTPHandlerFunc is like http.HandlerFunc but for ContextHTTPHandlers
type ContextHTTPHandlerFunc func(
	ctx context.Context, w http.ResponseWriter, r *http.Request)

func (f ContextHTTPHandlerFunc) ServeHTTP(ctx context.Context,
	w http.ResponseWriter, r *http.Request) {
	f(ctx, w, r)
}

// ContextWrapper will turn a ContextHTTPHandler into an http.Handler by
// passing a new Context into every request.
func ContextWrapper(h ContextHTTPHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(context.Background(), w, r)
	})
}
