// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package zipkin

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type AnnotationType int64

const (
	AnnotationType_BOOL   AnnotationType = 0
	AnnotationType_BYTES  AnnotationType = 1
	AnnotationType_I16    AnnotationType = 2
	AnnotationType_I32    AnnotationType = 3
	AnnotationType_I64    AnnotationType = 4
	AnnotationType_DOUBLE AnnotationType = 5
	AnnotationType_STRING AnnotationType = 6
)

func (p AnnotationType) String() string {
	switch p {
	case AnnotationType_BOOL:
		return "AnnotationType_BOOL"
	case AnnotationType_BYTES:
		return "AnnotationType_BYTES"
	case AnnotationType_I16:
		return "AnnotationType_I16"
	case AnnotationType_I32:
		return "AnnotationType_I32"
	case AnnotationType_I64:
		return "AnnotationType_I64"
	case AnnotationType_DOUBLE:
		return "AnnotationType_DOUBLE"
	case AnnotationType_STRING:
		return "AnnotationType_STRING"
	}
	return "<UNSET>"
}

func AnnotationTypeFromString(s string) (AnnotationType, error) {
	switch s {
	case "AnnotationType_BOOL":
		return AnnotationType_BOOL, nil
	case "AnnotationType_BYTES":
		return AnnotationType_BYTES, nil
	case "AnnotationType_I16":
		return AnnotationType_I16, nil
	case "AnnotationType_I32":
		return AnnotationType_I32, nil
	case "AnnotationType_I64":
		return AnnotationType_I64, nil
	case "AnnotationType_DOUBLE":
		return AnnotationType_DOUBLE, nil
	case "AnnotationType_STRING":
		return AnnotationType_STRING, nil
	}
	return AnnotationType(0), fmt.Errorf("not a valid AnnotationType string")
}

func AnnotationTypePtr(v AnnotationType) *AnnotationType { return &v }

type Endpoint struct {
	Ipv4        int32  `thrift:"ipv4,1" json:"ipv4"`
	Port        int16  `thrift:"port,2" json:"port"`
	ServiceName string `thrift:"service_name,3" json:"service_name"`
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (p *Endpoint) GetIpv4() int32 {
	return p.Ipv4
}

func (p *Endpoint) GetPort() int16 {
	return p.Port
}

func (p *Endpoint) GetServiceName() string {
	return p.ServiceName
}
func (p *Endpoint) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Endpoint) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Ipv4 = v
	}
	return nil
}

func (p *Endpoint) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Port = v
	}
	return nil
}

func (p *Endpoint) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.ServiceName = v
	}
	return nil
}

func (p *Endpoint) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Endpoint"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Endpoint) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ipv4", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:ipv4: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Ipv4)); err != nil {
		return fmt.Errorf("%T.ipv4 (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:ipv4: %s", p, err)
	}
	return err
}

func (p *Endpoint) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("port", thrift.I16, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:port: %s", p, err)
	}
	if err := oprot.WriteI16(int16(p.Port)); err != nil {
		return fmt.Errorf("%T.port (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:port: %s", p, err)
	}
	return err
}

func (p *Endpoint) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("service_name", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:service_name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.ServiceName)); err != nil {
		return fmt.Errorf("%T.service_name (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:service_name: %s", p, err)
	}
	return err
}

func (p *Endpoint) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Endpoint(%+v)", *p)
}

type Annotation struct {
	Timestamp int64     `thrift:"timestamp,1" json:"timestamp"`
	Value     string    `thrift:"value,2" json:"value"`
	Host      *Endpoint `thrift:"host,3" json:"host"`
	Duration  *int32    `thrift:"duration,4" json:"duration"`
}

func NewAnnotation() *Annotation {
	return &Annotation{}
}

func (p *Annotation) GetTimestamp() int64 {
	return p.Timestamp
}

func (p *Annotation) GetValue() string {
	return p.Value
}

var Annotation_Host_DEFAULT *Endpoint

func (p *Annotation) GetHost() *Endpoint {
	if !p.IsSetHost() {
		return Annotation_Host_DEFAULT
	}
	return p.Host
}

var Annotation_Duration_DEFAULT int32

func (p *Annotation) GetDuration() int32 {
	if !p.IsSetDuration() {
		return Annotation_Duration_DEFAULT
	}
	return *p.Duration
}
func (p *Annotation) IsSetHost() bool {
	return p.Host != nil
}

func (p *Annotation) IsSetDuration() bool {
	return p.Duration != nil
}

func (p *Annotation) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Annotation) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Timestamp = v
	}
	return nil
}

func (p *Annotation) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *Annotation) ReadField3(iprot thrift.TProtocol) error {
	p.Host = &Endpoint{}
	if err := p.Host.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Host, err)
	}
	return nil
}

func (p *Annotation) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Duration = &v
	}
	return nil
}

func (p *Annotation) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Annotation"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Annotation) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:timestamp: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Timestamp)); err != nil {
		return fmt.Errorf("%T.timestamp (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:timestamp: %s", p, err)
	}
	return err
}

func (p *Annotation) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:value: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Value)); err != nil {
		return fmt.Errorf("%T.value (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:value: %s", p, err)
	}
	return err
}

func (p *Annotation) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetHost() {
		if err := oprot.WriteFieldBegin("host", thrift.STRUCT, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:host: %s", p, err)
		}
		if err := p.Host.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Host, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:host: %s", p, err)
		}
	}
	return err
}

func (p *Annotation) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetDuration() {
		if err := oprot.WriteFieldBegin("duration", thrift.I32, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:duration: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.Duration)); err != nil {
			return fmt.Errorf("%T.duration (4) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:duration: %s", p, err)
		}
	}
	return err
}

func (p *Annotation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Annotation(%+v)", *p)
}

type BinaryAnnotation struct {
	Key            string         `thrift:"key,1" json:"key"`
	Value          []byte         `thrift:"value,2" json:"value"`
	AnnotationType AnnotationType `thrift:"annotation_type,3" json:"annotation_type"`
	Host           *Endpoint      `thrift:"host,4" json:"host"`
}

func NewBinaryAnnotation() *BinaryAnnotation {
	return &BinaryAnnotation{}
}

func (p *BinaryAnnotation) GetKey() string {
	return p.Key
}

func (p *BinaryAnnotation) GetValue() []byte {
	return p.Value
}

func (p *BinaryAnnotation) GetAnnotationType() AnnotationType {
	return p.AnnotationType
}

var BinaryAnnotation_Host_DEFAULT *Endpoint

func (p *BinaryAnnotation) GetHost() *Endpoint {
	if !p.IsSetHost() {
		return BinaryAnnotation_Host_DEFAULT
	}
	return p.Host
}
func (p *BinaryAnnotation) IsSetHost() bool {
	return p.Host != nil
}

func (p *BinaryAnnotation) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *BinaryAnnotation) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *BinaryAnnotation) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *BinaryAnnotation) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		temp := AnnotationType(v)
		p.AnnotationType = temp
	}
	return nil
}

func (p *BinaryAnnotation) ReadField4(iprot thrift.TProtocol) error {
	p.Host = &Endpoint{}
	if err := p.Host.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Host, err)
	}
	return nil
}

func (p *BinaryAnnotation) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BinaryAnnotation"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *BinaryAnnotation) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:key: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Key)); err != nil {
		return fmt.Errorf("%T.key (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:key: %s", p, err)
	}
	return err
}

func (p *BinaryAnnotation) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:value: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return fmt.Errorf("%T.value (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:value: %s", p, err)
	}
	return err
}

func (p *BinaryAnnotation) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("annotation_type", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:annotation_type: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.AnnotationType)); err != nil {
		return fmt.Errorf("%T.annotation_type (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:annotation_type: %s", p, err)
	}
	return err
}

func (p *BinaryAnnotation) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetHost() {
		if err := oprot.WriteFieldBegin("host", thrift.STRUCT, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:host: %s", p, err)
		}
		if err := p.Host.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Host, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:host: %s", p, err)
		}
	}
	return err
}

func (p *BinaryAnnotation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BinaryAnnotation(%+v)", *p)
}

type Span struct {
	TraceId int64 `thrift:"trace_id,1" json:"trace_id"`
	// unused field # 2
	Name        string        `thrift:"name,3" json:"name"`
	Id          int64         `thrift:"id,4" json:"id"`
	ParentId    *int64        `thrift:"parent_id,5" json:"parent_id"`
	Annotations []*Annotation `thrift:"annotations,6" json:"annotations"`
	// unused field # 7
	BinaryAnnotations []*BinaryAnnotation `thrift:"binary_annotations,8" json:"binary_annotations"`
	Debug             bool                `thrift:"debug,9" json:"debug"`
}

func NewSpan() *Span {
	return &Span{}
}

func (p *Span) GetTraceId() int64 {
	return p.TraceId
}

func (p *Span) GetName() string {
	return p.Name
}

func (p *Span) GetId() int64 {
	return p.Id
}

var Span_ParentId_DEFAULT int64

func (p *Span) GetParentId() int64 {
	if !p.IsSetParentId() {
		return Span_ParentId_DEFAULT
	}
	return *p.ParentId
}

func (p *Span) GetAnnotations() []*Annotation {
	return p.Annotations
}

func (p *Span) GetBinaryAnnotations() []*BinaryAnnotation {
	return p.BinaryAnnotations
}

var Span_Debug_DEFAULT bool = false

func (p *Span) GetDebug() bool {
	return p.Debug
}
func (p *Span) IsSetParentId() bool {
	return p.ParentId != nil
}

func (p *Span) IsSetDebug() bool {
	return p.Debug != Span_Debug_DEFAULT
}

func (p *Span) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.ReadField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.ReadField9(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Span) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.TraceId = v
	}
	return nil
}

func (p *Span) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *Span) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Id = v
	}
	return nil
}

func (p *Span) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.ParentId = &v
	}
	return nil
}

func (p *Span) ReadField6(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*Annotation, 0, size)
	p.Annotations = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &Annotation{}
		if err := _elem0.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem0, err)
		}
		p.Annotations = append(p.Annotations, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Span) ReadField8(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*BinaryAnnotation, 0, size)
	p.BinaryAnnotations = tSlice
	for i := 0; i < size; i++ {
		_elem1 := &BinaryAnnotation{}
		if err := _elem1.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem1, err)
		}
		p.BinaryAnnotations = append(p.BinaryAnnotations, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *Span) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 9: %s", err)
	} else {
		p.Debug = v
	}
	return nil
}

func (p *Span) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Span"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Span) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("trace_id", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:trace_id: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.TraceId)); err != nil {
		return fmt.Errorf("%T.trace_id (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:trace_id: %s", p, err)
	}
	return err
}

func (p *Span) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return fmt.Errorf("%T.name (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:name: %s", p, err)
	}
	return err
}

func (p *Span) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:id: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Id)); err != nil {
		return fmt.Errorf("%T.id (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:id: %s", p, err)
	}
	return err
}

func (p *Span) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetParentId() {
		if err := oprot.WriteFieldBegin("parent_id", thrift.I64, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:parent_id: %s", p, err)
		}
		if err := oprot.WriteI64(int64(*p.ParentId)); err != nil {
			return fmt.Errorf("%T.parent_id (5) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:parent_id: %s", p, err)
		}
	}
	return err
}

func (p *Span) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("annotations", thrift.LIST, 6); err != nil {
		return fmt.Errorf("%T write field begin error 6:annotations: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Annotations)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.Annotations {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 6:annotations: %s", p, err)
	}
	return err
}

func (p *Span) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("binary_annotations", thrift.LIST, 8); err != nil {
		return fmt.Errorf("%T write field begin error 8:binary_annotations: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.BinaryAnnotations)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.BinaryAnnotations {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 8:binary_annotations: %s", p, err)
	}
	return err
}

func (p *Span) writeField9(oprot thrift.TProtocol) (err error) {
	if p.IsSetDebug() {
		if err := oprot.WriteFieldBegin("debug", thrift.BOOL, 9); err != nil {
			return fmt.Errorf("%T write field begin error 9:debug: %s", p, err)
		}
		if err := oprot.WriteBool(bool(p.Debug)); err != nil {
			return fmt.Errorf("%T.debug (9) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 9:debug: %s", p, err)
		}
	}
	return err
}

func (p *Span) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Span(%+v)", *p)
}
