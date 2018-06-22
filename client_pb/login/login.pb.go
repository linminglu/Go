// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

/*
Package login is a generated protocol buffer package.

It is generated from these files:
	login.proto

It has these top-level messages:
	LoginAuthReq
	LoginAuthRsp
*/
package login

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

type ErrorCode int32

const (
	ErrorCode_SUCCESS           ErrorCode = 0
	ErrorCode_ABNORMAL          ErrorCode = 1
	ErrorCode_ERR_EXPIRE_TOKEN  ErrorCode = 101
	ErrorCode_ERR_INVALID_TOKEN ErrorCode = 102
)

var ErrorCode_name = map[int32]string{
	0:   "SUCCESS",
	1:   "ABNORMAL",
	101: "ERR_EXPIRE_TOKEN",
	102: "ERR_INVALID_TOKEN",
}
var ErrorCode_value = map[string]int32{
	"SUCCESS":           0,
	"ABNORMAL":          1,
	"ERR_EXPIRE_TOKEN":  101,
	"ERR_INVALID_TOKEN": 102,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// LoginAuthReq 认证请求
type LoginAuthReq struct {
	AccountId   *uint64 `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AccountName *string `protobuf:"bytes,2,opt,name=account_name,json=accountName" json:"account_name,omitempty"`
	// ...
	Expire           *int64  `protobuf:"varint,14,opt,name=expire" json:"expire,omitempty"`
	Token            *string `protobuf:"bytes,15,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginAuthReq) Reset()                    { *m = LoginAuthReq{} }
func (m *LoginAuthReq) String() string            { return proto.CompactTextString(m) }
func (*LoginAuthReq) ProtoMessage()               {}
func (*LoginAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginAuthReq) GetAccountId() uint64 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *LoginAuthReq) GetAccountName() string {
	if m != nil && m.AccountName != nil {
		return *m.AccountName
	}
	return ""
}

func (m *LoginAuthReq) GetExpire() int64 {
	if m != nil && m.Expire != nil {
		return *m.Expire
	}
	return 0
}

func (m *LoginAuthReq) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// LoginAuthRsp 认证应答
type LoginAuthRsp struct {
	ErrCode *ErrorCode `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=login.ErrorCode" json:"err_code,omitempty"`
	// ...
	PlayerId         *uint64 `protobuf:"varint,3,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	GateIp           *string `protobuf:"bytes,4,opt,name=gate_ip,json=gateIp" json:"gate_ip,omitempty"`
	GatePort         *int32  `protobuf:"varint,5,opt,name=gate_port,json=gatePort" json:"gate_port,omitempty"`
	Expire           *int64  `protobuf:"varint,14,opt,name=expire" json:"expire,omitempty"`
	GateToken        *string `protobuf:"bytes,15,opt,name=gate_token,json=gateToken" json:"gate_token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginAuthRsp) Reset()                    { *m = LoginAuthRsp{} }
func (m *LoginAuthRsp) String() string            { return proto.CompactTextString(m) }
func (*LoginAuthRsp) ProtoMessage()               {}
func (*LoginAuthRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginAuthRsp) GetErrCode() ErrorCode {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return ErrorCode_SUCCESS
}

func (m *LoginAuthRsp) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *LoginAuthRsp) GetGateIp() string {
	if m != nil && m.GateIp != nil {
		return *m.GateIp
	}
	return ""
}

func (m *LoginAuthRsp) GetGatePort() int32 {
	if m != nil && m.GatePort != nil {
		return *m.GatePort
	}
	return 0
}

func (m *LoginAuthRsp) GetExpire() int64 {
	if m != nil && m.Expire != nil {
		return *m.Expire
	}
	return 0
}

func (m *LoginAuthRsp) GetGateToken() string {
	if m != nil && m.GateToken != nil {
		return *m.GateToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginAuthReq)(nil), "login.LoginAuthReq")
	proto.RegisterType((*LoginAuthRsp)(nil), "login.LoginAuthRsp")
	proto.RegisterEnum("login.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x5d, 0xa1, 0xd0, 0x0e, 0x04, 0xeb, 0x06, 0xb5, 0x89, 0x21, 0xa9, 0x9c, 0x1a, 0x4d,
	0x38, 0xf8, 0x06, 0x15, 0x7b, 0x68, 0xc4, 0x42, 0xb6, 0x68, 0xbc, 0x35, 0x0d, 0x1d, 0xb1, 0x11,
	0xba, 0xeb, 0xba, 0x24, 0x7a, 0xf1, 0xed, 0x7c, 0x2f, 0xb3, 0xdb, 0x6a, 0xb8, 0x78, 0x9b, 0xff,
	0x9b, 0xfc, 0x3b, 0xdf, 0x42, 0x6f, 0xc3, 0xd7, 0x65, 0x35, 0x11, 0x92, 0x2b, 0x4e, 0x2d, 0x13,
	0xc6, 0x5f, 0xd0, 0x9f, 0xe9, 0x21, 0xdc, 0xa9, 0x17, 0x86, 0x6f, 0x74, 0x04, 0x90, 0xaf, 0x56,
	0x7c, 0x57, 0xa9, 0xac, 0x2c, 0x3c, 0xe2, 0x93, 0xa0, 0xcd, 0x9c, 0x86, 0xc4, 0x05, 0xbd, 0x80,
	0xfe, 0xef, 0xba, 0xca, 0xb7, 0xe8, 0x1d, 0xfa, 0x24, 0x70, 0x58, 0xaf, 0x61, 0x49, 0xbe, 0x45,
	0x7a, 0x0a, 0x1d, 0xfc, 0x10, 0xa5, 0x44, 0x6f, 0xe0, 0x93, 0xa0, 0xc5, 0x9a, 0x44, 0x87, 0x60,
	0x29, 0xfe, 0x8a, 0x95, 0x77, 0x64, 0x3a, 0x75, 0x18, 0x7f, 0x93, 0x7d, 0x81, 0x77, 0x41, 0xaf,
	0xc0, 0x46, 0x29, 0xb3, 0x15, 0x2f, 0xd0, 0x9c, 0x1f, 0x5c, 0xbb, 0x93, 0xda, 0x3b, 0x92, 0x92,
	0xcb, 0x29, 0x2f, 0x90, 0x75, 0x51, 0x9a, 0x81, 0x9e, 0x83, 0x23, 0x36, 0xf9, 0x27, 0x4a, 0x2d,
	0xdb, 0x32, 0xb2, 0x76, 0x0d, 0xe2, 0x82, 0x9e, 0x41, 0x77, 0x9d, 0x2b, 0xcc, 0x4a, 0xe1, 0xb5,
	0xcd, 0xc9, 0x8e, 0x8e, 0xb1, 0xd0, 0x2d, 0xb3, 0x10, 0x5c, 0x2a, 0xcf, 0xf2, 0x49, 0x60, 0x31,
	0x5b, 0x83, 0x05, 0x97, 0xea, 0x5f, 0xfd, 0x11, 0x80, 0x29, 0xed, 0xff, 0xc1, 0x3c, 0xb3, 0xd4,
	0xe0, 0x32, 0x05, 0xe7, 0xcf, 0x8f, 0xf6, 0xa0, 0x9b, 0x3e, 0x4c, 0xa7, 0x51, 0x9a, 0xba, 0x07,
	0xb4, 0x0f, 0x76, 0x78, 0x93, 0xcc, 0xd9, 0x7d, 0x38, 0x73, 0x09, 0x1d, 0x82, 0x1b, 0x31, 0x96,
	0x45, 0x4f, 0x8b, 0x98, 0x45, 0xd9, 0x72, 0x7e, 0x17, 0x25, 0x2e, 0xd2, 0x13, 0x38, 0xd6, 0x34,
	0x4e, 0x1e, 0xc3, 0x59, 0x7c, 0xdb, 0xe0, 0xe7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x6e,
	0xa2, 0xb1, 0xb1, 0x01, 0x00, 0x00,
}
