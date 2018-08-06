// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgserver.proto

package msgserver // import "steve/client_pb/msgserver"

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

// msgid:MSGSVR_GET_HORSE_RACE_REQ = 0x4001; 	// 获取跑马灯请求
// 获取跑马灯请求
type MsgSvrGetHorseRaceReq struct {
	PlayerId             *uint64  `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSvrGetHorseRaceReq) Reset()         { *m = MsgSvrGetHorseRaceReq{} }
func (m *MsgSvrGetHorseRaceReq) String() string { return proto.CompactTextString(m) }
func (*MsgSvrGetHorseRaceReq) ProtoMessage()    {}
func (*MsgSvrGetHorseRaceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgserver_bfcdafb8bc5cce3a, []int{0}
}
func (m *MsgSvrGetHorseRaceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSvrGetHorseRaceReq.Unmarshal(m, b)
}
func (m *MsgSvrGetHorseRaceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSvrGetHorseRaceReq.Marshal(b, m, deterministic)
}
func (dst *MsgSvrGetHorseRaceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSvrGetHorseRaceReq.Merge(dst, src)
}
func (m *MsgSvrGetHorseRaceReq) XXX_Size() int {
	return xxx_messageInfo_MsgSvrGetHorseRaceReq.Size(m)
}
func (m *MsgSvrGetHorseRaceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSvrGetHorseRaceReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSvrGetHorseRaceReq proto.InternalMessageInfo

func (m *MsgSvrGetHorseRaceReq) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

// msgid:MSGSVR_GET_HORSE_RACE_RSP = 0x4002; 	// 获取跑马灯应答
// 获取跑马灯应答
type MsgSvrGetHorseRaceRsp struct {
	ErrCode              *int32   `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrDesc              *string  `protobuf:"bytes,2,opt,name=err_desc,json=errDesc" json:"err_desc,omitempty"`
	Tick                 *int32   `protobuf:"varint,3,opt,name=tick" json:"tick,omitempty"`
	Sleep                *int32   `protobuf:"varint,4,opt,name=sleep" json:"sleep,omitempty"`
	Content              []string `protobuf:"bytes,5,rep,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSvrGetHorseRaceRsp) Reset()         { *m = MsgSvrGetHorseRaceRsp{} }
func (m *MsgSvrGetHorseRaceRsp) String() string { return proto.CompactTextString(m) }
func (*MsgSvrGetHorseRaceRsp) ProtoMessage()    {}
func (*MsgSvrGetHorseRaceRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgserver_bfcdafb8bc5cce3a, []int{1}
}
func (m *MsgSvrGetHorseRaceRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSvrGetHorseRaceRsp.Unmarshal(m, b)
}
func (m *MsgSvrGetHorseRaceRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSvrGetHorseRaceRsp.Marshal(b, m, deterministic)
}
func (dst *MsgSvrGetHorseRaceRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSvrGetHorseRaceRsp.Merge(dst, src)
}
func (m *MsgSvrGetHorseRaceRsp) XXX_Size() int {
	return xxx_messageInfo_MsgSvrGetHorseRaceRsp.Size(m)
}
func (m *MsgSvrGetHorseRaceRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSvrGetHorseRaceRsp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSvrGetHorseRaceRsp proto.InternalMessageInfo

func (m *MsgSvrGetHorseRaceRsp) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *MsgSvrGetHorseRaceRsp) GetErrDesc() string {
	if m != nil && m.ErrDesc != nil {
		return *m.ErrDesc
	}
	return ""
}

func (m *MsgSvrGetHorseRaceRsp) GetTick() int32 {
	if m != nil && m.Tick != nil {
		return *m.Tick
	}
	return 0
}

func (m *MsgSvrGetHorseRaceRsp) GetSleep() int32 {
	if m != nil && m.Sleep != nil {
		return *m.Sleep
	}
	return 0
}

func (m *MsgSvrGetHorseRaceRsp) GetContent() []string {
	if m != nil {
		return m.Content
	}
	return nil
}

// msgid:MSGSVR_HORSE_RACE_UPDATE_NTF = 0x4003; // 跑马灯变化通知
// 跑马灯变化通知
type MsgSvrHorseRaceChangeNtf struct {
	Channel              *int32   `protobuf:"varint,1,opt,name=channel" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSvrHorseRaceChangeNtf) Reset()         { *m = MsgSvrHorseRaceChangeNtf{} }
func (m *MsgSvrHorseRaceChangeNtf) String() string { return proto.CompactTextString(m) }
func (*MsgSvrHorseRaceChangeNtf) ProtoMessage()    {}
func (*MsgSvrHorseRaceChangeNtf) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgserver_bfcdafb8bc5cce3a, []int{2}
}
func (m *MsgSvrHorseRaceChangeNtf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSvrHorseRaceChangeNtf.Unmarshal(m, b)
}
func (m *MsgSvrHorseRaceChangeNtf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSvrHorseRaceChangeNtf.Marshal(b, m, deterministic)
}
func (dst *MsgSvrHorseRaceChangeNtf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSvrHorseRaceChangeNtf.Merge(dst, src)
}
func (m *MsgSvrHorseRaceChangeNtf) XXX_Size() int {
	return xxx_messageInfo_MsgSvrHorseRaceChangeNtf.Size(m)
}
func (m *MsgSvrHorseRaceChangeNtf) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSvrHorseRaceChangeNtf.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSvrHorseRaceChangeNtf proto.InternalMessageInfo

func (m *MsgSvrHorseRaceChangeNtf) GetChannel() int32 {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgSvrGetHorseRaceReq)(nil), "msgserver.MsgSvrGetHorseRaceReq")
	proto.RegisterType((*MsgSvrGetHorseRaceRsp)(nil), "msgserver.MsgSvrGetHorseRaceRsp")
	proto.RegisterType((*MsgSvrHorseRaceChangeNtf)(nil), "msgserver.MsgSvrHorseRaceChangeNtf")
}

func init() { proto.RegisterFile("msgserver.proto", fileDescriptor_msgserver_bfcdafb8bc5cce3a) }

var fileDescriptor_msgserver_bfcdafb8bc5cce3a = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x15, 0x9a, 0xa8, 0x8d, 0x17, 0x24, 0x0b, 0x24, 0x57, 0x5d, 0xa2, 0x4c, 0x9d, 0xe8,
	0xd2, 0x2f, 0xa0, 0x48, 0xc0, 0x00, 0x43, 0xd8, 0x58, 0xa2, 0x60, 0x5f, 0xd2, 0x08, 0x63, 0x9b,
	0xf7, 0xac, 0x48, 0x7c, 0x06, 0x7f, 0x8c, 0x70, 0xda, 0x4c, 0x6c, 0xbe, 0xf7, 0xea, 0x58, 0x47,
	0x4f, 0x5c, 0x7e, 0x72, 0xcf, 0xa0, 0x11, 0x74, 0x13, 0xc8, 0x47, 0x2f, 0xcb, 0xb9, 0xa8, 0xf7,
	0xe2, 0xfa, 0x89, 0xfb, 0x97, 0x91, 0xee, 0x11, 0x1f, 0x3c, 0x31, 0x9a, 0x4e, 0xa3, 0xc1, 0x97,
	0xdc, 0x88, 0x32, 0xd8, 0xee, 0x1b, 0xd4, 0x0e, 0x46, 0x65, 0x55, 0xb6, 0xcd, 0x9b, 0xd5, 0x54,
	0x3c, 0x9a, 0xfa, 0x27, 0xfb, 0x17, 0xe3, 0x20, 0xd7, 0x62, 0x05, 0xa2, 0x56, 0x7b, 0x83, 0x44,
	0x15, 0xcd, 0x12, 0x44, 0x07, 0x6f, 0x70, 0x9e, 0x0c, 0x58, 0xab, 0x8b, 0x2a, 0xdb, 0x96, 0x69,
	0xba, 0x03, 0x6b, 0x29, 0x45, 0x1e, 0x07, 0xfd, 0xa1, 0x16, 0x89, 0x48, 0x6f, 0x79, 0x25, 0x0a,
	0xb6, 0x40, 0x50, 0x79, 0x2a, 0xa7, 0x20, 0x95, 0x58, 0x6a, 0xef, 0x22, 0x5c, 0x54, 0x45, 0xb5,
	0xf8, 0xfb, 0xe3, 0x14, 0xeb, 0xbd, 0x50, 0x93, 0xd2, 0xec, 0x73, 0x38, 0x76, 0xae, 0xc7, 0x73,
	0x7c, 0x4f, 0xd4, 0xb1, 0x73, 0x0e, 0xf6, 0x2c, 0x75, 0x8a, 0xb7, 0x9b, 0xd7, 0x35, 0x47, 0x8c,
	0xd8, 0x69, 0x3b, 0xc0, 0xc5, 0x36, 0xbc, 0xed, 0xe6, 0xe3, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xff, 0x9d, 0xf8, 0xee, 0x39, 0x01, 0x00, 0x00,
}
