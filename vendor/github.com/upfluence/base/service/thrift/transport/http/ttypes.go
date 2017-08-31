// Autogenerated by Thrift Compiler (1.0.0-upfluence)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package http

import (
	"bytes"
	"fmt"
	"github.com/upfluence/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal
var GoUnusedProtection__ int

type Method int64

const (
	Method_POST   Method = 0
	Method_GET    Method = 1
	Method_PUT    Method = 2
	Method_DELETE Method = 3
)

func (p Method) String() string {
	switch p {
	case Method_POST:
		return "Method_POST"
	case Method_GET:
		return "Method_GET"
	case Method_PUT:
		return "Method_PUT"
	case Method_DELETE:
		return "Method_DELETE"
	}
	return "<UNSET>"
}

func MethodFromString(s string) (Method, error) {
	switch s {
	case "Method_POST":
		return Method_POST, nil
	case "Method_GET":
		return Method_GET, nil
	case "Method_PUT":
		return Method_PUT, nil
	case "Method_DELETE":
		return Method_DELETE, nil
	}
	return Method(0), fmt.Errorf("not a valid Method string")
}

func MethodPtr(v Method) *Method { return &v }

// Attributes:
//  - Url
//  - Method
type Transport struct {
	Url    string `thrift:"url,1,required" json:"url"`
	Method Method `thrift:"method,2,required" json:"method"`
}

func NewTransport() *Transport {
	return &Transport{}
}

func (p *Transport) GetUrl() string {
	return p.Url
}

func (p *Transport) GetMethod() Method {
	return p.Method
}
func (p *Transport) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetUrl bool = false
	var issetMethod bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
			issetUrl = true
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
			issetMethod = true
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetUrl {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Url is not set"))
	}
	if !issetMethod {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Method is not set"))
	}
	return nil
}

func (p *Transport) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Url = v
	}
	return nil
}

func (p *Transport) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := Method(v)
		p.Method = temp
	}
	return nil
}

func (p *Transport) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Transport"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Transport) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("url", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:url: ", p), err)
	}
	if err := oprot.WriteString(string(p.Url)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.url (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:url: ", p), err)
	}
	return err
}

func (p *Transport) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("method", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:method: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Method)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.method (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:method: ", p), err)
	}
	return err
}

func (p *Transport) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Transport(%+v)", *p)
}
