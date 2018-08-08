// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common.proto
	errors.proto

It has these top-level messages:
	GeographicalLocation
	GameConfig
	GameLevelConfig
	Result
*/
package common

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

// GameId 游戏 ID
type GameId int32

const (
	GameId_GAMEID_XUELIU   GameId = 1
	GameId_GAMEID_XUEZHAN  GameId = 2
	GameId_GAMEID_DOUDIZHU GameId = 3
	GameId_GAMEID_ERRENMJ  GameId = 4
)

var GameId_name = map[int32]string{
	1: "GAMEID_XUELIU",
	2: "GAMEID_XUEZHAN",
	3: "GAMEID_DOUDIZHU",
	4: "GAMEID_ERRENMJ",
}
var GameId_value = map[string]int32{
	"GAMEID_XUELIU":   1,
	"GAMEID_XUEZHAN":  2,
	"GAMEID_DOUDIZHU": 3,
	"GAMEID_ERRENMJ":  4,
}

func (x GameId) Enum() *GameId {
	p := new(GameId)
	*p = x
	return p
}
func (x GameId) String() string {
	return proto.EnumName(GameId_name, int32(x))
}
func (x *GameId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GameId_value, data, "GameId")
	if err != nil {
		return err
	}
	*x = GameId(value)
	return nil
}
func (GameId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// LocSourceType 位置信息来源类型
type LocSourceType int32

const (
	LocSourceType_LOC_SOURCE_BAIDU  LocSourceType = 0
	LocSourceType_LOC_SOURCE_JIZHAN LocSourceType = 1
)

var LocSourceType_name = map[int32]string{
	0: "LOC_SOURCE_BAIDU",
	1: "LOC_SOURCE_JIZHAN",
}
var LocSourceType_value = map[string]int32{
	"LOC_SOURCE_BAIDU":  0,
	"LOC_SOURCE_JIZHAN": 1,
}

func (x LocSourceType) Enum() *LocSourceType {
	p := new(LocSourceType)
	*p = x
	return p
}
func (x LocSourceType) String() string {
	return proto.EnumName(LocSourceType_name, int32(x))
}
func (x *LocSourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LocSourceType_value, data, "LocSourceType")
	if err != nil {
		return err
	}
	*x = LocSourceType(value)
	return nil
}
func (LocSourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// PlayerState 玩家状态
type PlayerState int32

const (
	PlayerState_PS_IDLE     PlayerState = 0
	PlayerState_PS_GAMEING  PlayerState = 1
	PlayerState_PS_MATCHING PlayerState = 2
)

var PlayerState_name = map[int32]string{
	0: "PS_IDLE",
	1: "PS_GAMEING",
	2: "PS_MATCHING",
}
var PlayerState_value = map[string]int32{
	"PS_IDLE":     0,
	"PS_GAMEING":  1,
	"PS_MATCHING": 2,
}

func (x PlayerState) Enum() *PlayerState {
	p := new(PlayerState)
	*p = x
	return p
}
func (x PlayerState) String() string {
	return proto.EnumName(PlayerState_name, int32(x))
}
func (x *PlayerState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlayerState_value, data, "PlayerState")
	if err != nil {
		return err
	}
	*x = PlayerState(value)
	return nil
}
func (PlayerState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// PlayerGender 玩家性别
type PlayerGender int32

const (
	PlayerGender_PG_NIL    PlayerGender = 0
	PlayerGender_PG_MALE   PlayerGender = 1
	PlayerGender_PG_FEMALE PlayerGender = 2
)

var PlayerGender_name = map[int32]string{
	0: "PG_NIL",
	1: "PG_MALE",
	2: "PG_FEMALE",
}
var PlayerGender_value = map[string]int32{
	"PG_NIL":    0,
	"PG_MALE":   1,
	"PG_FEMALE": 2,
}

func (x PlayerGender) Enum() *PlayerGender {
	p := new(PlayerGender)
	*p = x
	return p
}
func (x PlayerGender) String() string {
	return proto.EnumName(PlayerGender_name, int32(x))
}
func (x *PlayerGender) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlayerGender_value, data, "PlayerGender")
	if err != nil {
		return err
	}
	*x = PlayerGender(value)
	return nil
}
func (PlayerGender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Platform 平台
type Platform int32

const (
	Platform_Android Platform = 1
	Platform_Iphone  Platform = 2
)

var Platform_name = map[int32]string{
	1: "Android",
	2: "Iphone",
}
var Platform_value = map[string]int32{
	"Android": 1,
	"Iphone":  2,
}

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}
func (x Platform) String() string {
	return proto.EnumName(Platform_name, int32(x))
}
func (x *Platform) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Platform_value, data, "Platform")
	if err != nil {
		return err
	}
	*x = Platform(value)
	return nil
}
func (Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// GeographicalLocation 玩家地理位置
type GeographicalLocation struct {
	Type             *LocSourceType `protobuf:"varint,1,opt,name=type,enum=common.LocSourceType" json:"type,omitempty"`
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

// GameLevelConfig 游戏玩法
type GameConfig struct {
	GameId           *uint32 `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	GameName         *string `protobuf:"bytes,2,opt,name=game_name,json=gameName" json:"game_name,omitempty"`
	GameType         *uint32 `protobuf:"varint,3,opt,name=game_type,json=gameType" json:"game_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameConfig) Reset()                    { *m = GameConfig{} }
func (m *GameConfig) String() string            { return proto.CompactTextString(m) }
func (*GameConfig) ProtoMessage()               {}
func (*GameConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GameConfig) GetGameId() uint32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameConfig) GetGameName() string {
	if m != nil && m.GameName != nil {
		return *m.GameName
	}
	return ""
}

func (m *GameConfig) GetGameType() uint32 {
	if m != nil && m.GameType != nil {
		return *m.GameType
	}
	return 0
}

// GameLevelConfig 游戏等级
type GameLevelConfig struct {
	GameId           *uint32 `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	LevelId          *uint32 `protobuf:"varint,2,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	LevelName        *string `protobuf:"bytes,3,opt,name=level_name,json=levelName" json:"level_name,omitempty"`
	BaseScores       *uint32 `protobuf:"varint,4,opt,name=base_scores,json=baseScores" json:"base_scores,omitempty"`
	LowScores        *uint32 `protobuf:"varint,5,opt,name=low_scores,json=lowScores" json:"low_scores,omitempty"`
	HighScors        *uint32 `protobuf:"varint,6,opt,name=high_scors,json=highScors" json:"high_scors,omitempty"`
	MinPeople        *uint32 `protobuf:"varint,7,opt,name=min_people,json=minPeople" json:"min_people,omitempty"`
	MaxPeople        *uint32 `protobuf:"varint,8,opt,name=max_people,json=maxPeople" json:"max_people,omitempty"`
	ShowPeople       *uint32 `protobuf:"varint,9,opt,name=show_people,json=showPeople" json:"show_people,omitempty"`
	RealPeople       *uint32 `protobuf:"varint,10,opt,name=real_people,json=realPeople" json:"real_people,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameLevelConfig) Reset()                    { *m = GameLevelConfig{} }
func (m *GameLevelConfig) String() string            { return proto.CompactTextString(m) }
func (*GameLevelConfig) ProtoMessage()               {}
func (*GameLevelConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GameLevelConfig) GetGameId() uint32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameLevelConfig) GetLevelId() uint32 {
	if m != nil && m.LevelId != nil {
		return *m.LevelId
	}
	return 0
}

func (m *GameLevelConfig) GetLevelName() string {
	if m != nil && m.LevelName != nil {
		return *m.LevelName
	}
	return ""
}

func (m *GameLevelConfig) GetBaseScores() uint32 {
	if m != nil && m.BaseScores != nil {
		return *m.BaseScores
	}
	return 0
}

func (m *GameLevelConfig) GetLowScores() uint32 {
	if m != nil && m.LowScores != nil {
		return *m.LowScores
	}
	return 0
}

func (m *GameLevelConfig) GetHighScors() uint32 {
	if m != nil && m.HighScors != nil {
		return *m.HighScors
	}
	return 0
}

func (m *GameLevelConfig) GetMinPeople() uint32 {
	if m != nil && m.MinPeople != nil {
		return *m.MinPeople
	}
	return 0
}

func (m *GameLevelConfig) GetMaxPeople() uint32 {
	if m != nil && m.MaxPeople != nil {
		return *m.MaxPeople
	}
	return 0
}

func (m *GameLevelConfig) GetShowPeople() uint32 {
	if m != nil && m.ShowPeople != nil {
		return *m.ShowPeople
	}
	return 0
}

func (m *GameLevelConfig) GetRealPeople() uint32 {
	if m != nil && m.RealPeople != nil {
		return *m.RealPeople
	}
	return 0
}

func init() {
	proto.RegisterType((*GeographicalLocation)(nil), "common.GeographicalLocation")
	proto.RegisterType((*GameConfig)(nil), "common.GameConfig")
	proto.RegisterType((*GameLevelConfig)(nil), "common.GameLevelConfig")
	proto.RegisterEnum("common.GameId", GameId_name, GameId_value)
	proto.RegisterEnum("common.LocSourceType", LocSourceType_name, LocSourceType_value)
	proto.RegisterEnum("common.PlayerState", PlayerState_name, PlayerState_value)
	proto.RegisterEnum("common.PlayerGender", PlayerGender_name, PlayerGender_value)
	proto.RegisterEnum("common.Platform", Platform_name, Platform_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x86, 0xb1, 0xc3, 0x97, 0xc4, 0x27, 0x04, 0x86, 0xf9, 0xa0, 0x4d, 0xff, 0x04, 0xa2, 0x1b,
	0x9a, 0x05, 0x48, 0x5d, 0x74, 0xd3, 0x6e, 0x42, 0xe2, 0x1a, 0x23, 0x27, 0x58, 0x36, 0x96, 0x10,
	0x9b, 0xd1, 0x60, 0x0f, 0x89, 0x25, 0xdb, 0x63, 0xd9, 0xe6, 0x4f, 0xbd, 0xbc, 0xde, 0x58, 0x35,
	0x67, 0x12, 0x42, 0x57, 0x5d, 0x9e, 0xe7, 0x79, 0xcf, 0x9c, 0x57, 0x51, 0x0c, 0x5b, 0xb1, 0xcc,
	0x73, 0x59, 0x9c, 0x94, 0x95, 0x6c, 0x24, 0x6d, 0xeb, 0xe9, 0xe8, 0x17, 0xec, 0x39, 0x42, 0xce,
	0x2b, 0x5e, 0x2e, 0xd2, 0x98, 0x67, 0x9e, 0x8c, 0x79, 0x93, 0xca, 0x82, 0x7e, 0x81, 0xcd, 0xe6,
	0xb9, 0x14, 0x03, 0xe3, 0xd0, 0x38, 0xde, 0xfe, 0xba, 0x7f, 0xb2, 0x5c, 0xf6, 0x64, 0x1c, 0xca,
	0xfb, 0x2a, 0x16, 0x57, 0xcf, 0xa5, 0x08, 0x30, 0x42, 0x3f, 0x82, 0x95, 0xc9, 0x62, 0x9e, 0x36,
	0xf7, 0x89, 0x18, 0x98, 0x87, 0xc6, 0xb1, 0x11, 0xac, 0x01, 0x7d, 0x0f, 0xdd, 0x8c, 0x37, 0x5a,
	0xb6, 0x50, 0xbe, 0xcc, 0x47, 0x0c, 0xc0, 0xe1, 0xb9, 0x18, 0xcb, 0xe2, 0x2e, 0x9d, 0xd3, 0xb7,
	0xd0, 0x99, 0xf3, 0x5c, 0xb0, 0x34, 0xc1, 0xab, 0xfd, 0xa0, 0xad, 0x46, 0x37, 0xa1, 0x1f, 0xc0,
	0x42, 0x51, 0xf0, 0x5c, 0x1f, 0xb0, 0x82, 0xae, 0x02, 0x33, 0x9e, 0x8b, 0x17, 0x89, 0x6d, 0x5b,
	0xb8, 0x87, 0x52, 0x15, 0x3c, 0xfa, 0x6d, 0xc2, 0x8e, 0xba, 0xe0, 0x89, 0x07, 0x91, 0xfd, 0xeb,
	0xcc, 0x3b, 0xe8, 0x66, 0x2a, 0xa7, 0x8c, 0x89, 0xa6, 0x83, 0xb3, 0x9b, 0xd0, 0x4f, 0x00, 0x5a,
	0x61, 0x85, 0x16, 0x56, 0xb0, 0x90, 0x60, 0x87, 0x03, 0xe8, 0xdd, 0xf2, 0x5a, 0xb0, 0x3a, 0x96,
	0x95, 0xa8, 0x07, 0x9b, 0xb8, 0x0c, 0x0a, 0x85, 0x48, 0x70, 0x5f, 0x3e, 0xae, 0xfc, 0x7f, 0xe8,
	0xad, 0x4c, 0x3e, 0xae, 0xf5, 0x22, 0x9d, 0x2f, 0xd0, 0xd7, 0x83, 0xb6, 0xd6, 0x8a, 0x28, 0x8f,
	0x3a, 0x4f, 0x0b, 0x56, 0x0a, 0x59, 0x66, 0x62, 0xd0, 0xd1, 0x3a, 0x4f, 0x0b, 0x1f, 0x01, 0x6a,
	0xfe, 0xb4, 0xd2, 0xdd, 0xa5, 0xe6, 0x4f, 0x4b, 0x7d, 0x00, 0xbd, 0x7a, 0x21, 0x1f, 0x57, 0xde,
	0xd2, 0xe5, 0x14, 0x5a, 0x07, 0x2a, 0xc1, 0xb3, 0x55, 0x00, 0x74, 0x40, 0x21, 0x1d, 0x18, 0x5e,
	0x43, 0xdb, 0xd1, 0x3f, 0xd1, 0x2e, 0xf4, 0x9d, 0xd1, 0xd4, 0x76, 0x27, 0xec, 0x3a, 0xb2, 0x3d,
	0x37, 0x22, 0x06, 0xa5, 0xb0, 0xbd, 0x46, 0x37, 0xe7, 0xa3, 0x19, 0x31, 0xe9, 0xff, 0xb0, 0xb3,
	0x64, 0x93, 0xcb, 0x68, 0xe2, 0xde, 0x9c, 0x47, 0xa4, 0xf5, 0x2a, 0x68, 0x07, 0x81, 0x3d, 0x9b,
	0x5e, 0x90, 0xcd, 0xe1, 0x0f, 0xe8, 0xff, 0xf5, 0x8f, 0xa2, 0x7b, 0x40, 0xbc, 0xcb, 0x31, 0x0b,
	0x2f, 0xa3, 0x60, 0x6c, 0xb3, 0xb3, 0x91, 0x3b, 0x89, 0xc8, 0x06, 0xdd, 0x87, 0xdd, 0x57, 0xf4,
	0xc2, 0xc5, 0x33, 0xc6, 0xf0, 0x3b, 0xf4, 0xfc, 0x8c, 0x3f, 0x8b, 0x2a, 0x6c, 0x78, 0x23, 0x68,
	0x0f, 0x3a, 0x7e, 0xc8, 0xdc, 0x89, 0x67, 0x93, 0x0d, 0xba, 0x0d, 0xe0, 0x87, 0x0c, 0x0f, 0xce,
	0x1c, 0x62, 0xd0, 0x1d, 0xe8, 0xf9, 0x21, 0x9b, 0x8e, 0xae, 0xc6, 0xe7, 0x0a, 0x98, 0xc3, 0x6f,
	0xb0, 0xa5, 0x97, 0x1d, 0x51, 0x24, 0xa2, 0xa2, 0x00, 0x6d, 0xdf, 0x61, 0x33, 0xd7, 0x23, 0x1b,
	0xf8, 0x92, 0xc3, 0xa6, 0x23, 0xcf, 0x26, 0x06, 0xed, 0x83, 0xe5, 0x3b, 0xec, 0xa7, 0x8d, 0xa3,
	0x39, 0xfc, 0x0c, 0x5d, 0x3f, 0xe3, 0xcd, 0x9d, 0xac, 0x72, 0x95, 0x1b, 0x15, 0x49, 0x25, 0xd3,
	0x84, 0x18, 0xea, 0x01, 0xb7, 0x5c, 0xc8, 0x42, 0x10, 0xf3, 0x6c, 0x70, 0xf3, 0xa6, 0x6e, 0xc4,
	0x83, 0x38, 0x8d, 0xb3, 0x54, 0x14, 0x0d, 0x2b, 0x6f, 0x4f, 0xf5, 0x07, 0xf4, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xf3, 0x40, 0xeb, 0x91, 0x86, 0x03, 0x00, 0x00,
}
