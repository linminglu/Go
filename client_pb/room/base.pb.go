// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

/*
Package room is a generated protocol buffer package.

It is generated from these files:
	base.proto
	enum.proto
	error.proto
	game_common.proto
	game_ddz.proto
	game_erren.proto
	game_scxz.proto
	poker.proto

It has these top-level messages:
	GeographicalLocation
	RoomLoginReq
	RoomLoginRsp
	DeviceInfo
	RoomVisitorLoginReq
	RoomVisitorLoginRsp
	RoomJoinDeskReq
	RoomJoinDeskRsp
	RoomPlayerInfo
	RoomDeskCreatedNtf
	RoomDeskQuitReq
	RoomDeskQuitRsp
	RoomDeskDismissNtf
	RoomDeskQuitEnterNtf
	RoomDeskContinueReq
	RoomDeskContinueRsp
	RoomDeskNeedReusmeReq
	RoomChangePlayersReq
	RoomChangePlayersRsp
	RoomDeskNeedReusmeRsp
	RoomDeskChatReq
	RoomDeskChatNtf
	PlayerLocation
	RoomPlayerLocationReq
	RoomPlayerLocationRsp
	Card
	Fan
	RoomStartGameNtf
	RoomXipaiNtf
	PlayerCardCount
	RoomFapaiNtf
	RoomHuansanzhangNtf
	RoomHuansanzhangReq
	RoomHuansanzhangRsp
	RoomHuansanzhangFinishNtf
	RoomDingqueNtf
	RoomDingqueReq
	RoomDingqueRsp
	PlayerDingqueColor
	RoomDingqueFinishNtf
	RoomChiInfo
	RoomChupaiWenxunNtf
	RoomXingpaiActionReq
	RoomChiNtf
	RoomPengNtf
	RoomGangNtf
	TingCardInfo
	CanTingCardInfo
	RoomZixunNtf
	TingAction
	RoomChupaiReq
	RoomChupaiNtf
	RoomMopaiNtf
	RoomWaitQianggangHuNtf
	RoomHuNtf
	RoomTingInfoNtf
	CardsGroup
	PlayerCardsGroup
	RoomGameOverNtf
	BillDetail
	BillPlayerInfo
	RoomBalanceInfoRsp
	RoomSettleInstantRsp
	RoomResumeGameReq
	RoomCartoonFinishReq
	GamePlayerInfo
	GameDeskInfo
	RoomResumeGameRsp
	RoomTuoGuanNtf
	RoomCancelTuoGuanReq
	Result
	NextStage
	DDZStartGameNtf
	DDZDealNtf
	DDZGrabLordReq
	DDZGrabLordNtf
	DDZLordNtf
	DDZDoubleReq
	DDZDoubleNtf
	DDZPlayCardReq
	DDZPlayCardRsp
	DDZPlayCardNtf
	DDZGameOverNtf
	DDZBillPlayerInfo
	DDZTuoGuanReq
	DDZTuoGuanNtf
	DDZResumeGameReq
	DDZResumeGameRsp
	DDZDeskInfo
	DDZPlayerInfo
	RoomBuHuaInfo
	RoomBuHuaNtf
	RoomGiveUpReq
	RoomGiveUpRsp
	RoomGiveUpNtf
	Poker
*/
package room

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

// ChatType 聊天类型
type ChatType int32

const (
	ChatType_CT_EXPRESSION ChatType = 0
	ChatType_CT_QUICK      ChatType = 1
	ChatType_CT_VOICE      ChatType = 2
	ChatType_CT_WRITE      ChatType = 3
)

var ChatType_name = map[int32]string{
	0: "CT_EXPRESSION",
	1: "CT_QUICK",
	2: "CT_VOICE",
	3: "CT_WRITE",
}
var ChatType_value = map[string]int32{
	"CT_EXPRESSION": 0,
	"CT_QUICK":      1,
	"CT_VOICE":      2,
	"CT_WRITE":      3,
}

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}
func (x ChatType) String() string {
	return proto.EnumName(ChatType_name, int32(x))
}
func (x *ChatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChatType_value, data, "ChatType")
	if err != nil {
		return err
	}
	*x = ChatType(value)
	return nil
}
func (ChatType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// GeographicalLocation 玩家地理位置
type GeographicalLocation struct {
	Type             *LocSourceType `protobuf:"varint,1,opt,name=type,enum=room.LocSourceType" json:"type,omitempty"`
	Longitude        *float64       `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
	Latitude         *float64       `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GeographicalLocation) Reset()                    { *m = GeographicalLocation{} }
func (m *GeographicalLocation) String() string            { return proto.CompactTextString(m) }
func (*GeographicalLocation) ProtoMessage()               {}
func (*GeographicalLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GeographicalLocation) GetType() LocSourceType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return LocSourceType_LOC_SOURCE_BAIDU
}

func (m *GeographicalLocation) GetLongitude() float64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *GeographicalLocation) GetLatitude() float64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

// RoomLoginReq 房间登录请求
type RoomLoginReq struct {
	UserName         *string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomLoginReq) Reset()                    { *m = RoomLoginReq{} }
func (m *RoomLoginReq) String() string            { return proto.CompactTextString(m) }
func (*RoomLoginReq) ProtoMessage()               {}
func (*RoomLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RoomLoginReq) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

// RoomLoginRsp 登录房间响应
type RoomLoginRsp struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Coin             *uint64    `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	ErrCode          *RoomError `protobuf:"varint,3,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomLoginRsp) Reset()                    { *m = RoomLoginRsp{} }
func (m *RoomLoginRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomLoginRsp) ProtoMessage()               {}
func (*RoomLoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RoomLoginRsp) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *RoomLoginRsp) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *RoomLoginRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

// DeviceInfo 设备信息
type DeviceInfo struct {
	DeviceType       *DeviceType `protobuf:"varint,1,opt,name=device_type,json=deviceType,enum=room.DeviceType" json:"device_type,omitempty"`
	MacAddr          *string     `protobuf:"bytes,2,opt,name=mac_addr,json=macAddr" json:"mac_addr,omitempty"`
	Uuid             *string     `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	AndroidId        *string     `protobuf:"bytes,4,opt,name=android_id,json=androidId" json:"android_id,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeviceInfo) GetDeviceType() DeviceType {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return DeviceType_DT_ANDROID
}

func (m *DeviceInfo) GetMacAddr() string {
	if m != nil && m.MacAddr != nil {
		return *m.MacAddr
	}
	return ""
}

func (m *DeviceInfo) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *DeviceInfo) GetAndroidId() string {
	if m != nil && m.AndroidId != nil {
		return *m.AndroidId
	}
	return ""
}

// RoomVisitorLoginReq 游客登录请求
type RoomVisitorLoginReq struct {
	DeviceInfo       *DeviceInfo `protobuf:"bytes,1,opt,name=device_info,json=deviceInfo" json:"device_info,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RoomVisitorLoginReq) Reset()                    { *m = RoomVisitorLoginReq{} }
func (m *RoomVisitorLoginReq) String() string            { return proto.CompactTextString(m) }
func (*RoomVisitorLoginReq) ProtoMessage()               {}
func (*RoomVisitorLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RoomVisitorLoginReq) GetDeviceInfo() *DeviceInfo {
	if m != nil {
		return m.DeviceInfo
	}
	return nil
}

// RoomVisitorLoginRsp 游客登录应答
type RoomVisitorLoginRsp struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	UserName         *string    `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	PlayerId         *uint64    `protobuf:"varint,3,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Coin             *uint64    `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomVisitorLoginRsp) Reset()                    { *m = RoomVisitorLoginRsp{} }
func (m *RoomVisitorLoginRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomVisitorLoginRsp) ProtoMessage()               {}
func (*RoomVisitorLoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RoomVisitorLoginRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

func (m *RoomVisitorLoginRsp) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *RoomVisitorLoginRsp) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *RoomVisitorLoginRsp) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// RoomJoinDeskReq 申请进入房间请求
type RoomJoinDeskReq struct {
	Reserve          *uint32                 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	GameId           *GameId                 `protobuf:"varint,2,opt,name=game_id,json=gameId,enum=room.GameId" json:"game_id,omitempty"`
	Location         []*GeographicalLocation `protobuf:"bytes,3,rep,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *RoomJoinDeskReq) Reset()                    { *m = RoomJoinDeskReq{} }
func (m *RoomJoinDeskReq) String() string            { return proto.CompactTextString(m) }
func (*RoomJoinDeskReq) ProtoMessage()               {}
func (*RoomJoinDeskReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RoomJoinDeskReq) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

func (m *RoomJoinDeskReq) GetGameId() GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return GameId_GAMEID_XUELIU
}

func (m *RoomJoinDeskReq) GetLocation() []*GeographicalLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

// RoomJoinDeskRsp 响应加入房间
type RoomJoinDeskRsp struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomJoinDeskRsp) Reset()                    { *m = RoomJoinDeskRsp{} }
func (m *RoomJoinDeskRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomJoinDeskRsp) ProtoMessage()               {}
func (*RoomJoinDeskRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RoomJoinDeskRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

// RoomPlayerInfo 房间玩家数据
type RoomPlayerInfo struct {
	PlayerId         *uint64                 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Name             *string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Coin             *uint64                 `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Seat             *uint32                 `protobuf:"varint,4,opt,name=seat" json:"seat,omitempty"`
	Location         []*GeographicalLocation `protobuf:"bytes,5,rep,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *RoomPlayerInfo) Reset()                    { *m = RoomPlayerInfo{} }
func (m *RoomPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*RoomPlayerInfo) ProtoMessage()               {}
func (*RoomPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RoomPlayerInfo) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *RoomPlayerInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RoomPlayerInfo) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *RoomPlayerInfo) GetSeat() uint32 {
	if m != nil && m.Seat != nil {
		return *m.Seat
	}
	return 0
}

func (m *RoomPlayerInfo) GetLocation() []*GeographicalLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

// RoomDeskCreatedNtf 房间创建成功通知
type RoomDeskCreatedNtf struct {
	Players          []*RoomPlayerInfo `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
	GameId           *GameId           `protobuf:"varint,2,opt,name=game_id,json=gameId,enum=room.GameId" json:"game_id,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RoomDeskCreatedNtf) Reset()                    { *m = RoomDeskCreatedNtf{} }
func (m *RoomDeskCreatedNtf) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskCreatedNtf) ProtoMessage()               {}
func (*RoomDeskCreatedNtf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RoomDeskCreatedNtf) GetPlayers() []*RoomPlayerInfo {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *RoomDeskCreatedNtf) GetGameId() GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return GameId_GAMEID_XUELIU
}

// RoomQuitReq 退出牌局请求
type RoomDeskQuitReq struct {
	UserData         *uint32 `protobuf:"varint,1,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomDeskQuitReq) Reset()                    { *m = RoomDeskQuitReq{} }
func (m *RoomDeskQuitReq) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskQuitReq) ProtoMessage()               {}
func (*RoomDeskQuitReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RoomDeskQuitReq) GetUserData() uint32 {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return 0
}

// RommDeskQuitRsp 退出牌局响应
type RoomDeskQuitRsp struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	UserData         *uint32    `protobuf:"varint,2,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomDeskQuitRsp) Reset()                    { *m = RoomDeskQuitRsp{} }
func (m *RoomDeskQuitRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskQuitRsp) ProtoMessage()               {}
func (*RoomDeskQuitRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RoomDeskQuitRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

func (m *RoomDeskQuitRsp) GetUserData() uint32 {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return 0
}

// RoomDeskDismissNtf 牌桌解散通知
type RoomDeskDismissNtf struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomDeskDismissNtf) Reset()                    { *m = RoomDeskDismissNtf{} }
func (m *RoomDeskDismissNtf) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskDismissNtf) ProtoMessage()               {}
func (*RoomDeskDismissNtf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *RoomDeskDismissNtf) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

// RoomDeskQuitEnterNtf 退出加入牌局通知
type RoomDeskQuitEnterNtf struct {
	PlayerId         *uint64         `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Type             *QuitEnterType  `protobuf:"varint,2,opt,name=type,enum=room.QuitEnterType" json:"type,omitempty"`
	PlayerInfo       *RoomPlayerInfo `protobuf:"bytes,3,opt,name=player_info,json=playerInfo" json:"player_info,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RoomDeskQuitEnterNtf) Reset()                    { *m = RoomDeskQuitEnterNtf{} }
func (m *RoomDeskQuitEnterNtf) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskQuitEnterNtf) ProtoMessage()               {}
func (*RoomDeskQuitEnterNtf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RoomDeskQuitEnterNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *RoomDeskQuitEnterNtf) GetType() QuitEnterType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return QuitEnterType_QET_QUIT
}

func (m *RoomDeskQuitEnterNtf) GetPlayerInfo() *RoomPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

// RoomDeskContinueReq 续局请求
type RoomDeskContinueReq struct {
	Reserver         *uint32                 `protobuf:"varint,1,opt,name=reserver" json:"reserver,omitempty"`
	GameId           *GameId                 `protobuf:"varint,2,opt,name=game_id,json=gameId,enum=room.GameId" json:"game_id,omitempty"`
	Location         []*GeographicalLocation `protobuf:"bytes,3,rep,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *RoomDeskContinueReq) Reset()                    { *m = RoomDeskContinueReq{} }
func (m *RoomDeskContinueReq) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskContinueReq) ProtoMessage()               {}
func (*RoomDeskContinueReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *RoomDeskContinueReq) GetReserver() uint32 {
	if m != nil && m.Reserver != nil {
		return *m.Reserver
	}
	return 0
}

func (m *RoomDeskContinueReq) GetGameId() GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return GameId_GAMEID_XUELIU
}

func (m *RoomDeskContinueReq) GetLocation() []*GeographicalLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

// RoomDeskContinueRsp 续局应答
type RoomDeskContinueRsp struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomDeskContinueRsp) Reset()                    { *m = RoomDeskContinueRsp{} }
func (m *RoomDeskContinueRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskContinueRsp) ProtoMessage()               {}
func (*RoomDeskContinueRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *RoomDeskContinueRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

// RoomGameStatusReq 是否需要恢复牌局请求
type RoomDeskNeedReusmeReq struct {
	Reserve          *uint32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomDeskNeedReusmeReq) Reset()                    { *m = RoomDeskNeedReusmeReq{} }
func (m *RoomDeskNeedReusmeReq) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskNeedReusmeReq) ProtoMessage()               {}
func (*RoomDeskNeedReusmeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *RoomDeskNeedReusmeReq) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// RoomChangePlayersReq 换对手请求
type RoomChangePlayersReq struct {
	GameId           *GameId `protobuf:"varint,1,opt,name=game_id,json=gameId,enum=room.GameId" json:"game_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomChangePlayersReq) Reset()                    { *m = RoomChangePlayersReq{} }
func (m *RoomChangePlayersReq) String() string            { return proto.CompactTextString(m) }
func (*RoomChangePlayersReq) ProtoMessage()               {}
func (*RoomChangePlayersReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *RoomChangePlayersReq) GetGameId() GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return GameId_GAMEID_XUELIU
}

// RoomChangePlayersRsp 换对手应答
type RoomChangePlayersRsp struct {
	ErrCode          *RoomError `protobuf:"varint,1,opt,name=err_code,json=errCode,enum=room.RoomError" json:"err_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoomChangePlayersRsp) Reset()                    { *m = RoomChangePlayersRsp{} }
func (m *RoomChangePlayersRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomChangePlayersRsp) ProtoMessage()               {}
func (*RoomChangePlayersRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *RoomChangePlayersRsp) GetErrCode() RoomError {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return RoomError_SUCCESS
}

// RoomGameStatusRsp 是否需要恢复牌局应答
type RoomDeskNeedReusmeRsp struct {
	IsNeed           *bool   `protobuf:"varint,1,opt,name=is_need,json=isNeed" json:"is_need,omitempty"`
	GameId           *GameId `protobuf:"varint,2,opt,name=game_id,json=gameId,enum=room.GameId" json:"game_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomDeskNeedReusmeRsp) Reset()                    { *m = RoomDeskNeedReusmeRsp{} }
func (m *RoomDeskNeedReusmeRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskNeedReusmeRsp) ProtoMessage()               {}
func (*RoomDeskNeedReusmeRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *RoomDeskNeedReusmeRsp) GetIsNeed() bool {
	if m != nil && m.IsNeed != nil {
		return *m.IsNeed
	}
	return false
}

func (m *RoomDeskNeedReusmeRsp) GetGameId() GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return GameId_GAMEID_XUELIU
}

// RoomDeskChatReq 聊天请求
type RoomDeskChatReq struct {
	ChatType         *ChatType `protobuf:"varint,1,opt,name=chat_type,json=chatType,enum=room.ChatType" json:"chat_type,omitempty"`
	ChatInfo         *string   `protobuf:"bytes,2,opt,name=chat_info,json=chatInfo" json:"chat_info,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *RoomDeskChatReq) Reset()                    { *m = RoomDeskChatReq{} }
func (m *RoomDeskChatReq) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskChatReq) ProtoMessage()               {}
func (*RoomDeskChatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *RoomDeskChatReq) GetChatType() ChatType {
	if m != nil && m.ChatType != nil {
		return *m.ChatType
	}
	return ChatType_CT_EXPRESSION
}

func (m *RoomDeskChatReq) GetChatInfo() string {
	if m != nil && m.ChatInfo != nil {
		return *m.ChatInfo
	}
	return ""
}

// RoomDeskChatNtf 聊天通知
type RoomDeskChatNtf struct {
	PlayerId         *uint64   `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	ChatType         *ChatType `protobuf:"varint,2,opt,name=chat_type,json=chatType,enum=room.ChatType" json:"chat_type,omitempty"`
	ChatInfo         *string   `protobuf:"bytes,3,opt,name=chat_info,json=chatInfo" json:"chat_info,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *RoomDeskChatNtf) Reset()                    { *m = RoomDeskChatNtf{} }
func (m *RoomDeskChatNtf) String() string            { return proto.CompactTextString(m) }
func (*RoomDeskChatNtf) ProtoMessage()               {}
func (*RoomDeskChatNtf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *RoomDeskChatNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *RoomDeskChatNtf) GetChatType() ChatType {
	if m != nil && m.ChatType != nil {
		return *m.ChatType
	}
	return ChatType_CT_EXPRESSION
}

func (m *RoomDeskChatNtf) GetChatInfo() string {
	if m != nil && m.ChatInfo != nil {
		return *m.ChatInfo
	}
	return ""
}

// PlayerLocation 玩家地理位置
type PlayerLocation struct {
	PlayerId         *uint64                 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Location         []*GeographicalLocation `protobuf:"bytes,2,rep,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *PlayerLocation) Reset()                    { *m = PlayerLocation{} }
func (m *PlayerLocation) String() string            { return proto.CompactTextString(m) }
func (*PlayerLocation) ProtoMessage()               {}
func (*PlayerLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *PlayerLocation) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *PlayerLocation) GetLocation() []*GeographicalLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

// RoomPlayerLocationReq 玩家地理位置请求
type RoomPlayerLocationReq struct {
	Reserve          *uint32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RoomPlayerLocationReq) Reset()                    { *m = RoomPlayerLocationReq{} }
func (m *RoomPlayerLocationReq) String() string            { return proto.CompactTextString(m) }
func (*RoomPlayerLocationReq) ProtoMessage()               {}
func (*RoomPlayerLocationReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *RoomPlayerLocationReq) GetReserve() uint32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// RoomPlayerLocationRsp 玩家地理位置应答
type RoomPlayerLocationRsp struct {
	Locations        []*PlayerLocation `protobuf:"bytes,1,rep,name=locations" json:"locations,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RoomPlayerLocationRsp) Reset()                    { *m = RoomPlayerLocationRsp{} }
func (m *RoomPlayerLocationRsp) String() string            { return proto.CompactTextString(m) }
func (*RoomPlayerLocationRsp) ProtoMessage()               {}
func (*RoomPlayerLocationRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *RoomPlayerLocationRsp) GetLocations() []*PlayerLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

func init() {
	proto.RegisterType((*GeographicalLocation)(nil), "room.GeographicalLocation")
	proto.RegisterType((*RoomLoginReq)(nil), "room.RoomLoginReq")
	proto.RegisterType((*RoomLoginRsp)(nil), "room.RoomLoginRsp")
	proto.RegisterType((*DeviceInfo)(nil), "room.DeviceInfo")
	proto.RegisterType((*RoomVisitorLoginReq)(nil), "room.RoomVisitorLoginReq")
	proto.RegisterType((*RoomVisitorLoginRsp)(nil), "room.RoomVisitorLoginRsp")
	proto.RegisterType((*RoomJoinDeskReq)(nil), "room.RoomJoinDeskReq")
	proto.RegisterType((*RoomJoinDeskRsp)(nil), "room.RoomJoinDeskRsp")
	proto.RegisterType((*RoomPlayerInfo)(nil), "room.RoomPlayerInfo")
	proto.RegisterType((*RoomDeskCreatedNtf)(nil), "room.RoomDeskCreatedNtf")
	proto.RegisterType((*RoomDeskQuitReq)(nil), "room.RoomDeskQuitReq")
	proto.RegisterType((*RoomDeskQuitRsp)(nil), "room.RoomDeskQuitRsp")
	proto.RegisterType((*RoomDeskDismissNtf)(nil), "room.RoomDeskDismissNtf")
	proto.RegisterType((*RoomDeskQuitEnterNtf)(nil), "room.RoomDeskQuitEnterNtf")
	proto.RegisterType((*RoomDeskContinueReq)(nil), "room.RoomDeskContinueReq")
	proto.RegisterType((*RoomDeskContinueRsp)(nil), "room.RoomDeskContinueRsp")
	proto.RegisterType((*RoomDeskNeedReusmeReq)(nil), "room.RoomDeskNeedReusmeReq")
	proto.RegisterType((*RoomChangePlayersReq)(nil), "room.RoomChangePlayersReq")
	proto.RegisterType((*RoomChangePlayersRsp)(nil), "room.RoomChangePlayersRsp")
	proto.RegisterType((*RoomDeskNeedReusmeRsp)(nil), "room.RoomDeskNeedReusmeRsp")
	proto.RegisterType((*RoomDeskChatReq)(nil), "room.RoomDeskChatReq")
	proto.RegisterType((*RoomDeskChatNtf)(nil), "room.RoomDeskChatNtf")
	proto.RegisterType((*PlayerLocation)(nil), "room.PlayerLocation")
	proto.RegisterType((*RoomPlayerLocationReq)(nil), "room.RoomPlayerLocationReq")
	proto.RegisterType((*RoomPlayerLocationRsp)(nil), "room.RoomPlayerLocationRsp")
	proto.RegisterEnum("room.ChatType", ChatType_name, ChatType_value)
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 889 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xe1, 0x6e, 0x1b, 0x45,
	0x10, 0xe6, 0x6c, 0x53, 0x9f, 0xc7, 0x89, 0x6b, 0xb6, 0x01, 0x8c, 0x01, 0x29, 0x3a, 0x09, 0x11,
	0xb5, 0x92, 0xab, 0x46, 0x82, 0x7f, 0x95, 0x68, 0x1d, 0x2b, 0x98, 0x46, 0x69, 0xbb, 0x09, 0x2d,
	0x2a, 0x3f, 0x4e, 0xdb, 0xdb, 0x89, 0xb3, 0x8a, 0xef, 0xf6, 0xd8, 0xbd, 0x8b, 0xf0, 0x2b, 0x20,
	0x21, 0x90, 0x78, 0x02, 0xde, 0x14, 0xed, 0xee, 0x9d, 0x7d, 0x67, 0x19, 0xb7, 0x46, 0xe2, 0x57,
	0x66, 0xe7, 0x66, 0x67, 0xbe, 0xf9, 0xe6, 0x9b, 0x8d, 0x01, 0xde, 0x32, 0x8d, 0xa3, 0x54, 0xc9,
	0x4c, 0x92, 0x96, 0x92, 0x32, 0x1e, 0x76, 0x51, 0x29, 0xa9, 0x9c, 0x6b, 0x08, 0x98, 0xe4, 0xb1,
	0xb3, 0x83, 0x05, 0x1c, 0x9c, 0xa2, 0x9c, 0x29, 0x96, 0x5e, 0x8b, 0x88, 0xcd, 0xcf, 0x64, 0xc4,
	0x32, 0x21, 0x13, 0xf2, 0x35, 0xb4, 0xb2, 0x45, 0x8a, 0x03, 0xef, 0xd0, 0x3b, 0xea, 0x1d, 0xdf,
	0x1b, 0x99, 0x2c, 0xa3, 0x33, 0x19, 0x5d, 0xc8, 0x5c, 0x45, 0x78, 0xb9, 0x48, 0x91, 0xda, 0x00,
	0xf2, 0x05, 0x74, 0xe6, 0x32, 0x99, 0x89, 0x2c, 0xe7, 0x38, 0x68, 0x1c, 0x7a, 0x47, 0x1e, 0x5d,
	0x39, 0xc8, 0x10, 0xfc, 0x39, 0xcb, 0xdc, 0xc7, 0xa6, 0xfd, 0xb8, 0x3c, 0x07, 0x0f, 0x60, 0x8f,
	0x4a, 0x19, 0x9f, 0xc9, 0x99, 0x48, 0x28, 0xfe, 0x42, 0x3e, 0x87, 0x4e, 0xae, 0x51, 0x85, 0x09,
	0x8b, 0x5d, 0xdd, 0x0e, 0xf5, 0x8d, 0xe3, 0x9c, 0xc5, 0x18, 0xdc, 0x54, 0x83, 0x75, 0x6a, 0x82,
	0xd3, 0x39, 0x5b, 0xa0, 0x0a, 0x05, 0xb7, 0xc1, 0x2d, 0xea, 0x3b, 0xc7, 0x94, 0x13, 0x02, 0xad,
	0x48, 0x8a, 0xc4, 0xc2, 0x69, 0x51, 0x6b, 0x93, 0xfb, 0xe0, 0xa3, 0x52, 0x61, 0x24, 0x0b, 0x24,
	0xbd, 0xe3, 0xbb, 0xae, 0x29, 0x93, 0x76, 0x62, 0xd8, 0xa1, 0x6d, 0x54, 0x6a, 0x2c, 0x39, 0x06,
	0xbf, 0x7b, 0x00, 0x27, 0x78, 0x2b, 0x22, 0x9c, 0x26, 0x57, 0x92, 0x3c, 0x82, 0x2e, 0xb7, 0xa7,
	0xb0, 0x42, 0x49, 0xdf, 0xdd, 0x76, 0x61, 0x96, 0x0f, 0xe0, 0x4b, 0x9b, 0x7c, 0x06, 0x7e, 0xcc,
	0xa2, 0x90, 0x71, 0xae, 0x2c, 0x8a, 0x0e, 0x6d, 0xc7, 0x2c, 0x7a, 0xc2, 0xb9, 0x32, 0xe0, 0xf2,
	0x5c, 0x70, 0x0b, 0xa2, 0x43, 0xad, 0x4d, 0xbe, 0x04, 0x60, 0x09, 0x57, 0x52, 0x70, 0xd3, 0x4e,
	0xcb, 0x7e, 0xe9, 0x14, 0x9e, 0x29, 0x0f, 0xbe, 0x87, 0x7b, 0x06, 0xe5, 0x2b, 0xa1, 0x45, 0x26,
	0xd5, 0x92, 0xb0, 0x15, 0x2e, 0x91, 0x5c, 0x49, 0x8b, 0xab, 0x5b, 0xc7, 0x65, 0xe0, 0x97, 0xb8,
	0x8c, 0x1d, 0xfc, 0xe1, 0x6d, 0x48, 0xa5, 0xd3, 0x1a, 0x3b, 0xde, 0x76, 0x76, 0xea, 0x73, 0x6a,
	0xd4, 0xe7, 0x54, 0x9f, 0x4b, 0xf3, 0x5f, 0xe6, 0xd2, 0x5a, 0xcd, 0x25, 0xf8, 0xcd, 0x83, 0xbb,
	0xa6, 0xc8, 0x0f, 0x52, 0x24, 0x27, 0xa8, 0x6f, 0x4c, 0x63, 0x03, 0x68, 0x2b, 0xd4, 0xa8, 0x6e,
	0x1d, 0x98, 0x7d, 0x5a, 0x1e, 0xc9, 0x57, 0xd0, 0x9e, 0xb1, 0x18, 0x4d, 0xf2, 0x86, 0x85, 0xb9,
	0xe7, 0x60, 0x9e, 0xb2, 0x18, 0xa7, 0x9c, 0xde, 0x99, 0xd9, 0xbf, 0xe4, 0x5b, 0xf0, 0xe7, 0x85,
	0x92, 0x07, 0xcd, 0xc3, 0xe6, 0x51, 0xf7, 0x78, 0x58, 0xc4, 0x6d, 0xd0, 0x3a, 0x5d, 0xc6, 0x06,
	0x8f, 0xd7, 0xb0, 0xec, 0xc6, 0x4c, 0xf0, 0xb7, 0x07, 0x3d, 0xe3, 0x7e, 0xe1, 0x1a, 0x36, 0xda,
	0x79, 0x97, 0x4e, 0x2b, 0x24, 0x5a, 0x7b, 0xc9, 0x51, 0xb3, 0xa2, 0x5d, 0x02, 0x2d, 0x8d, 0x2c,
	0xb3, 0xbc, 0xed, 0x53, 0x6b, 0xd7, 0x5a, 0xfc, 0x70, 0x87, 0x16, 0x6f, 0x80, 0x18, 0x88, 0xa6,
	0xbd, 0xb1, 0x42, 0x96, 0x21, 0x3f, 0xcf, 0xae, 0xc8, 0x08, 0xda, 0x0e, 0x95, 0x1e, 0x78, 0x36,
	0xd9, 0xc1, 0xaa, 0xc9, 0x55, 0x37, 0xb4, 0x0c, 0x7a, 0xcf, 0x39, 0x04, 0x23, 0xc7, 0xa7, 0x29,
	0xf6, 0x32, 0x17, 0x59, 0x75, 0xcb, 0x39, 0xcb, 0x58, 0x31, 0x5d, 0xab, 0x9e, 0x13, 0x96, 0xb1,
	0xe0, 0xcd, 0x5a, 0xfc, 0x7f, 0x54, 0xa6, 0xcd, 0xdd, 0x58, 0xcb, 0xfd, 0xdd, 0xaa, 0xf1, 0x13,
	0xa1, 0x63, 0xa1, 0xb5, 0x69, 0x7c, 0x97, 0xf1, 0xfe, 0xe5, 0xc1, 0x41, 0x15, 0xde, 0x24, 0xc9,
	0x50, 0x99, 0x24, 0x5b, 0x87, 0x5c, 0xbe, 0xa4, 0x8d, 0xea, 0x4b, 0xba, 0xbc, 0x5e, 0x79, 0x49,
	0xbf, 0x81, 0x6e, 0x99, 0xc5, 0xac, 0x73, 0xd3, 0xae, 0xf3, 0xe6, 0x39, 0x40, 0xba, 0xb4, 0x83,
	0x3f, 0x8b, 0x95, 0xb6, 0x13, 0x95, 0x49, 0x26, 0x92, 0x1c, 0x0d, 0xd1, 0x43, 0xf0, 0x8b, 0xad,
	0x51, 0x25, 0xcf, 0xe5, 0xf9, 0xff, 0x5e, 0xa3, 0x27, 0x1b, 0x10, 0xed, 0xb8, 0x4a, 0x8f, 0xe0,
	0xe3, 0x32, 0xc5, 0x39, 0x22, 0xa7, 0x98, 0xeb, 0x18, 0xb7, 0xbe, 0x0d, 0xc1, 0x63, 0x37, 0x9d,
	0xf1, 0x35, 0x4b, 0x66, 0xe8, 0xc8, 0xd2, 0xe6, 0x46, 0xa5, 0x59, 0x6f, 0x8b, 0x56, 0x9f, 0x6e,
	0xba, 0xbe, 0x23, 0xea, 0xd7, 0x1b, 0x51, 0xeb, 0x94, 0x7c, 0x0a, 0x6d, 0xa1, 0xc3, 0x04, 0xd1,
	0x61, 0xf0, 0xe9, 0x1d, 0xa1, 0x4d, 0xc4, 0xfb, 0x2e, 0xd2, 0xcf, 0xab, 0xc5, 0x18, 0x5f, 0x33,
	0xbb, 0x48, 0x0f, 0xa0, 0x13, 0x5d, 0xb3, 0xac, 0xfa, 0x3f, 0xa9, 0xe7, 0xee, 0x9a, 0x08, 0xab,
	0x2b, 0x3f, 0x2a, 0x2c, 0xa3, 0x50, 0x1b, 0x6c, 0x95, 0x55, 0xbc, 0xd9, 0xc6, 0x61, 0x15, 0xf4,
	0x6b, 0x3d, 0xf9, 0x3b, 0x15, 0x5d, 0xab, 0xdc, 0xd8, 0xa5, 0x72, 0x73, 0xad, 0x32, 0x42, 0xcf,
	0x31, 0xbd, 0xfc, 0xdd, 0xb1, 0xb5, 0x70, 0x55, 0x8f, 0x8d, 0x1d, 0xf4, 0x58, 0x88, 0xa9, 0x5e,
	0x6a, 0xbb, 0x98, 0x9e, 0x6d, 0xbc, 0xa2, 0x53, 0x72, 0x6c, 0x7e, 0xef, 0xb8, 0xe3, 0xda, 0x5b,
	0xb9, 0x16, 0xbb, 0x0a, 0xbb, 0x7f, 0x0a, 0x7e, 0xc9, 0x0c, 0xf9, 0x08, 0xf6, 0xc7, 0x97, 0xe1,
	0xe4, 0xa7, 0x17, 0x74, 0x72, 0x71, 0x31, 0x7d, 0x7e, 0xde, 0xff, 0x80, 0xec, 0x81, 0x3f, 0xbe,
	0x0c, 0x5f, 0xfe, 0x38, 0x1d, 0x3f, 0xeb, 0x7b, 0xc5, 0xe9, 0xd5, 0xf3, 0xe9, 0x78, 0xd2, 0x6f,
	0x14, 0xa7, 0xd7, 0x74, 0x7a, 0x39, 0xe9, 0x37, 0x9f, 0x7e, 0xf2, 0xe6, 0x40, 0x67, 0x78, 0x8b,
	0x0f, 0xa3, 0xb9, 0xc0, 0x24, 0x0b, 0xd3, 0xb7, 0x0f, 0x4d, 0xe9, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0xa9, 0x91, 0xb2, 0xf1, 0x09, 0x00, 0x00,
}