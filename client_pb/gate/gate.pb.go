// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gate.proto

/*
Package gate is a generated protocol buffer package.

It is generated from these files:
	gate.proto

It has these top-level messages:
	GateAuthReq
	GateAuthRsp
	GateHeartBeatReq
	GateHeartBeatRsp
	GateAnotherLoginNtf
*/
package gate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ErrCode int32

const (
	ErrCode_SUCCESS            ErrCode = 0
	ErrCode_UNKNOW             ErrCode = 1
	ErrCode_ERR_INVALID_TOKEN  ErrCode = 101
	ErrCode_ERR_EXPIRE_TOKEN   ErrCode = 102
	ErrCode_ERR_ALREADY_AUTHED ErrCode = 103
)

var ErrCode_name = map[int32]string{
	0:   "SUCCESS",
	1:   "UNKNOW",
	101: "ERR_INVALID_TOKEN",
	102: "ERR_EXPIRE_TOKEN",
	103: "ERR_ALREADY_AUTHED",
}
var ErrCode_value = map[string]int32{
	"SUCCESS":            0,
	"UNKNOW":             1,
	"ERR_INVALID_TOKEN":  101,
	"ERR_EXPIRE_TOKEN":   102,
	"ERR_ALREADY_AUTHED": 103,
}

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}
func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}
func (x *ErrCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrCode_value, data, "ErrCode")
	if err != nil {
		return err
	}
	*x = ErrCode(value)
	return nil
}
func (ErrCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// GateAuthReq 网关认证请求
type GateAuthReq struct {
	PlayerId *uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	// ...
	Expire           *int64  `protobuf:"varint,14,opt,name=expire" json:"expire,omitempty"`
	Token            *string `protobuf:"bytes,15,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GateAuthReq) Reset()                    { *m = GateAuthReq{} }
func (m *GateAuthReq) String() string            { return proto.CompactTextString(m) }
func (*GateAuthReq) ProtoMessage()               {}
func (*GateAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GateAuthReq) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *GateAuthReq) GetExpire() int64 {
	if m != nil && m.Expire != nil {
		return *m.Expire
	}
	return 0
}

func (m *GateAuthReq) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// GateAuthRsp 网关认证应答
type GateAuthRsp struct {
	ErrCode          *ErrCode `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=gate.ErrCode" json:"err_code,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GateAuthRsp) Reset()                    { *m = GateAuthRsp{} }
func (m *GateAuthRsp) String() string            { return proto.CompactTextString(m) }
func (*GateAuthRsp) ProtoMessage()               {}
func (*GateAuthRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GateAuthRsp) GetErrCode() ErrCode {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return ErrCode_SUCCESS
}

// GateHeartBeatReq 心跳检测请求
type GateHeartBeatReq struct {
	Reserve          *uint32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GateHeartBeatReq) Reset()                    { *m = GateHeartBeatReq{} }
func (m *GateHeartBeatReq) String() string            { return proto.CompactTextString(m) }
func (*GateHeartBeatReq) ProtoMessage()               {}
func (*GateHeartBeatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GateHeartBeatReq) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// GateHeartBeatRsp 心跳检测应答
type GateHeartBeatRsp struct {
	Reserve          *uint32 `protobuf:"varint,2,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GateHeartBeatRsp) Reset()                    { *m = GateHeartBeatRsp{} }
func (m *GateHeartBeatRsp) String() string            { return proto.CompactTextString(m) }
func (*GateHeartBeatRsp) ProtoMessage()               {}
func (*GateHeartBeatRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GateHeartBeatRsp) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// GateAnotherLoginNtf 顶号通知
type GateAnotherLoginNtf struct {
	Reserve          *uint32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GateAnotherLoginNtf) Reset()                    { *m = GateAnotherLoginNtf{} }
func (m *GateAnotherLoginNtf) String() string            { return proto.CompactTextString(m) }
func (*GateAnotherLoginNtf) ProtoMessage()               {}
func (*GateAnotherLoginNtf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GateAnotherLoginNtf) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

func init() {
	proto.RegisterType((*GateAuthReq)(nil), "gate.GateAuthReq")
	proto.RegisterType((*GateAuthRsp)(nil), "gate.GateAuthRsp")
	proto.RegisterType((*GateHeartBeatReq)(nil), "gate.GateHeartBeatReq")
	proto.RegisterType((*GateHeartBeatRsp)(nil), "gate.GateHeartBeatRsp")
	proto.RegisterType((*GateAnotherLoginNtf)(nil), "gate.GateAnotherLoginNtf")
	proto.RegisterEnum("gate.ErrCode", ErrCode_name, ErrCode_value)
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0x8d, 0xd6, 0xa6, 0x7d, 0xa5, 0x75, 0x5d, 0xb5, 0x04, 0xbc, 0x94, 0x9e, 0x82, 0x48,
	0x05, 0x2f, 0x9e, 0x63, 0xbb, 0xd8, 0xd0, 0x90, 0xca, 0xa6, 0xd1, 0x7a, 0x0a, 0xc1, 0xbc, 0xa6,
	0x41, 0xc9, 0xc6, 0xcd, 0x2a, 0xfa, 0xef, 0x65, 0x93, 0x08, 0x15, 0xf1, 0xf6, 0x66, 0x98, 0x19,
	0x3e, 0x1e, 0x40, 0x1a, 0x2b, 0x9c, 0x14, 0x52, 0x28, 0x41, 0x5b, 0xfa, 0x1e, 0xaf, 0xa1, 0x77,
	0x17, 0x2b, 0x74, 0xde, 0xd5, 0x96, 0xe3, 0x1b, 0x3d, 0x87, 0x6e, 0xf1, 0x1a, 0x7f, 0xa1, 0x8c,
	0xb2, 0xc4, 0x32, 0x46, 0x86, 0xdd, 0xe2, 0x9d, 0xda, 0x70, 0x13, 0x3a, 0x84, 0x36, 0x7e, 0x16,
	0x99, 0x44, 0x6b, 0x30, 0x32, 0xec, 0x03, 0xde, 0x28, 0x7a, 0x0a, 0x87, 0x4a, 0xbc, 0x60, 0x6e,
	0x1d, 0x8d, 0x0c, 0xbb, 0xcb, 0x6b, 0x31, 0xbe, 0xd9, 0x59, 0x2e, 0x0b, 0x6a, 0x43, 0x07, 0xa5,
	0x8c, 0x9e, 0x45, 0x82, 0xd5, 0xf0, 0xe0, 0xba, 0x3f, 0xa9, 0x68, 0x98, 0x94, 0x53, 0x91, 0x20,
	0x37, 0xb1, 0x3e, 0xc6, 0x97, 0x40, 0x74, 0x71, 0x8e, 0xb1, 0x54, 0xb7, 0x18, 0x2b, 0xcd, 0x65,
	0x81, 0x29, 0xb1, 0x44, 0xf9, 0x51, 0x97, 0xfb, 0xfc, 0x47, 0xfe, 0x4d, 0x97, 0xc5, 0x6e, 0x7a,
	0xff, 0x77, 0xfa, 0x0a, 0x4e, 0x2a, 0xa8, 0x5c, 0xa8, 0x2d, 0x4a, 0x4f, 0xa4, 0x59, 0xee, 0xab,
	0xcd, 0xff, 0xf3, 0x17, 0x29, 0x98, 0x0d, 0x20, 0xed, 0x81, 0x19, 0x84, 0xd3, 0x29, 0x0b, 0x02,
	0xb2, 0x47, 0x01, 0xda, 0xa1, 0xbf, 0xf0, 0x97, 0x8f, 0xc4, 0xa0, 0x67, 0x70, 0xcc, 0x38, 0x8f,
	0x5c, 0xff, 0xc1, 0xf1, 0xdc, 0x59, 0xb4, 0x5a, 0x2e, 0x98, 0x4f, 0xf4, 0x5b, 0x88, 0xb6, 0xd9,
	0xfa, 0xde, 0xe5, 0xac, 0x71, 0x37, 0x74, 0x08, 0x54, 0xbb, 0x8e, 0xc7, 0x99, 0x33, 0x7b, 0x8a,
	0x9c, 0x70, 0x35, 0x67, 0x33, 0x92, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x5e, 0x46, 0x66,
	0x9b, 0x01, 0x00, 0x00,
}
