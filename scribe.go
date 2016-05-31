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
	"encoding/base64"
	"errors"
	"log"
	"net"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"gopkg.in/spacemonkeygo/monkit-zipkin.v2/gen-go/scribe"
	"gopkg.in/spacemonkeygo/monkit-zipkin.v2/gen-go/zipkin"
)

// ScribeCollector matches the TraceCollector interface, but writes directly
// to a connected Scribe socket.
type ScribeCollector struct {
	addr *net.TCPAddr
	done chan struct{}

	logs chan *scribe.LogEntry
}

// NewScribeCollector creates a ScribeCollector. scribe_addr is the address
// of the Scribe endpoint, typically "127.0.0.1:9410"
func NewScribeCollector(scribe_addr string) (*ScribeCollector, error) {
	sa, err := net.ResolveTCPAddr("tcp", scribe_addr)
	if err != nil {
		return nil, err
	}

	s := ScribeCollector{
		addr: sa,
		done: make(chan struct{}),
		logs: make(chan *scribe.LogEntry, 100),
	}

	go s.pumpWrites()

	return &s, nil
}

// pumpWrites sends all messages on s.logs to scribe. It creates the
// scribe connections, recreates them when write errors occur, and
// backs off on consecutive connection errors.
//
// When a write error occurs, pumpWrites will lose that log entry.
func (s *ScribeCollector) pumpWrites() {
	var backoff int

	for {
		select {
		case <-s.done:
			return
		case <-time.After(defaultBackoff.duration(backoff)):
		}

		conn, err := newScribeConn(s.addr)
		if err != nil {
			log.Printf("connect error: %s", err)
			backoff++
			continue
		}

		err = writeAll(conn, s.logs, s.done)
		if err != nil {
			log.Printf("write error: %s", err)
		}

		backoff = 0
	}
}

// writeAll sends all logs to c, stopping when done is signaled.
func writeAll(c *scribeConn, logs <-chan *scribe.LogEntry,
	done <-chan struct{}) error {
	for {
		select {
		case l := <-logs:
			rc, err := c.client.Log([]*scribe.LogEntry{l})
			if err != nil {
				return err
			}

			// Non-OK responses are logged but not fatal
			// for the connection.
			if rc != scribe.ResultCode_OK {
				log.Printf("scribe result code not OK: %s", rc)
			}
		case <-done:
			return nil
		}
	}
}

// Close closes an existing ScribeCollector
func (s *ScribeCollector) Close() error {
	close(s.done)
	return nil
}

// CollectSerialized buffers a serialized zipkin.Span to be sent to
// the Scribe endpoint. It returns an error and loses the log entry if
// the buffer is full.
func (c *ScribeCollector) CollectSerialized(serialized []byte) error {
	entry := scribe.LogEntry{
		Category: "zipkin",
		Message:  base64.StdEncoding.EncodeToString(serialized),
	}

	select {
	case c.logs <- &entry:
		return nil
	default:
		return errors.New("skipping scribe log: buffer full")
	}
}

// Collect will serialize and send a zipkin.Span to the configured Scribe
// endpoint
func (c *ScribeCollector) Collect(s *zipkin.Span) {
	t := thrift.NewTMemoryBuffer()
	p := thrift.NewTBinaryProtocolTransport(t)
	err := s.Write(p)
	if err != nil {
		log.Printf("failed write: %v", err)
		return
	}
	err = c.CollectSerialized(t.Buffer.Bytes())
	if err != nil {
		log.Printf("failed collect: %v", err)
	}
}

type scribeConn struct {
	transport *thrift.TFramedTransport
	client    *scribe.ScribeClient
}

func newScribeConn(addr *net.TCPAddr) (*scribeConn, error) {
	transport := thrift.NewTFramedTransport(
		thrift.NewTSocketFromAddrTimeout(addr, 10*time.Second))
	err := transport.Open()
	if err != nil {
		return nil, err
	}

	proto := thrift.NewTBinaryProtocolTransport(transport)
	conn := scribeConn{
		transport: transport,
		client:    scribe.NewScribeClientProtocol(transport, proto, proto),
	}

	return &conn, nil
}

func (c *scribeConn) Close() {
	c.transport.Close()
}

var _ TraceCollector = (*ScribeCollector)(nil)
