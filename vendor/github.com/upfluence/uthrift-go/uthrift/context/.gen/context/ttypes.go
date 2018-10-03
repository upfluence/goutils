// Autogenerated by Thrift Compiler (2.0.1-upfluence)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package context

import (
	"bytes"
	"fmt"
	"github.com/upfluence/thrift/lib/go/thrift"
	"github.com/upfluence/base/version"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal
var _ = version.GoUnusedProtection__
var GoUnusedProtection__ int

// Attributes:
//  - InstanceName
//  - AppName
//  - ProjectName
//  - Environment
//  - Version
//  - InterfaceVersions
type Peer struct {
	InstanceName      *string                     `thrift:"instance_name,1" json:"instance_name,omitempty"`
	AppName           *string                     `thrift:"app_name,2" json:"app_name,omitempty"`
	ProjectName       *string                     `thrift:"project_name,3" json:"project_name,omitempty"`
	Environment       *string                     `thrift:"environment,4" json:"environment,omitempty"`
	Version           *version.Version            `thrift:"version,5" json:"version,omitempty"`
	InterfaceVersions map[string]*version.Version `thrift:"interface_versions,6" json:"interface_versions,omitempty"`
}

func NewPeer() *Peer {
	return &Peer{}
}

var Peer_InstanceName_DEFAULT string

func (p *Peer) GetInstanceName() string {
	if !p.IsSetInstanceName() {
		return Peer_InstanceName_DEFAULT
	}
	return *p.InstanceName
}

func (p *Peer) SetInstanceName(v string) {
	p.InstanceName = &v
}

var Peer_AppName_DEFAULT string

func (p *Peer) GetAppName() string {
	if !p.IsSetAppName() {
		return Peer_AppName_DEFAULT
	}
	return *p.AppName
}

func (p *Peer) SetAppName(v string) {
	p.AppName = &v
}

var Peer_ProjectName_DEFAULT string

func (p *Peer) GetProjectName() string {
	if !p.IsSetProjectName() {
		return Peer_ProjectName_DEFAULT
	}
	return *p.ProjectName
}

func (p *Peer) SetProjectName(v string) {
	p.ProjectName = &v
}

var Peer_Environment_DEFAULT string

func (p *Peer) GetEnvironment() string {
	if !p.IsSetEnvironment() {
		return Peer_Environment_DEFAULT
	}
	return *p.Environment
}

func (p *Peer) SetEnvironment(v string) {
	p.Environment = &v
}

var Peer_Version_DEFAULT *version.Version

func (p *Peer) GetVersion() *version.Version {
	if !p.IsSetVersion() {
		return Peer_Version_DEFAULT
	}
	return p.Version
}

func (p *Peer) SetVersion(v *version.Version) {
	p.Version = v
}

var Peer_InterfaceVersions_DEFAULT map[string]*version.Version

func (p *Peer) GetInterfaceVersions() map[string]*version.Version {
	return p.InterfaceVersions
}

func (p *Peer) SetInterfaceVersions(v map[string]*version.Version) {
	p.InterfaceVersions = v
}
func (p *Peer) IsSetInstanceName() bool {
	return p.InstanceName != nil
}

func (p *Peer) IsSetAppName() bool {
	return p.AppName != nil
}

func (p *Peer) IsSetProjectName() bool {
	return p.ProjectName != nil
}

func (p *Peer) IsSetEnvironment() bool {
	return p.Environment != nil
}

func (p *Peer) IsSetVersion() bool {
	return p.Version != nil
}

func (p *Peer) IsSetInterfaceVersions() bool {
	return p.InterfaceVersions != nil
}

func (p *Peer) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Peer) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.InstanceName = &v
	}
	return nil
}

func (p *Peer) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AppName = &v
	}
	return nil
}

func (p *Peer) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.ProjectName = &v
	}
	return nil
}

func (p *Peer) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Environment = &v
	}
	return nil
}

func (p *Peer) ReadField5(iprot thrift.TProtocol) error {
	p.Version = version.NewVersion()
	if err := p.Version.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Version), err)
	}
	return nil
}

func (p *Peer) ReadField6(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]*version.Version, size)
	p.InterfaceVersions = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key0 = v
		}
		_val1 := version.NewVersion()
		if err := _val1.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _val1), err)
		}
		p.InterfaceVersions[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *Peer) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Peer"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
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
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
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

func (p *Peer) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetInstanceName() {
		if err := oprot.WriteFieldBegin("instance_name", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:instance_name: ", p), err)
		}
		if err := oprot.WriteString(string(*p.InstanceName)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.instance_name (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:instance_name: ", p), err)
		}
	}
	return err
}

func (p *Peer) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetAppName() {
		if err := oprot.WriteFieldBegin("app_name", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:app_name: ", p), err)
		}
		if err := oprot.WriteString(string(*p.AppName)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.app_name (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:app_name: ", p), err)
		}
	}
	return err
}

func (p *Peer) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetProjectName() {
		if err := oprot.WriteFieldBegin("project_name", thrift.STRING, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:project_name: ", p), err)
		}
		if err := oprot.WriteString(string(*p.ProjectName)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.project_name (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:project_name: ", p), err)
		}
	}
	return err
}

func (p *Peer) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetEnvironment() {
		if err := oprot.WriteFieldBegin("environment", thrift.STRING, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:environment: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Environment)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.environment (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:environment: ", p), err)
		}
	}
	return err
}

func (p *Peer) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetVersion() {
		if err := oprot.WriteFieldBegin("version", thrift.STRUCT, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:version: ", p), err)
		}
		if err := p.Version.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Version), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:version: ", p), err)
		}
	}
	return err
}

func (p *Peer) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetInterfaceVersions() {
		if err := oprot.WriteFieldBegin("interface_versions", thrift.MAP, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:interface_versions: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRUCT, len(p.InterfaceVersions)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.InterfaceVersions {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:interface_versions: ", p), err)
		}
	}
	return err
}

func (p *Peer) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Peer(%+v)", *p)
}

// Attributes:
//  - DeadlineSec
//  - DeadlineNano
//  - Values
//  - Client
//  - TraceID
//  - SpanID
type Context struct {
	DeadlineSec  *int64            `thrift:"deadline_sec,1" json:"deadline_sec,omitempty"`
	DeadlineNano *int64            `thrift:"deadline_nano,2" json:"deadline_nano,omitempty"`
	Values       map[string][]byte `thrift:"values,3" json:"values,omitempty"`
	Client       *Peer             `thrift:"client,4" json:"client,omitempty"`
	TraceID      *string           `thrift:"trace_id,5" json:"trace_id,omitempty"`
	SpanID       *string           `thrift:"span_id,6" json:"span_id,omitempty"`
}

func NewContext() *Context {
	return &Context{}
}

var Context_DeadlineSec_DEFAULT int64

func (p *Context) GetDeadlineSec() int64 {
	if !p.IsSetDeadlineSec() {
		return Context_DeadlineSec_DEFAULT
	}
	return *p.DeadlineSec
}

func (p *Context) SetDeadlineSec(v int64) {
	p.DeadlineSec = &v
}

var Context_DeadlineNano_DEFAULT int64

func (p *Context) GetDeadlineNano() int64 {
	if !p.IsSetDeadlineNano() {
		return Context_DeadlineNano_DEFAULT
	}
	return *p.DeadlineNano
}

func (p *Context) SetDeadlineNano(v int64) {
	p.DeadlineNano = &v
}

var Context_Values_DEFAULT map[string][]byte

func (p *Context) GetValues() map[string][]byte {
	return p.Values
}

func (p *Context) SetValues(v map[string][]byte) {
	p.Values = v
}

var Context_Client_DEFAULT *Peer

func (p *Context) GetClient() *Peer {
	if !p.IsSetClient() {
		return Context_Client_DEFAULT
	}
	return p.Client
}

func (p *Context) SetClient(v *Peer) {
	p.Client = v
}

var Context_TraceID_DEFAULT string

func (p *Context) GetTraceID() string {
	if !p.IsSetTraceID() {
		return Context_TraceID_DEFAULT
	}
	return *p.TraceID
}

func (p *Context) SetTraceID(v string) {
	p.TraceID = &v
}

var Context_SpanID_DEFAULT string

func (p *Context) GetSpanID() string {
	if !p.IsSetSpanID() {
		return Context_SpanID_DEFAULT
	}
	return *p.SpanID
}

func (p *Context) SetSpanID(v string) {
	p.SpanID = &v
}
func (p *Context) IsSetDeadlineSec() bool {
	return p.DeadlineSec != nil
}

func (p *Context) IsSetDeadlineNano() bool {
	return p.DeadlineNano != nil
}

func (p *Context) IsSetValues() bool {
	return p.Values != nil
}

func (p *Context) IsSetClient() bool {
	return p.Client != nil
}

func (p *Context) IsSetTraceID() bool {
	return p.TraceID != nil
}

func (p *Context) IsSetSpanID() bool {
	return p.SpanID != nil
}

func (p *Context) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Context) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.DeadlineSec = &v
	}
	return nil
}

func (p *Context) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.DeadlineNano = &v
	}
	return nil
}

func (p *Context) ReadField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string][]byte, size)
	p.Values = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key2 = v
		}
		var _val3 []byte
		if v, err := iprot.ReadBinary(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val3 = v
		}
		p.Values[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *Context) ReadField4(iprot thrift.TProtocol) error {
	p.Client = NewPeer()
	if err := p.Client.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Client), err)
	}
	return nil
}

func (p *Context) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.TraceID = &v
	}
	return nil
}

func (p *Context) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.SpanID = &v
	}
	return nil
}

func (p *Context) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Context"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
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
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
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

func (p *Context) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetDeadlineSec() {
		if err := oprot.WriteFieldBegin("deadline_sec", thrift.I64, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:deadline_sec: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.DeadlineSec)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.deadline_sec (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:deadline_sec: ", p), err)
		}
	}
	return err
}

func (p *Context) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetDeadlineNano() {
		if err := oprot.WriteFieldBegin("deadline_nano", thrift.I64, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:deadline_nano: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.DeadlineNano)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.deadline_nano (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:deadline_nano: ", p), err)
		}
	}
	return err
}

func (p *Context) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetValues() {
		if err := oprot.WriteFieldBegin("values", thrift.MAP, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:values: ", p), err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Values)); err != nil {
			return thrift.PrependError("error writing map begin: ", err)
		}
		for k, v := range p.Values {
			if err := oprot.WriteString(string(k)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
			if err := oprot.WriteBinary(v); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return thrift.PrependError("error writing map end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:values: ", p), err)
		}
	}
	return err
}

func (p *Context) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetClient() {
		if err := oprot.WriteFieldBegin("client", thrift.STRUCT, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:client: ", p), err)
		}
		if err := p.Client.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Client), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:client: ", p), err)
		}
	}
	return err
}

func (p *Context) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetTraceID() {
		if err := oprot.WriteFieldBegin("trace_id", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:trace_id: ", p), err)
		}
		if err := oprot.WriteString(string(*p.TraceID)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.trace_id (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:trace_id: ", p), err)
		}
	}
	return err
}

func (p *Context) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetSpanID() {
		if err := oprot.WriteFieldBegin("span_id", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:span_id: ", p), err)
		}
		if err := oprot.WriteString(string(*p.SpanID)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.span_id (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:span_id: ", p), err)
		}
	}
	return err
}

func (p *Context) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Context(%+v)", *p)
}