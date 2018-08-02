// Code generated by protoc-gen-go. DO NOT EDIT.
// source: robot.proto

package robot

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GameConfig 游戏玩法信息
type GameConfig struct {
	GameId  uint32 `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	LevelId uint32 `protobuf:"varint,2,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
}

func (m *GameConfig) Reset()                    { *m = GameConfig{} }
func (m *GameConfig) String() string            { return proto.CompactTextString(m) }
func (*GameConfig) ProtoMessage()               {}
func (*GameConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GameConfig) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *GameConfig) GetLevelId() uint32 {
	if m != nil {
		return m.LevelId
	}
	return 0
}

// GameWinRate 游戏对应的胜率
type GameWinRate struct {
	Game    *GameConfig `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
	WinRate int32       `protobuf:"varint,2,opt,name=win_rate,json=winRate" json:"win_rate,omitempty"`
}

func (m *GameWinRate) Reset()                    { *m = GameWinRate{} }
func (m *GameWinRate) String() string            { return proto.CompactTextString(m) }
func (*GameWinRate) ProtoMessage()               {}
func (*GameWinRate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GameWinRate) GetGame() *GameConfig {
	if m != nil {
		return m.Game
	}
	return nil
}

func (m *GameWinRate) GetWinRate() int32 {
	if m != nil {
		return m.WinRate
	}
	return 0
}

// RobotPlayerInfo 机器人玩家信息
type RobotPlayerInfo struct {
	NickName    string           `protobuf:"bytes,1,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Avatar      string           `protobuf:"bytes,2,opt,name=avatar" json:"avatar,omitempty"`
	Coin        uint64           `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	State       RobotPlayerState `protobuf:"varint,4,opt,name=state,enum=robot.RobotPlayerState" json:"state,omitempty"`
	GameWinRate *GameWinRate     `protobuf:"bytes,5,opt,name=game_win_rate,json=gameWinRate" json:"game_win_rate,omitempty"`
}

func (m *RobotPlayerInfo) Reset()                    { *m = RobotPlayerInfo{} }
func (m *RobotPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*RobotPlayerInfo) ProtoMessage()               {}
func (*RobotPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *RobotPlayerInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *RobotPlayerInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *RobotPlayerInfo) GetCoin() uint64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

func (m *RobotPlayerInfo) GetState() RobotPlayerState {
	if m != nil {
		return m.State
	}
	return RobotPlayerState_RPS_IDIE
}

func (m *RobotPlayerInfo) GetGameWinRate() *GameWinRate {
	if m != nil {
		return m.GameWinRate
	}
	return nil
}

// WinRateRange 胜率范围
type WinRateRange struct {
	High int32 `protobuf:"varint,1,opt,name=high" json:"high,omitempty"`
	Low  int32 `protobuf:"varint,2,opt,name=low" json:"low,omitempty"`
}

func (m *WinRateRange) Reset()                    { *m = WinRateRange{} }
func (m *WinRateRange) String() string            { return proto.CompactTextString(m) }
func (*WinRateRange) ProtoMessage()               {}
func (*WinRateRange) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *WinRateRange) GetHigh() int32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *WinRateRange) GetLow() int32 {
	if m != nil {
		return m.Low
	}
	return 0
}

// CoinsRange 金币范围
type CoinsRange struct {
	High int64 `protobuf:"varint,1,opt,name=high" json:"high,omitempty"`
	Low  int64 `protobuf:"varint,2,opt,name=low" json:"low,omitempty"`
}

func (m *CoinsRange) Reset()                    { *m = CoinsRange{} }
func (m *CoinsRange) String() string            { return proto.CompactTextString(m) }
func (*CoinsRange) ProtoMessage()               {}
func (*CoinsRange) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CoinsRange) GetHigh() int64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *CoinsRange) GetLow() int64 {
	if m != nil {
		return m.Low
	}
	return 0
}

// GetRobotPlayerIDReq 获取机器人ID请求
type GetRobotPlayerIDReq struct {
	Game         *GameConfig      `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
	WinRateRange *WinRateRange    `protobuf:"bytes,2,opt,name=win_rate_range,json=winRateRange" json:"win_rate_range,omitempty"`
	CoinsRange   *CoinsRange      `protobuf:"bytes,3,opt,name=coins_range,json=coinsRange" json:"coins_range,omitempty"`
	NewState     RobotPlayerState `protobuf:"varint,4,opt,name=new_state,json=newState,enum=robot.RobotPlayerState" json:"new_state,omitempty"`
}

func (m *GetRobotPlayerIDReq) Reset()                    { *m = GetRobotPlayerIDReq{} }
func (m *GetRobotPlayerIDReq) String() string            { return proto.CompactTextString(m) }
func (*GetRobotPlayerIDReq) ProtoMessage()               {}
func (*GetRobotPlayerIDReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetRobotPlayerIDReq) GetGame() *GameConfig {
	if m != nil {
		return m.Game
	}
	return nil
}

func (m *GetRobotPlayerIDReq) GetWinRateRange() *WinRateRange {
	if m != nil {
		return m.WinRateRange
	}
	return nil
}

func (m *GetRobotPlayerIDReq) GetCoinsRange() *CoinsRange {
	if m != nil {
		return m.CoinsRange
	}
	return nil
}

func (m *GetRobotPlayerIDReq) GetNewState() RobotPlayerState {
	if m != nil {
		return m.NewState
	}
	return RobotPlayerState_RPS_IDIE
}

// GetRobotPlayerIDRsp 机器人ID响应
type GetRobotPlayerIDRsp struct {
	RobotPlayerId uint64 `protobuf:"varint,1,opt,name=robot_player_id,json=robotPlayerId" json:"robot_player_id,omitempty"`
	ErrCode       int32  `protobuf:"varint,2,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
}

func (m *GetRobotPlayerIDRsp) Reset()                    { *m = GetRobotPlayerIDRsp{} }
func (m *GetRobotPlayerIDRsp) String() string            { return proto.CompactTextString(m) }
func (*GetRobotPlayerIDRsp) ProtoMessage()               {}
func (*GetRobotPlayerIDRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *GetRobotPlayerIDRsp) GetRobotPlayerId() uint64 {
	if m != nil {
		return m.RobotPlayerId
	}
	return 0
}

func (m *GetRobotPlayerIDRsp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

// SetRobotPlayerStateReq 設置机器人玩家状态請求
type SetRobotPlayerStateReq struct {
	RobotPlayerId uint64           `protobuf:"varint,1,opt,name=robot_player_id,json=robotPlayerId" json:"robot_player_id,omitempty"`
	NewState      RobotPlayerState `protobuf:"varint,2,opt,name=new_state,json=newState,enum=robot.RobotPlayerState" json:"new_state,omitempty"`
	OldState      RobotPlayerState `protobuf:"varint,3,opt,name=old_state,json=oldState,enum=robot.RobotPlayerState" json:"old_state,omitempty"`
	ServerType    ServerType       `protobuf:"varint,4,opt,name=server_type,json=serverType,enum=robot.ServerType" json:"server_type,omitempty"`
	ServerAddr    string           `protobuf:"bytes,5,opt,name=server_addr,json=serverAddr" json:"server_addr,omitempty"`
}

func (m *SetRobotPlayerStateReq) Reset()                    { *m = SetRobotPlayerStateReq{} }
func (m *SetRobotPlayerStateReq) String() string            { return proto.CompactTextString(m) }
func (*SetRobotPlayerStateReq) ProtoMessage()               {}
func (*SetRobotPlayerStateReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *SetRobotPlayerStateReq) GetRobotPlayerId() uint64 {
	if m != nil {
		return m.RobotPlayerId
	}
	return 0
}

func (m *SetRobotPlayerStateReq) GetNewState() RobotPlayerState {
	if m != nil {
		return m.NewState
	}
	return RobotPlayerState_RPS_IDIE
}

func (m *SetRobotPlayerStateReq) GetOldState() RobotPlayerState {
	if m != nil {
		return m.OldState
	}
	return RobotPlayerState_RPS_IDIE
}

func (m *SetRobotPlayerStateReq) GetServerType() ServerType {
	if m != nil {
		return m.ServerType
	}
	return ServerType_ST_GATE
}

func (m *SetRobotPlayerStateReq) GetServerAddr() string {
	if m != nil {
		return m.ServerAddr
	}
	return ""
}

// SetRobotPlayerStateRsp 設置机器人玩家状态响应
type SetRobotPlayerStateRsp struct {
	Result  bool  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	ErrCode int32 `protobuf:"varint,2,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
}

func (m *SetRobotPlayerStateRsp) Reset()                    { *m = SetRobotPlayerStateRsp{} }
func (m *SetRobotPlayerStateRsp) String() string            { return proto.CompactTextString(m) }
func (*SetRobotPlayerStateRsp) ProtoMessage()               {}
func (*SetRobotPlayerStateRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *SetRobotPlayerStateRsp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *SetRobotPlayerStateRsp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func init() {
	proto.RegisterType((*GameConfig)(nil), "robot.GameConfig")
	proto.RegisterType((*GameWinRate)(nil), "robot.GameWinRate")
	proto.RegisterType((*RobotPlayerInfo)(nil), "robot.RobotPlayerInfo")
	proto.RegisterType((*WinRateRange)(nil), "robot.WinRateRange")
	proto.RegisterType((*CoinsRange)(nil), "robot.CoinsRange")
	proto.RegisterType((*GetRobotPlayerIDReq)(nil), "robot.GetRobotPlayerIDReq")
	proto.RegisterType((*GetRobotPlayerIDRsp)(nil), "robot.GetRobotPlayerIDRsp")
	proto.RegisterType((*SetRobotPlayerStateReq)(nil), "robot.SetRobotPlayerStateReq")
	proto.RegisterType((*SetRobotPlayerStateRsp)(nil), "robot.SetRobotPlayerStateRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RobotService service

type RobotServiceClient interface {
	GetRobotPlayerIDByInfo(ctx context.Context, in *GetRobotPlayerIDReq, opts ...grpc.CallOption) (*GetRobotPlayerIDRsp, error)
	SetRobotPlayerState(ctx context.Context, in *SetRobotPlayerStateReq, opts ...grpc.CallOption) (*SetRobotPlayerStateRsp, error)
}

type robotServiceClient struct {
	cc *grpc.ClientConn
}

func NewRobotServiceClient(cc *grpc.ClientConn) RobotServiceClient {
	return &robotServiceClient{cc}
}

func (c *robotServiceClient) GetRobotPlayerIDByInfo(ctx context.Context, in *GetRobotPlayerIDReq, opts ...grpc.CallOption) (*GetRobotPlayerIDRsp, error) {
	out := new(GetRobotPlayerIDRsp)
	err := grpc.Invoke(ctx, "/robot.RobotService/GetRobotPlayerIDByInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) SetRobotPlayerState(ctx context.Context, in *SetRobotPlayerStateReq, opts ...grpc.CallOption) (*SetRobotPlayerStateRsp, error) {
	out := new(SetRobotPlayerStateRsp)
	err := grpc.Invoke(ctx, "/robot.RobotService/SetRobotPlayerState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RobotService service

type RobotServiceServer interface {
	GetRobotPlayerIDByInfo(context.Context, *GetRobotPlayerIDReq) (*GetRobotPlayerIDRsp, error)
	SetRobotPlayerState(context.Context, *SetRobotPlayerStateReq) (*SetRobotPlayerStateRsp, error)
}

func RegisterRobotServiceServer(s *grpc.Server, srv RobotServiceServer) {
	s.RegisterService(&_RobotService_serviceDesc, srv)
}

func _RobotService_GetRobotPlayerIDByInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRobotPlayerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).GetRobotPlayerIDByInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.RobotService/GetRobotPlayerIDByInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).GetRobotPlayerIDByInfo(ctx, req.(*GetRobotPlayerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_SetRobotPlayerState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRobotPlayerStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).SetRobotPlayerState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.RobotService/SetRobotPlayerState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).SetRobotPlayerState(ctx, req.(*SetRobotPlayerStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RobotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "robot.RobotService",
	HandlerType: (*RobotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRobotPlayerIDByInfo",
			Handler:    _RobotService_GetRobotPlayerIDByInfo_Handler,
		},
		{
			MethodName: "SetRobotPlayerState",
			Handler:    _RobotService_SetRobotPlayerState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robot.proto",
}

func init() { proto.RegisterFile("robot.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
<<<<<<< Updated upstream
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x75, 0xd2, 0x24, 0xe3, 0xa4, 0x85, 0x8d, 0x94, 0x96, 0x20, 0x44, 0x64, 0x09, 0x94,
	0x0b, 0x3d, 0x98, 0x0a, 0x89, 0x1b, 0x10, 0xa4, 0x2a, 0x42, 0x02, 0xb4, 0x01, 0xc1, 0xcd, 0x72,
	0xb3, 0xd3, 0xd4, 0xc2, 0xd9, 0x75, 0xd7, 0x6e, 0xac, 0x1c, 0xf9, 0x2c, 0x8e, 0x7c, 0x0a, 0x7f,
	0x82, 0x76, 0x6c, 0x27, 0x0e, 0x09, 0x25, 0xb7, 0x99, 0x9d, 0x7d, 0x6f, 0xe6, 0xbd, 0x1d, 0x1b,
	0x1c, 0xad, 0x2e, 0x55, 0x7a, 0x16, 0x6b, 0x95, 0x2a, 0x56, 0xa7, 0xa4, 0x0f, 0x28, 0x6f, 0xe7,
	0xf9, 0x91, 0xfb, 0x1a, 0xe0, 0x22, 0x98, 0xe3, 0x48, 0xc9, 0xab, 0x70, 0xc6, 0x4e, 0xa0, 0x31,
	0x0b, 0xe6, 0xe8, 0x87, 0xe2, 0xd4, 0x1a, 0x58, 0xc3, 0x0e, 0x3f, 0x34, 0xe9, 0x58, 0xb0, 0x87,
	0xd0, 0x8c, 0x70, 0x81, 0x91, 0xa9, 0x1c, 0x50, 0xa5, 0x41, 0xf9, 0x58, 0xb8, 0x1f, 0xc1, 0x31,
	0x0c, 0x5f, 0x43, 0xc9, 0x83, 0x14, 0xd9, 0x53, 0xa8, 0x19, 0x0c, 0xe1, 0x1d, 0xef, 0xc1, 0x59,
	0xde, 0x7f, 0xdd, 0x83, 0x53, 0xd9, 0x10, 0x66, 0xa1, 0xf4, 0x75, 0x90, 0x22, 0x11, 0xd6, 0x79,
	0x23, 0xcb, 0x19, 0xdc, 0x5f, 0x16, 0x1c, 0x73, 0x83, 0xfa, 0x14, 0x05, 0x4b, 0xd4, 0x63, 0x79,
	0xa5, 0xd8, 0x23, 0x68, 0xc9, 0x70, 0xfa, 0xdd, 0x97, 0x25, 0x75, 0x8b, 0x37, 0xcd, 0xc1, 0x07,
	0xc3, 0xd5, 0x83, 0xc3, 0x60, 0x11, 0xa4, 0x81, 0x26, 0xa6, 0x16, 0x2f, 0x32, 0xc6, 0xa0, 0x36,
	0x55, 0xa1, 0x3c, 0xb5, 0x07, 0xd6, 0xb0, 0xc6, 0x29, 0x66, 0xcf, 0xa1, 0x9e, 0xa4, 0xa6, 0x69,
	0x6d, 0x60, 0x0d, 0x8f, 0xbc, 0x93, 0x62, 0xbe, 0x4a, 0xbf, 0x89, 0x29, 0xf3, 0xfc, 0x16, 0x7b,
	0x09, 0x1d, 0x32, 0x64, 0x35, 0x6b, 0x9d, 0x64, 0xb1, 0x8a, 0xac, 0x42, 0x38, 0x77, 0x66, 0xeb,
	0xc4, 0x3d, 0x87, 0x76, 0x79, 0x1e, 0xc8, 0x19, 0x9a, 0x51, 0xae, 0xc3, 0xd9, 0x35, 0x8d, 0x5e,
	0xe7, 0x14, 0xb3, 0xfb, 0x60, 0x47, 0x2a, 0x2b, 0xd4, 0x9b, 0xd0, 0xf5, 0x00, 0x46, 0x2a, 0x94,
	0xc9, 0x36, 0xc6, 0xde, 0xc6, 0xd8, 0x39, 0xe6, 0xb7, 0x05, 0xdd, 0x0b, 0x4c, 0xab, 0x86, 0xbd,
	0xe3, 0x78, 0xb3, 0xef, 0x3b, 0xbc, 0x82, 0xa3, 0x52, 0x9b, 0xaf, 0x4d, 0x5b, 0xe2, 0x76, 0xbc,
	0x6e, 0x01, 0xa8, 0xaa, 0xe0, 0xed, 0xac, 0xaa, 0xc9, 0x03, 0xc7, 0x58, 0x9a, 0x14, 0x38, 0x7b,
	0xa3, 0xd1, 0x5a, 0x07, 0x87, 0xe9, 0x5a, 0xd3, 0x39, 0xb4, 0x24, 0x66, 0xfe, 0x5e, 0x4f, 0xd0,
	0x94, 0x98, 0x51, 0xe4, 0x7e, 0xdb, 0x21, 0x31, 0x89, 0xd9, 0x33, 0x38, 0x26, 0xa8, 0x1f, 0xd3,
	0x61, 0xb9, 0xb5, 0x35, 0xde, 0xd1, 0x95, 0xab, 0xb4, 0xbc, 0xa8, 0xb5, 0x3f, 0x55, 0x62, 0xb5,
	0x6b, 0xa8, 0xf5, 0x48, 0x09, 0x74, 0x7f, 0x1c, 0x40, 0x6f, 0xb2, 0x41, 0x9d, 0xf7, 0xc6, 0x9b,
	0xbd, 0xd9, 0x37, 0x24, 0x1d, 0xec, 0x29, 0xc9, 0xa0, 0x54, 0x24, 0x0a, 0x94, 0xfd, 0x1f, 0x94,
	0x8a, 0x44, 0x8e, 0xf2, 0xc0, 0x49, 0x50, 0x2f, 0x50, 0xfb, 0xe9, 0x32, 0x2e, 0x0d, 0x2c, 0x2d,
	0x9f, 0x50, 0xe5, 0xf3, 0x32, 0x46, 0x0e, 0xc9, 0x2a, 0x66, 0x4f, 0x56, 0x98, 0x40, 0x08, 0x4d,
	0x0b, 0xdc, 0x2a, 0x2f, 0xbc, 0x11, 0x42, 0xbb, 0xef, 0x77, 0x5b, 0x90, 0xc4, 0xe6, 0xc3, 0xd2,
	0x98, 0xdc, 0x46, 0x29, 0x29, 0x6f, 0xf2, 0x22, 0xbb, 0xc3, 0x50, 0xef, 0xa7, 0x05, 0x6d, 0xa2,
	0x32, 0xd3, 0x84, 0x53, 0x64, 0x1c, 0x7a, 0x7f, 0xbf, 0xdd, 0xdb, 0x25, 0x7d, 0xd3, 0xfd, 0x72,
	0x27, 0xb7, 0xb7, 0xb7, 0xff, 0xcf, 0x5a, 0x12, 0xbb, 0xf7, 0xd8, 0x17, 0xe8, 0xee, 0x98, 0x98,
	0x3d, 0x5e, 0x19, 0xb1, 0xeb, 0x41, 0xfb, 0x77, 0x95, 0x0d, 0xed, 0xe5, 0x21, 0xfd, 0x12, 0x5f,
	0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x89, 0x75, 0x3f, 0xfd, 0x34, 0x05, 0x00, 0x00,
=======
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xcd, 0xa5, 0xc9, 0xb8, 0x49, 0x61, 0x23, 0xa5, 0x21, 0xa8, 0xa2, 0xb2, 0x04, 0xea,
	0x0b, 0x7d, 0x70, 0x11, 0x12, 0x6f, 0x40, 0x90, 0xaa, 0x08, 0x09, 0xd0, 0x06, 0x04, 0x6f, 0x96,
	0x9b, 0x9d, 0x26, 0x16, 0x8e, 0xd7, 0x5d, 0xbb, 0x8d, 0xf2, 0x4d, 0x7c, 0x01, 0x1f, 0xc4, 0x77,
	0x80, 0x66, 0x7c, 0x89, 0x5b, 0x42, 0x28, 0x6f, 0x3b, 0x33, 0x7b, 0xce, 0xcc, 0x39, 0x3b, 0x36,
	0xd8, 0x46, 0x9f, 0xeb, 0xf4, 0x24, 0x36, 0x3a, 0xd5, 0xa2, 0xc1, 0xc1, 0x10, 0x30, 0xba, 0x5a,
	0x64, 0x29, 0xe7, 0x15, 0xc0, 0x99, 0xbf, 0xc0, 0x91, 0x8e, 0x2e, 0x82, 0x99, 0x38, 0x80, 0xdd,
	0x99, 0xbf, 0x40, 0x2f, 0x50, 0x03, 0xeb, 0xc8, 0x3a, 0xee, 0xc8, 0x26, 0x85, 0x63, 0x25, 0x1e,
	0x42, 0x2b, 0xc4, 0x6b, 0x0c, 0xa9, 0xb2, 0xc3, 0x95, 0x5d, 0x8e, 0xc7, 0xca, 0xf9, 0x00, 0x36,
	0x31, 0x7c, 0x09, 0x22, 0xe9, 0xa7, 0x28, 0x9e, 0x40, 0x9d, 0x30, 0x8c, 0xb7, 0xdd, 0x07, 0x27,
	0x59, 0xff, 0x75, 0x0f, 0xc9, 0x65, 0x22, 0x5c, 0x06, 0x91, 0x67, 0xfc, 0x14, 0x99, 0xb0, 0x21,
	0x77, 0x97, 0x19, 0x83, 0xf3, 0xd3, 0x82, 0x7d, 0x49, 0xa8, 0x8f, 0xa1, 0xbf, 0x42, 0x33, 0x8e,
	0x2e, 0xb4, 0x78, 0x04, 0xed, 0x98, 0xa3, 0x62, 0xb4, 0xba, 0x6c, 0x65, 0x89, 0xb1, 0xa2, 0x62,
	0x14, 0x4c, 0xbf, 0x79, 0x11, 0xf5, 0x25, 0xb2, 0xb6, 0x6c, 0x51, 0xe2, 0x3d, 0x35, 0x3a, 0x04,
	0x98, 0xa3, 0xaf, 0xbc, 0x60, 0xe1, 0xcf, 0x70, 0x50, 0xe3, 0x6a, 0x9b, 0x32, 0x63, 0x4a, 0x08,
	0x01, 0xf5, 0xa9, 0x0e, 0xa2, 0x41, 0x9d, 0x39, 0xf9, 0x2c, 0x9e, 0x41, 0x23, 0x49, 0x69, 0xb0,
	0xc6, 0x91, 0x75, 0xdc, 0x75, 0x0f, 0x72, 0x0d, 0x95, 0x99, 0x26, 0x54, 0x96, 0xd9, 0x2d, 0xf1,
	0x02, 0x3a, 0x6c, 0x5a, 0xa9, 0xa7, 0xc9, 0xd2, 0x45, 0x45, 0x7a, 0x6e, 0x8e, 0xb4, 0x67, 0xeb,
	0xc0, 0x79, 0x0e, 0x7b, 0x45, 0xde, 0x8f, 0xb2, 0x51, 0xe6, 0xc1, 0x6c, 0xce, 0xf2, 0x1a, 0x92,
	0xcf, 0xe2, 0x3e, 0xd4, 0x42, 0xbd, 0xcc, 0x1d, 0xa2, 0xa3, 0xe3, 0x02, 0x8c, 0x74, 0x10, 0x25,
	0xff, 0x83, 0xf9, 0x6e, 0x41, 0xef, 0x0c, 0xd3, 0xaa, 0xa9, 0x6f, 0x25, 0x5e, 0xde, 0xf5, 0xad,
	0x5e, 0x42, 0xb7, 0xd0, 0xe6, 0x19, 0x6a, 0xcb, 0xdc, 0xb6, 0xdb, 0xcb, 0x01, 0x55, 0x15, 0x72,
	0x6f, 0x59, 0xd5, 0xe4, 0x82, 0x4d, 0x96, 0x26, 0x39, 0xae, 0x76, 0xa3, 0xd1, 0x5a, 0x87, 0x84,
	0x69, 0x79, 0x76, 0xbe, 0x6e, 0x18, 0x36, 0x89, 0xc5, 0x53, 0xd8, 0x67, 0x98, 0x77, 0x7b, 0x11,
	0x3a, 0xa6, 0x72, 0x95, 0x57, 0x15, 0x8d, 0xf1, 0xa6, 0x5a, 0x95, 0x9b, 0x85, 0xc6, 0x8c, 0xb4,
	0x42, 0xe7, 0x97, 0x05, 0xfd, 0xc9, 0x0d, 0xea, 0xec, 0x21, 0xf1, 0xf2, 0xce, 0xec, 0xa7, 0xd0,
	0x8a, 0x70, 0x99, 0xad, 0xc7, 0xce, 0xf6, 0xf5, 0x28, 0x2f, 0x12, 0x48, 0x87, 0x2a, 0x03, 0xd5,
	0xfe, 0x01, 0x2a, 0x2e, 0x92, 0x75, 0x09, 0x9a, 0x6b, 0x34, 0x5e, 0xba, 0x8a, 0x91, 0x17, 0xb4,
	0x5b, 0x5a, 0x37, 0xe1, 0xca, 0xa7, 0x55, 0x8c, 0x12, 0x92, 0xf2, 0x2c, 0x1e, 0x97, 0x18, 0x5f,
	0x29, 0xc3, 0xfb, 0xdb, 0x2e, 0x2e, 0xbc, 0x56, 0xca, 0x38, 0xef, 0x36, 0x1b, 0x90, 0xc4, 0xa2,
	0x0f, 0x4d, 0x83, 0xc9, 0x55, 0x98, 0xb2, 0xee, 0x96, 0xcc, 0xa3, 0x2d, 0x76, 0xba, 0x3f, 0x2c,
	0xd8, 0x63, 0x2a, 0x9a, 0x26, 0x98, 0xa2, 0x90, 0xd0, 0xbf, 0xfd, 0x72, 0x6f, 0x56, 0xfc, 0xfd,
	0x0e, 0x8b, 0xdd, 0xfa, 0x73, 0x0b, 0x87, 0x7f, 0xad, 0x25, 0xb1, 0x73, 0x4f, 0x7c, 0x86, 0xde,
	0x86, 0x89, 0xc5, 0x61, 0x69, 0xc4, 0xa6, 0xe7, 0x1c, 0x6e, 0x2b, 0x13, 0xed, 0x79, 0x93, 0x7f,
	0x7f, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x18, 0x6d, 0x9f, 0x20, 0x05, 0x00, 0x00,
>>>>>>> Stashed changes
}
