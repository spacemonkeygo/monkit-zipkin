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
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"sync"
)

type locker struct {
	l sync.Mutex
	s rand.Source
}

func (l *locker) Int63() (rv int64) {
	l.l.Lock()
	rv = l.s.Int63()
	l.l.Unlock()
	return rv
}

func (l *locker) Seed(seed int64) {
	l.l.Lock()
	l.s.Seed(seed)
	l.l.Unlock()
}

func seed() int64 {
	var seed [8]byte
	_, err := crypto_rand.Read(seed[:])
	if err != nil {
		panic(err)
	}
	return int64(binary.BigEndian.Uint64(seed[:]))
}

var rng = rand.New(&locker{s: rand.NewSource(seed())})
