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

/*
Package zipkin provides a monkit plugin for sending traces to Zipkin.

Example usage

Your main method gets set up something like this:

  package main

  import (
    "net/http"

    "gopkg.in/spacemonkeygo/monkit.v2"
    "gopkg.in/spacemonkeygo/monkit.v2/environment"
    "gopkg.in/spacemonkeygo/monkit.v2/present"
    "gopkg.in/spacemonkeygo/monkit-zipkin.v2"
  )

  func main() {
    environment.Register(monkit.Default)
    go http.ListenAndServe("localhost:9000", present.HTTP(monkit.Default))
    collector, err := zipkin.NewScribeCollector("zipkin.whatever:9042")
    if err != nil {
      panic(err)
    }
    zipkin.RegisterZipkin(monkit.Default, collector, zipkin.Options{})

    ...
  }

Once you've done that, you need to make sure your HTTP handlers pull Zipkin
Context info from the Request. That's easy to do with zipkin.ContextWrapper.

  func HandleRequest(ctx context.Context, w http.ResponseWriter,
    r *http.Request) {
    defer mon.Task()(&ctx)(nil)

    ... whatever
  }

  func DoTheThing(ctx context.Context) (err error) {
    defer mon.Task()(&ctx)(&err)
    return http.Serve(listener, zipkin.ContextWrapper(
      zipkin.ContextHTTPHandlerFunc(HandleRequest)))
  }

Last, your outbound HTTP requests need to pass through Context info:

  func MakeRequest(ctx context.Context) (err error) {
    defer mon.Task()(&ctx)(&err)
    req, err := http.NewRequest(...)
    if err != nil {
      return err
    }
    resp, err := zipkin.TraceRequest(ctx, http.DefaultClient, req)
    ...
  }
*/
package zipkin // import "gopkg.in/spacemonkeygo/monkit-zipkin.v2"
