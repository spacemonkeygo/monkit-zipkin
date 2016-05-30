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
	"gopkg.in/spacemonkeygo/monitor-zipkin.v2/gen-go/zipkin"
)

// TraceCollector is an interface dealing with completed Spans on a
// SpanManager. See RegisterZipkin.
type TraceCollector interface {
	// Collect gets called with a Span whenever a Span is completed on a
	// SpanManager.
	Collect(span *zipkin.Span)
}
