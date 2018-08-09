// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room_mgr.proto

/*
Package roommgr is a generated protocol buffer package.

It is generated from these files:
	room_mgr.proto

It has these top-level messages:
	DeskPlayer
	CreateDeskRequest
	CreateDeskResponse
*/
package roommgr

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoomError int32

const (
	RoomError_SUCCESS RoomError = 0
	RoomError_FAILED  RoomError = 1
)

var RoomError_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
}
var RoomError_value = map[string]int32{
	"SUCCESS": 0,
	"FAILED":  1,
}

func (x RoomError) String() string {
	return proto.EnumName(RoomError_name, int32(x))
}
func (RoomError) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 牌桌玩家
type DeskPlayer struct {
	PlayerId   uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	RobotLevel int32  `protobuf:"varint,2,opt,name=robot_level,json=robotLevel" json:"robot_level,omitempty"`
	Seat       uint32 `protobuf:"varint,3,opt,name=seat" json:"seat,omitempty"`
}

func (m *DeskPlayer) Reset()                    { *m = DeskPlayer{} }
func (m *DeskPlayer) String() string            { return proto.CompactTextString(m) }
func (*DeskPlayer) ProtoMessage()               {}
func (*DeskPlayer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeskPlayer) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *DeskPlayer) GetRobotLevel() int32 {
	if m != nil {
		return m.RobotLevel
	}
	return 0
}

func (m *DeskPlayer) GetSeat() uint32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

// 创建桌子的请求
type CreateDeskRequest struct {
	GameId  uint32        `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	LevelId uint32        `protobuf:"varint,2,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	DeskId  uint64        `protobuf:"varint,3,opt,name=desk_id,json=deskId" json:"desk_id,omitempty"`
	Players []*DeskPlayer `protobuf:"bytes,4,rep,name=players" json:"players,omitempty"`
}

func (m *CreateDeskRequest) Reset()                    { *m = CreateDeskRequest{} }
func (m *CreateDeskRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDeskRequest) ProtoMessage()               {}
func (*CreateDeskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateDeskRequest) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *CreateDeskRequest) GetLevelId() uint32 {
	if m != nil {
		return m.LevelId
	}
	return 0
}

func (m *CreateDeskRequest) GetDeskId() uint64 {
	if m != nil {
		return m.DeskId
	}
	return 0
}

func (m *CreateDeskRequest) GetPlayers() []*DeskPlayer {
	if m != nil {
		return m.Players
	}
	return nil
}

// 创建桌子的回复
type CreateDeskResponse struct {
	ErrCode RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=roommgr.RoomError" json:"err_code,omitempty"`
}

func (m *CreateDeskResponse) Reset()                    { *m = CreateDeskResponse{} }
func (m *CreateDeskResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDeskResponse) ProtoMessage()               {}
func (*CreateDeskResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateDeskResponse) GetErrCode() RoomError {
	if m != nil {
		return m.ErrCode
	}
	return RoomError_SUCCESS
}

func init() {
	proto.RegisterType((*DeskPlayer)(nil), "roommgr.DeskPlayer")
	proto.RegisterType((*CreateDeskRequest)(nil), "roommgr.CreateDeskRequest")
	proto.RegisterType((*CreateDeskResponse)(nil), "roommgr.CreateDeskResponse")
	proto.RegisterEnum("roommgr.RoomError", RoomError_name, RoomError_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RoomMgr service

type RoomMgrClient interface {
	CreateDesk(ctx context.Context, in *CreateDeskRequest, opts ...grpc.CallOption) (*CreateDeskResponse, error)
}

type roomMgrClient struct {
	cc *grpc.ClientConn
}

func NewRoomMgrClient(cc *grpc.ClientConn) RoomMgrClient {
	return &roomMgrClient{cc}
}

func (c *roomMgrClient) CreateDesk(ctx context.Context, in *CreateDeskRequest, opts ...grpc.CallOption) (*CreateDeskResponse, error) {
	out := new(CreateDeskResponse)
	err := grpc.Invoke(ctx, "/roommgr.RoomMgr/CreateDesk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomMgr service

type RoomMgrServer interface {
	CreateDesk(context.Context, *CreateDeskRequest) (*CreateDeskResponse, error)
}

func RegisterRoomMgrServer(s *grpc.Server, srv RoomMgrServer) {
	s.RegisterService(&_RoomMgr_serviceDesc, srv)
}

func _RoomMgr_CreateDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomMgrServer).CreateDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roommgr.RoomMgr/CreateDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomMgrServer).CreateDesk(ctx, req.(*CreateDeskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roommgr.RoomMgr",
	HandlerType: (*RoomMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDesk",
			Handler:    _RoomMgr_CreateDesk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room_mgr.proto",
}

func init() { proto.RegisterFile("room_mgr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xa5, 0xc0, 0xd7, 0x85, 0x21, 0x10, 0xbe, 0xf1, 0x60, 0x85, 0x83, 0x4d, 0xe3, 0xa1, 0x31,
	0x81, 0x03, 0xfe, 0x02, 0x53, 0xd0, 0x34, 0xc1, 0xc4, 0x2c, 0xf1, 0x6a, 0x53, 0xdc, 0x49, 0x63,
	0xa0, 0x2e, 0xce, 0x56, 0x13, 0x7f, 0x85, 0x7f, 0xd9, 0xec, 0x22, 0xc5, 0x44, 0x6f, 0xf3, 0xde,
	0x9b, 0xbe, 0xf7, 0xa6, 0x0b, 0x03, 0xd6, 0xba, 0xcc, 0xca, 0x82, 0xa7, 0x3b, 0xd6, 0x95, 0x46,
	0x61, 0x71, 0x59, 0x70, 0xf4, 0x08, 0x30, 0x27, 0xb3, 0xb9, 0xdf, 0xe6, 0x1f, 0xc4, 0x38, 0x86,
	0xee, 0xce, 0x4d, 0xd9, 0xb3, 0x0a, 0xbc, 0xd0, 0x8b, 0xdb, 0xb2, 0xb3, 0x27, 0x52, 0x85, 0xe7,
	0xd0, 0x63, 0xbd, 0xd6, 0x55, 0xb6, 0xa5, 0x77, 0xda, 0x06, 0xcd, 0xd0, 0x8b, 0xff, 0x49, 0x70,
	0xd4, 0xd2, 0x32, 0x88, 0xd0, 0x36, 0x94, 0x57, 0x41, 0x2b, 0xf4, 0xe2, 0xbe, 0x74, 0x73, 0xf4,
	0xe9, 0xc1, 0xff, 0x84, 0x29, 0xaf, 0xc8, 0xc6, 0x48, 0x7a, 0x7d, 0x23, 0x53, 0xe1, 0x29, 0x88,
	0x22, 0x2f, 0xe9, 0x90, 0xd2, 0x97, 0xbe, 0x85, 0xa9, 0xc2, 0x33, 0xe8, 0x38, 0x77, 0xab, 0x34,
	0x9d, 0x22, 0x1c, 0x4e, 0x95, 0xfd, 0x46, 0x91, 0xd9, 0x58, 0xa5, 0xe5, 0x9a, 0xf9, 0x16, 0xa6,
	0x0a, 0x27, 0x20, 0xf6, 0x1d, 0x4d, 0xd0, 0x0e, 0x5b, 0x71, 0x6f, 0x76, 0x32, 0xfd, 0xbe, 0x6e,
	0x7a, 0x3c, 0x4d, 0x1e, 0x76, 0xa2, 0x04, 0xf0, 0x67, 0x21, 0xb3, 0xd3, 0x2f, 0x86, 0x70, 0x02,
	0x1d, 0x62, 0xce, 0x9e, 0xb4, 0x22, 0x57, 0x69, 0x30, 0xc3, 0xda, 0x45, 0x6a, 0x5d, 0x2e, 0x98,
	0x35, 0x4b, 0x41, 0xcc, 0x89, 0x56, 0x74, 0x79, 0x01, 0xdd, 0x9a, 0xc5, 0x1e, 0x88, 0xd5, 0x43,
	0x92, 0x2c, 0x56, 0xab, 0x61, 0x03, 0x01, 0xfc, 0x9b, 0xeb, 0x74, 0xb9, 0x98, 0x0f, 0xbd, 0x99,
	0x04, 0x61, 0xb7, 0xee, 0x0a, 0xc6, 0x5b, 0x80, 0x63, 0x2a, 0x8e, 0x6a, 0xef, 0x5f, 0xff, 0x66,
	0x34, 0xfe, 0x53, 0xdb, 0xd7, 0x8c, 0x1a, 0x6b, 0xdf, 0x3d, 0xe0, 0xd5, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0xc5, 0x56, 0xee, 0xd2, 0x01, 0x00, 0x00,
}
