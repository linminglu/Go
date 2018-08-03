// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hall.proto

/*
Package hall is a generated protocol buffer package.

It is generated from these files:
	hall.proto

It has these top-level messages:
	HallGetPlayerInfoReq
	HallGetPlayerInfoRsp
	HallGetPlayerStateReq
	HallGetPlayerStateRsp
	HallGetGameListInfoReq
	HallGetGameListInfoRsp
*/
package hall

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "steve/client_pb/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HallGetPlayerInfoReq 获取玩家信息请求
type HallGetPlayerInfoReq struct {
	Reserve          *int32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HallGetPlayerInfoReq) Reset()                    { *m = HallGetPlayerInfoReq{} }
func (m *HallGetPlayerInfoReq) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerInfoReq) ProtoMessage()               {}
func (*HallGetPlayerInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HallGetPlayerInfoReq) GetReserve() int32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// HallGetPlayerInfoRsp 获取玩家信息应答
type HallGetPlayerInfoRsp struct {
	ErrCode          *uint32             `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrDesc          *string             `protobuf:"bytes,2,opt,name=err_desc,json=errDesc" json:"err_desc,omitempty"`
	NickName         *string             `protobuf:"bytes,3,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Coin             *uint64             `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	PlayerState      *common.PlayerState `protobuf:"varint,5,opt,name=player_state,json=playerState,enum=common.PlayerState" json:"player_state,omitempty"`
	GameId           *common.GameId      `protobuf:"varint,6,opt,name=game_id,json=gameId,enum=common.GameId" json:"game_id,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *HallGetPlayerInfoRsp) Reset()                    { *m = HallGetPlayerInfoRsp{} }
func (m *HallGetPlayerInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerInfoRsp) ProtoMessage()               {}
func (*HallGetPlayerInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HallGetPlayerInfoRsp) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetErrDesc() string {
	if m != nil && m.ErrDesc != nil {
		return *m.ErrDesc
	}
	return ""
}

func (m *HallGetPlayerInfoRsp) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *HallGetPlayerInfoRsp) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetPlayerState() common.PlayerState {
	if m != nil && m.PlayerState != nil {
		return *m.PlayerState
	}
	return common.PlayerState_PS_IDLE
}

func (m *HallGetPlayerInfoRsp) GetGameId() common.GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return common.GameId_GAMEID_XUELIU
}

// HallGetPlayerStateReq 获取玩家当前状态请求
type HallGetPlayerStateReq struct {
	Reserve          *int32  `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	UserData         *uint32 `protobuf:"varint,2,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallGetPlayerStateReq) Reset()                    { *m = HallGetPlayerStateReq{} }
func (m *HallGetPlayerStateReq) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerStateReq) ProtoMessage()               {}
func (*HallGetPlayerStateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HallGetPlayerStateReq) GetReserve() int32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

func (m *HallGetPlayerStateReq) GetUserData() uint32 {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return 0
}

// HallGetPlayerStateRsp 获取玩家当前状态响应
type HallGetPlayerStateRsp struct {
	PlayerState      *common.PlayerState `protobuf:"varint,1,opt,name=player_state,json=playerState,enum=common.PlayerState" json:"player_state,omitempty"`
	GameId           *common.GameId      `protobuf:"varint,6,opt,name=game_id,json=gameId,enum=common.GameId" json:"game_id,omitempty"`
	UserData         *uint32             `protobuf:"varint,2,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *HallGetPlayerStateRsp) Reset()                    { *m = HallGetPlayerStateRsp{} }
func (m *HallGetPlayerStateRsp) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerStateRsp) ProtoMessage()               {}
func (*HallGetPlayerStateRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HallGetPlayerStateRsp) GetPlayerState() common.PlayerState {
	if m != nil && m.PlayerState != nil {
		return *m.PlayerState
	}
	return common.PlayerState_PS_IDLE
}

func (m *HallGetPlayerStateRsp) GetGameId() common.GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return common.GameId_GAMEID_XUELIU
}

func (m *HallGetPlayerStateRsp) GetUserData() uint32 {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return 0
}

// HallGetGameListInfoReq 获取游戏列表信息请求
type HallGetGameListInfoReq struct {
	Reserve          *int32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HallGetGameListInfoReq) Reset()                    { *m = HallGetGameListInfoReq{} }
func (m *HallGetGameListInfoReq) String() string            { return proto.CompactTextString(m) }
func (*HallGetGameListInfoReq) ProtoMessage()               {}
func (*HallGetGameListInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HallGetGameListInfoReq) GetReserve() int32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// HallGetGameListInfoRsp 获取游戏列表信息响应
type HallGetGameListInfoRsp struct {
	ErrCode          *uint32                   `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	GameConfig       []*common.GameConfig      `protobuf:"bytes,2,rep,name=game_config,json=gameConfig" json:"game_config,omitempty"`
	GameLevelConfig  []*common.GameLevelConfig `protobuf:"bytes,3,rep,name=game_level_config,json=gameLevelConfig" json:"game_level_config,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *HallGetGameListInfoRsp) Reset()                    { *m = HallGetGameListInfoRsp{} }
func (m *HallGetGameListInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*HallGetGameListInfoRsp) ProtoMessage()               {}
func (*HallGetGameListInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HallGetGameListInfoRsp) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *HallGetGameListInfoRsp) GetGameConfig() []*common.GameConfig {
	if m != nil {
		return m.GameConfig
	}
	return nil
}

func (m *HallGetGameListInfoRsp) GetGameLevelConfig() []*common.GameLevelConfig {
	if m != nil {
		return m.GameLevelConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*HallGetPlayerInfoReq)(nil), "hall.HallGetPlayerInfoReq")
	proto.RegisterType((*HallGetPlayerInfoRsp)(nil), "hall.HallGetPlayerInfoRsp")
	proto.RegisterType((*HallGetPlayerStateReq)(nil), "hall.HallGetPlayerStateReq")
	proto.RegisterType((*HallGetPlayerStateRsp)(nil), "hall.HallGetPlayerStateRsp")
	proto.RegisterType((*HallGetGameListInfoReq)(nil), "hall.HallGetGameListInfoReq")
	proto.RegisterType((*HallGetGameListInfoRsp)(nil), "hall.HallGetGameListInfoRsp")
}

func init() { proto.RegisterFile("hall.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4d, 0xae, 0xd3, 0x30,
	0x14, 0x85, 0xe5, 0x36, 0xfd, 0x73, 0x7f, 0x10, 0xa6, 0x14, 0x03, 0x93, 0x28, 0x13, 0x32, 0x6a,
	0x51, 0x91, 0x58, 0x00, 0xad, 0x54, 0x2a, 0x55, 0x15, 0x32, 0x33, 0x26, 0x91, 0x71, 0x6e, 0x43,
	0x84, 0x13, 0x07, 0xdb, 0x54, 0x62, 0x21, 0x2c, 0x83, 0x2d, 0xb1, 0x16, 0x64, 0xa7, 0xe1, 0xb5,
	0x7a, 0x55, 0xdf, 0xe4, 0x8d, 0x72, 0x7d, 0x8e, 0x3e, 0xe7, 0x9c, 0x5c, 0x05, 0xe3, 0x6f, 0x5c,
	0xca, 0x79, 0xa5, 0x95, 0x55, 0x24, 0x70, 0xf3, 0xab, 0x91, 0x50, 0x45, 0xa1, 0xca, 0x5a, 0x8b,
	0xde, 0xe2, 0xe9, 0x47, 0x2e, 0xe5, 0x06, 0xec, 0x27, 0xc9, 0x7f, 0x81, 0xde, 0x96, 0x07, 0xc5,
	0xe0, 0x07, 0xa1, 0xb8, 0xa7, 0xc1, 0x80, 0x3e, 0x02, 0x45, 0x21, 0x8a, 0x3b, 0xac, 0x39, 0x46,
	0x7f, 0xd1, 0x35, 0xc4, 0x54, 0xe4, 0x25, 0xee, 0x83, 0xd6, 0x89, 0x50, 0x69, 0xcd, 0x8c, 0x59,
	0x0f, 0xb4, 0x5e, 0xa9, 0x14, 0x1a, 0x2b, 0x05, 0x23, 0x68, 0x2b, 0x44, 0xf1, 0xc0, 0x5b, 0x6b,
	0x30, 0x82, 0xbc, 0xc6, 0x83, 0x32, 0x17, 0xdf, 0x93, 0x92, 0x17, 0x40, 0xdb, 0xde, 0xeb, 0x3b,
	0x61, 0xcf, 0x0b, 0x20, 0x04, 0x07, 0x42, 0xe5, 0x25, 0x0d, 0x42, 0x14, 0x07, 0xcc, 0xcf, 0xe4,
	0x3d, 0x1e, 0x55, 0xfe, 0xbd, 0x89, 0xb1, 0xdc, 0x02, 0xed, 0x84, 0x28, 0x9e, 0x2c, 0x9f, 0xcd,
	0x4f, 0xb5, 0xea, 0x4c, 0x9f, 0x9d, 0xc5, 0x86, 0xd5, 0xdd, 0x81, 0xbc, 0xc1, 0xbd, 0x8c, 0x17,
	0x90, 0xe4, 0x29, 0xed, 0x7a, 0x64, 0xd2, 0x20, 0x1b, 0x5e, 0xc0, 0x36, 0x65, 0xdd, 0xcc, 0x3f,
	0xa3, 0x3d, 0x7e, 0x7e, 0xd1, 0xaf, 0xbe, 0xeb, 0xd6, 0x37, 0x71, 0x25, 0x7e, 0x1a, 0xd0, 0x49,
	0xca, 0x2d, 0xf7, 0x05, 0xc7, 0xac, 0xef, 0x84, 0x35, 0xb7, 0x3c, 0xfa, 0x8d, 0xae, 0x5e, 0x68,
	0xaa, 0x7b, 0x55, 0xd0, 0x23, 0x57, 0xb9, 0x9d, 0x6b, 0x89, 0x67, 0xa7, 0x58, 0x8e, 0xda, 0xe5,
	0xc6, 0x3e, 0xbc, 0xfc, 0x3f, 0xe8, 0x3a, 0x74, 0x7b, 0xfd, 0xef, 0xf0, 0xd0, 0xe7, 0x15, 0xaa,
	0x3c, 0xe4, 0x19, 0x6d, 0x85, 0xed, 0x78, 0xb8, 0x24, 0xe7, 0x99, 0x57, 0xde, 0x61, 0x38, 0xfb,
	0x3f, 0x93, 0x15, 0x7e, 0xea, 0x21, 0x09, 0x47, 0x90, 0x0d, 0xda, 0xf6, 0xe8, 0x8b, 0x73, 0x74,
	0xe7, 0xfc, 0x13, 0xff, 0x24, 0xbb, 0x14, 0x3e, 0xcc, 0xbe, 0x4c, 0x8d, 0x85, 0x23, 0x2c, 0x84,
	0xcc, 0xa1, 0xb4, 0x49, 0xf5, 0x75, 0xe1, 0x7e, 0x82, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xd7, 0x15, 0x1b, 0x17, 0x03, 0x00, 0x00,
}
