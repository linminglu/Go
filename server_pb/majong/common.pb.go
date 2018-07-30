// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package majong

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

type CardColor int32

const (
	CardColor_ColorWan  CardColor = 0
	CardColor_ColorTong CardColor = 1
	CardColor_ColorTiao CardColor = 2
	CardColor_ColorZi   CardColor = 3
	CardColor_ColorHua  CardColor = 4
)

var CardColor_name = map[int32]string{
	0: "ColorWan",
	1: "ColorTong",
	2: "ColorTiao",
	3: "ColorZi",
	4: "ColorHua",
}
var CardColor_value = map[string]int32{
	"ColorWan":  0,
	"ColorTong": 1,
	"ColorTiao": 2,
	"ColorZi":   3,
	"ColorHua":  4,
}

func (x CardColor) String() string {
	return proto.EnumName(CardColor_name, int32(x))
}
func (CardColor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_bccb12901ab3a744, []int{0}
}

// TingType 听牌类型
type TingType int32

const (
	TingType_TT_NORMAL_TING TingType = 0
	TingType_TT_TIAN_TING   TingType = 1
)

var TingType_name = map[int32]string{
	0: "TT_NORMAL_TING",
	1: "TT_TIAN_TING",
}
var TingType_value = map[string]int32{
	"TT_NORMAL_TING": 0,
	"TT_TIAN_TING":   1,
}

func (x TingType) String() string {
	return proto.EnumName(TingType_name, int32(x))
}
func (TingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_bccb12901ab3a744, []int{1}
}

// Card 卡牌结构
type Card struct {
	Color                CardColor `protobuf:"varint,1,opt,name=color,proto3,enum=majong.CardColor" json:"color,omitempty"`
	Point                int32     `protobuf:"varint,2,opt,name=point,proto3" json:"point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_bccb12901ab3a744, []int{0}
}
func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (dst *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(dst, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetColor() CardColor {
	if m != nil {
		return m.Color
	}
	return CardColor_ColorWan
}

func (m *Card) GetPoint() int32 {
	if m != nil {
		return m.Point
	}
	return 0
}

// TingAction 听牌动作
type TingAction struct {
	EnableTing           bool     `protobuf:"varint,1,opt,name=enable_ting,json=enableTing,proto3" json:"enable_ting,omitempty"`
	TingType             TingType `protobuf:"varint,2,opt,name=ting_type,json=tingType,proto3,enum=majong.TingType" json:"ting_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TingAction) Reset()         { *m = TingAction{} }
func (m *TingAction) String() string { return proto.CompactTextString(m) }
func (*TingAction) ProtoMessage()    {}
func (*TingAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_bccb12901ab3a744, []int{1}
}
func (m *TingAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TingAction.Unmarshal(m, b)
}
func (m *TingAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TingAction.Marshal(b, m, deterministic)
}
func (dst *TingAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TingAction.Merge(dst, src)
}
func (m *TingAction) XXX_Size() int {
	return xxx_messageInfo_TingAction.Size(m)
}
func (m *TingAction) XXX_DiscardUnknown() {
	xxx_messageInfo_TingAction.DiscardUnknown(m)
}

var xxx_messageInfo_TingAction proto.InternalMessageInfo

func (m *TingAction) GetEnableTing() bool {
	if m != nil {
		return m.EnableTing
	}
	return false
}

func (m *TingAction) GetTingType() TingType {
	if m != nil {
		return m.TingType
	}
	return TingType_TT_NORMAL_TING
}

func init() {
	proto.RegisterType((*Card)(nil), "majong.Card")
	proto.RegisterType((*TingAction)(nil), "majong.TingAction")
	proto.RegisterEnum("majong.CardColor", CardColor_name, CardColor_value)
	proto.RegisterEnum("majong.TingType", TingType_name, TingType_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_bccb12901ab3a744) }

var fileDescriptor_common_bccb12901ab3a744 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x6d, 0x6a, 0x3a, 0x8d, 0x61, 0x1d, 0x3c, 0xf4, 0xa6, 0xf4, 0xa2, 0x14, 0x0c,
	0x52, 0x3f, 0x41, 0x28, 0xa2, 0x05, 0x8d, 0xb0, 0x2c, 0x08, 0x22, 0x84, 0x6d, 0x0c, 0x65, 0xa5,
	0xd9, 0x09, 0x61, 0x3d, 0xf8, 0xed, 0xdd, 0x3f, 0x36, 0xbd, 0xcd, 0x7b, 0xbf, 0x37, 0x6f, 0x96,
	0x85, 0xb4, 0xa6, 0xb6, 0x25, 0x9d, 0x77, 0x3d, 0x19, 0xc2, 0x49, 0x2b, 0xbf, 0x49, 0xef, 0x16,
	0x8f, 0x30, 0x5e, 0xcb, 0xfe, 0x0b, 0x6f, 0x20, 0xae, 0x69, 0x4f, 0xfd, 0x3c, 0xba, 0x8e, 0x6e,
	0xb3, 0xd5, 0x45, 0x1e, 0x78, 0xee, 0xe0, 0xda, 0x01, 0x1e, 0x38, 0x5e, 0x42, 0xdc, 0x91, 0xd2,
	0x66, 0x3e, 0xb2, 0xc1, 0x98, 0x07, 0xb1, 0xf8, 0x04, 0x10, 0x4a, 0xef, 0x8a, 0xda, 0x28, 0xd2,
	0x78, 0x05, 0xb3, 0x46, 0xcb, 0xed, 0xbe, 0xa9, 0x8c, 0x35, 0x7d, 0x65, 0xc2, 0x21, 0x58, 0x2e,
	0x86, 0x77, 0x30, 0x75, 0xa4, 0x32, 0xbf, 0x5d, 0xe3, 0x8b, 0xb2, 0x15, 0x3b, 0x5c, 0x74, 0x01,
	0x61, 0x7d, 0x9e, 0x98, 0xff, 0x69, 0xc9, 0x61, 0x3a, 0xbc, 0x03, 0x53, 0x48, 0xfc, 0xf0, 0x2e,
	0x35, 0x3b, 0xc1, 0x73, 0x8b, 0x9c, 0x12, 0x76, 0x95, 0x45, 0x47, 0xa9, 0x24, 0xb1, 0x11, 0xce,
	0xe0, 0xcc, 0xcb, 0x0f, 0xc5, 0x4e, 0x87, 0xc5, 0xe7, 0x1f, 0xc9, 0xc6, 0xcb, 0x7b, 0x48, 0x0e,
	0x97, 0x10, 0x21, 0x13, 0xa2, 0x2a, 0xdf, 0xf8, 0x6b, 0xf1, 0x52, 0x89, 0x4d, 0xf9, 0x64, 0x8b,
	0x19, 0xa4, 0xd6, 0x13, 0x9b, 0xa2, 0x0c, 0x4e, 0xb4, 0x9d, 0xf8, 0x9f, 0x7b, 0xf8, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0x69, 0xbf, 0xf2, 0x49, 0x01, 0x00, 0x00,
}
