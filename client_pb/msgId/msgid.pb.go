// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgid.proto

/*
Package msgid is a generated protocol buffer package.

It is generated from these files:
	msgid.proto

It has these top-level messages:
*/
package msgid

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

// MsgID 消息 ID
type MsgID int32

const (
	MsgID_LOGIN_AUTH_REQ      MsgID = 1
	MsgID_LOGIN_AUTH_RSP      MsgID = 2
	MsgID_GATE_AUTH_REQ       MsgID = 4097
	MsgID_GATE_AUTH_RSP       MsgID = 4098
	MsgID_GATE_HEART_BEAT_REQ MsgID = 4099
	MsgID_GATE_HEART_BEAT_RSP MsgID = 4100
	MsgID_MATCH_REQ           MsgID = 8193
	MsgID_MATCH_RSP           MsgID = 8194
	// ROOM_BEGIN 消息区间开始
	MsgID_ROOM_BEGIN MsgID = 65536
	// ROOM_BASE_BEGIN 房间逻辑消息区间开始
	MsgID_ROOM_BASE_BEGIN           MsgID = 65537
	MsgID_ROOM_LOGIN_REQ            MsgID = 65538
	MsgID_ROOM_LOGIN_RSP            MsgID = 65539
	MsgID_ROOM_JOIN_DESK_REQ        MsgID = 65540
	MsgID_ROOM_JOIN_DESK_RSP        MsgID = 65541
	MsgID_ROOM_DESK_CREATED_NTF     MsgID = 65542
	MsgID_ROOM_DESK_QUIT_REQ        MsgID = 65543
	MsgID_ROOM_DESK_DISMISS_NTF     MsgID = 65544
	MsgID_ROOM_DESK_CONTINUE_REQ    MsgID = 65545
	MsgID_ROOM_DESK_CONTINUE_RSP    MsgID = 65546
	MsgID_ROOM_VISITOR_LOGIN_REQ    MsgID = 65547
	MsgID_ROOM_VISITOR_LOGIN_RSP    MsgID = 65548
	MsgID_ROOM_DESK_NEED_RESUME_REQ MsgID = 65549
	MsgID_ROOM_DESK_NEED_RESUME_RSP MsgID = 65550
	MsgID_ROOM_DESK_QUIT_RSP        MsgID = 65551
	MsgID_ROOM_PLAYER_LOCATION_REQ  MsgID = 65552
	MsgID_ROOM_PLAYER_LOCATION_RSP  MsgID = 65553
	MsgID_ROOM_DESK_QUIT_ENTER_NTF  MsgID = 65554
	MsgID_ROOM_CHANGE_PLAYERS_REQ   MsgID = 65556
	MsgID_ROOM_CHANGE_PLAYERS_RSP   MsgID = 65557
	MsgID_ROOM_PLAYER_GIVEUP_REQ    MsgID = 65792
	MsgID_ROOM_PLAYER_GIVEUP_RSP    MsgID = 65793
	MsgID_ROOM_PLAYER_GIVEUP_NTF    MsgID = 65794
	// ROOM_BASE_END ROOM 房间逻辑消息区间结束
	MsgID_ROOM_BASE_END MsgID = 69631
	// ROOM_GAME_BEGIN ROOM 游戏逻辑消息区间开始
	MsgID_ROOM_GAME_BEGIN     MsgID = 69632
	MsgID_ROOM_START_GAME_NTF MsgID = 69633
	MsgID_ROOM_XIPAI_NTF      MsgID = 69634
	MsgID_ROOM_FAPAI_NTF      MsgID = 69635
	// ROOM_HUANSANZHANG_NTF = 0x11020;	// 换三张通知  方向： 服务器->客户端
	MsgID_ROOM_HUANSANZHANG_REQ        MsgID = 69665
	MsgID_ROOM_HUANSANZHANG_RSP        MsgID = 69666
	MsgID_ROOM_HUANSANZHANG_FINISH_NTF MsgID = 69667
	MsgID_ROOM_DINGQUE_NTF             MsgID = 69696
	MsgID_ROOM_DINGQUE_REQ             MsgID = 69697
	MsgID_ROOM_DINGQUE_RSP             MsgID = 69698
	MsgID_ROOM_DINGQUE_FINISH_NTF      MsgID = 69699
	MsgID_ROOM_CHUPAIWENXUN_NTF        MsgID = 69728
	MsgID_ROOM_XINGPAI_ACTION_REQ      MsgID = 69744
	// ROOM_PENG_REQ = 0x11080;	// 请求碰 客户端->服务器
	MsgID_ROOM_PENG_RSP MsgID = 69761
	MsgID_ROOM_PENG_NTF MsgID = 69762
	// ROOM_GANG_REQ = 0x110A0;	// 请求杠 客户端->服务器
	MsgID_ROOM_GANG_RSP MsgID = 69793
	MsgID_ROOM_GANG_NTF MsgID = 69794
	// ROOM_HU_REQ = 0x110C0;	   // 胡请求 客户端->服务端
	MsgID_ROOM_HU_RSP MsgID = 69825
	MsgID_ROOM_HU_NTF MsgID = 69826
	// ROOM_QI_REQ = 0x110E0;	   // 请求弃 客户端->服务端
	MsgID_ROOM_QI_RSP               MsgID = 69857
	MsgID_ROOM_CHI_RSP              MsgID = 69872
	MsgID_ROOM_CHI_NTF              MsgID = 69873
	MsgID_ROOM_ZIXUN_NTF            MsgID = 69888
	MsgID_ROOM_CHUPAI_REQ           MsgID = 69889
	MsgID_ROOM_CHUPAI_NTF           MsgID = 69890
	MsgID_ROOM_MOPAI_NTF            MsgID = 69891
	MsgID_ROOM_BUHUA_NTF            MsgID = 69892
	MsgID_ROOM_WAIT_QIANGGANGHU_NTF MsgID = 69920
	MsgID_ROOM_TINGINFO_NTF         MsgID = 69921
	MsgID_ROOM_CARTOON_FINISH_REQ   MsgID = 69925
	MsgID_ROOM_GAMEOVER_NTF         MsgID = 69952
	MsgID_ROOM_RESUME_GAME_REQ      MsgID = 69969
	MsgID_ROOM_RESUME_GAME_RSP      MsgID = 69970
	// ROOM_GAME_END 游戏逻辑消息区间结束
	MsgID_ROOM_GAME_END MsgID = 73727
	// ROOM_SETTLE_BEGIN ROOM 游戏结算消息区间开始
	MsgID_ROOM_SETTLE_BEGIN   MsgID = 73728
	MsgID_ROOM_INSTANT_SETTLE MsgID = 73729
	MsgID_ROOM_ROUND_SETTLE   MsgID = 73730
	// ROOM_SETTLE_END 游戏结算消息区间结束
	MsgID_ROOM_SETTLE_END MsgID = 77823
	// 托管逻辑区间开始
	MsgID_ROOM_TUOGUAN_BEGIN      MsgID = 77824
	MsgID_ROOM_TUOGUAN_NTF        MsgID = 77825
	MsgID_ROOM_CANCEL_TUOGUAN_REQ MsgID = 77826
	// 托管逻辑区间结束
	MsgID_ROOM_TUOGUAN_END MsgID = 78079
	// 房间聊天
	MsgID_ROOM_CHAT_REQ           MsgID = 81920
	MsgID_ROOM_CHAT_NTF           MsgID = 81921
	MsgID_ROOM_DDZ_START_GAME_NTF MsgID = 86016
	MsgID_ROOM_DDZ_DEAL_NTF       MsgID = 86017
	MsgID_ROOM_DDZ_GRAB_LORD_REQ  MsgID = 86018
	MsgID_ROOM_DDZ_GRAB_LORD_RSP  MsgID = 86019
	MsgID_ROOM_DDZ_GRAB_LORD_NTF  MsgID = 86020
	MsgID_ROOM_DDZ_LORD_NTF       MsgID = 86021
	MsgID_ROOM_DDZ_DOUBLE_REQ     MsgID = 86022
	MsgID_ROOM_DDZ_DOUBLE_RSP     MsgID = 86023
	MsgID_ROOM_DDZ_DOUBLE_NTF     MsgID = 86024
	MsgID_ROOM_DDZ_PLAY_CARD_REQ  MsgID = 86025
	MsgID_ROOM_DDZ_PLAY_CARD_RSP  MsgID = 86026
	MsgID_ROOM_DDZ_PLAY_CARD_NTF  MsgID = 86027
	MsgID_ROOM_DDZ_GAME_OVER_NTF  MsgID = 86028
	MsgID_ROOM_DDZ_TUOGUAN_REQ    MsgID = 86029
	MsgID_ROOM_DDZ_TUOGUAN_RSP    MsgID = 86030
	MsgID_ROOM_DDZ_TUOGUAN_NTF    MsgID = 86031
	MsgID_ROOM_DDZ_RESUME_REQ     MsgID = 86032
	MsgID_ROOM_DDZ_RESUME_RSP     MsgID = 86033
	// ROOM END 房间消息区间结束
	MsgID_ROOM_END MsgID = 131071
)

var MsgID_name = map[int32]string{
	1:      "LOGIN_AUTH_REQ",
	2:      "LOGIN_AUTH_RSP",
	4097:   "GATE_AUTH_REQ",
	4098:   "GATE_AUTH_RSP",
	4099:   "GATE_HEART_BEAT_REQ",
	4100:   "GATE_HEART_BEAT_RSP",
	8193:   "MATCH_REQ",
	8194:   "MATCH_RSP",
	65536:  "ROOM_BEGIN",
	65537:  "ROOM_BASE_BEGIN",
	65538:  "ROOM_LOGIN_REQ",
	65539:  "ROOM_LOGIN_RSP",
	65540:  "ROOM_JOIN_DESK_REQ",
	65541:  "ROOM_JOIN_DESK_RSP",
	65542:  "ROOM_DESK_CREATED_NTF",
	65543:  "ROOM_DESK_QUIT_REQ",
	65544:  "ROOM_DESK_DISMISS_NTF",
	65545:  "ROOM_DESK_CONTINUE_REQ",
	65546:  "ROOM_DESK_CONTINUE_RSP",
	65547:  "ROOM_VISITOR_LOGIN_REQ",
	65548:  "ROOM_VISITOR_LOGIN_RSP",
	65549:  "ROOM_DESK_NEED_RESUME_REQ",
	65550:  "ROOM_DESK_NEED_RESUME_RSP",
	65551:  "ROOM_DESK_QUIT_RSP",
	65552:  "ROOM_PLAYER_LOCATION_REQ",
	65553:  "ROOM_PLAYER_LOCATION_RSP",
	65554:  "ROOM_DESK_QUIT_ENTER_NTF",
	65556:  "ROOM_CHANGE_PLAYERS_REQ",
	65557:  "ROOM_CHANGE_PLAYERS_RSP",
	65792:  "ROOM_PLAYER_GIVEUP_REQ",
	65793:  "ROOM_PLAYER_GIVEUP_RSP",
	65794:  "ROOM_PLAYER_GIVEUP_NTF",
	69631:  "ROOM_BASE_END",
	69632:  "ROOM_GAME_BEGIN",
	69633:  "ROOM_START_GAME_NTF",
	69634:  "ROOM_XIPAI_NTF",
	69635:  "ROOM_FAPAI_NTF",
	69665:  "ROOM_HUANSANZHANG_REQ",
	69666:  "ROOM_HUANSANZHANG_RSP",
	69667:  "ROOM_HUANSANZHANG_FINISH_NTF",
	69696:  "ROOM_DINGQUE_NTF",
	69697:  "ROOM_DINGQUE_REQ",
	69698:  "ROOM_DINGQUE_RSP",
	69699:  "ROOM_DINGQUE_FINISH_NTF",
	69728:  "ROOM_CHUPAIWENXUN_NTF",
	69744:  "ROOM_XINGPAI_ACTION_REQ",
	69761:  "ROOM_PENG_RSP",
	69762:  "ROOM_PENG_NTF",
	69793:  "ROOM_GANG_RSP",
	69794:  "ROOM_GANG_NTF",
	69825:  "ROOM_HU_RSP",
	69826:  "ROOM_HU_NTF",
	69857:  "ROOM_QI_RSP",
	69872:  "ROOM_CHI_RSP",
	69873:  "ROOM_CHI_NTF",
	69888:  "ROOM_ZIXUN_NTF",
	69889:  "ROOM_CHUPAI_REQ",
	69890:  "ROOM_CHUPAI_NTF",
	69891:  "ROOM_MOPAI_NTF",
	69892:  "ROOM_BUHUA_NTF",
	69920:  "ROOM_WAIT_QIANGGANGHU_NTF",
	69921:  "ROOM_TINGINFO_NTF",
	69925:  "ROOM_CARTOON_FINISH_REQ",
	69952:  "ROOM_GAMEOVER_NTF",
	69969:  "ROOM_RESUME_GAME_REQ",
	69970:  "ROOM_RESUME_GAME_RSP",
	73727:  "ROOM_GAME_END",
	73728:  "ROOM_SETTLE_BEGIN",
	73729:  "ROOM_INSTANT_SETTLE",
	73730:  "ROOM_ROUND_SETTLE",
	77823:  "ROOM_SETTLE_END",
	77824:  "ROOM_TUOGUAN_BEGIN",
	77825:  "ROOM_TUOGUAN_NTF",
	77826:  "ROOM_CANCEL_TUOGUAN_REQ",
	78079:  "ROOM_TUOGUAN_END",
	81920:  "ROOM_CHAT_REQ",
	81921:  "ROOM_CHAT_NTF",
	86016:  "ROOM_DDZ_START_GAME_NTF",
	86017:  "ROOM_DDZ_DEAL_NTF",
	86018:  "ROOM_DDZ_GRAB_LORD_REQ",
	86019:  "ROOM_DDZ_GRAB_LORD_RSP",
	86020:  "ROOM_DDZ_GRAB_LORD_NTF",
	86021:  "ROOM_DDZ_LORD_NTF",
	86022:  "ROOM_DDZ_DOUBLE_REQ",
	86023:  "ROOM_DDZ_DOUBLE_RSP",
	86024:  "ROOM_DDZ_DOUBLE_NTF",
	86025:  "ROOM_DDZ_PLAY_CARD_REQ",
	86026:  "ROOM_DDZ_PLAY_CARD_RSP",
	86027:  "ROOM_DDZ_PLAY_CARD_NTF",
	86028:  "ROOM_DDZ_GAME_OVER_NTF",
	86029:  "ROOM_DDZ_TUOGUAN_REQ",
	86030:  "ROOM_DDZ_TUOGUAN_RSP",
	86031:  "ROOM_DDZ_TUOGUAN_NTF",
	86032:  "ROOM_DDZ_RESUME_REQ",
	86033:  "ROOM_DDZ_RESUME_RSP",
	131071: "ROOM_END",
}
var MsgID_value = map[string]int32{
	"LOGIN_AUTH_REQ":               1,
	"LOGIN_AUTH_RSP":               2,
	"GATE_AUTH_REQ":                4097,
	"GATE_AUTH_RSP":                4098,
	"GATE_HEART_BEAT_REQ":          4099,
	"GATE_HEART_BEAT_RSP":          4100,
	"MATCH_REQ":                    8193,
	"MATCH_RSP":                    8194,
	"ROOM_BEGIN":                   65536,
	"ROOM_BASE_BEGIN":              65537,
	"ROOM_LOGIN_REQ":               65538,
	"ROOM_LOGIN_RSP":               65539,
	"ROOM_JOIN_DESK_REQ":           65540,
	"ROOM_JOIN_DESK_RSP":           65541,
	"ROOM_DESK_CREATED_NTF":        65542,
	"ROOM_DESK_QUIT_REQ":           65543,
	"ROOM_DESK_DISMISS_NTF":        65544,
	"ROOM_DESK_CONTINUE_REQ":       65545,
	"ROOM_DESK_CONTINUE_RSP":       65546,
	"ROOM_VISITOR_LOGIN_REQ":       65547,
	"ROOM_VISITOR_LOGIN_RSP":       65548,
	"ROOM_DESK_NEED_RESUME_REQ":    65549,
	"ROOM_DESK_NEED_RESUME_RSP":    65550,
	"ROOM_DESK_QUIT_RSP":           65551,
	"ROOM_PLAYER_LOCATION_REQ":     65552,
	"ROOM_PLAYER_LOCATION_RSP":     65553,
	"ROOM_DESK_QUIT_ENTER_NTF":     65554,
	"ROOM_CHANGE_PLAYERS_REQ":      65556,
	"ROOM_CHANGE_PLAYERS_RSP":      65557,
	"ROOM_PLAYER_GIVEUP_REQ":       65792,
	"ROOM_PLAYER_GIVEUP_RSP":       65793,
	"ROOM_PLAYER_GIVEUP_NTF":       65794,
	"ROOM_BASE_END":                69631,
	"ROOM_GAME_BEGIN":              69632,
	"ROOM_START_GAME_NTF":          69633,
	"ROOM_XIPAI_NTF":               69634,
	"ROOM_FAPAI_NTF":               69635,
	"ROOM_HUANSANZHANG_REQ":        69665,
	"ROOM_HUANSANZHANG_RSP":        69666,
	"ROOM_HUANSANZHANG_FINISH_NTF": 69667,
	"ROOM_DINGQUE_NTF":             69696,
	"ROOM_DINGQUE_REQ":             69697,
	"ROOM_DINGQUE_RSP":             69698,
	"ROOM_DINGQUE_FINISH_NTF":      69699,
	"ROOM_CHUPAIWENXUN_NTF":        69728,
	"ROOM_XINGPAI_ACTION_REQ":      69744,
	"ROOM_PENG_RSP":                69761,
	"ROOM_PENG_NTF":                69762,
	"ROOM_GANG_RSP":                69793,
	"ROOM_GANG_NTF":                69794,
	"ROOM_HU_RSP":                  69825,
	"ROOM_HU_NTF":                  69826,
	"ROOM_QI_RSP":                  69857,
	"ROOM_CHI_RSP":                 69872,
	"ROOM_CHI_NTF":                 69873,
	"ROOM_ZIXUN_NTF":               69888,
	"ROOM_CHUPAI_REQ":              69889,
	"ROOM_CHUPAI_NTF":              69890,
	"ROOM_MOPAI_NTF":               69891,
	"ROOM_BUHUA_NTF":               69892,
	"ROOM_WAIT_QIANGGANGHU_NTF":    69920,
	"ROOM_TINGINFO_NTF":            69921,
	"ROOM_CARTOON_FINISH_REQ":      69925,
	"ROOM_GAMEOVER_NTF":            69952,
	"ROOM_RESUME_GAME_REQ":         69969,
	"ROOM_RESUME_GAME_RSP":         69970,
	"ROOM_GAME_END":                73727,
	"ROOM_SETTLE_BEGIN":            73728,
	"ROOM_INSTANT_SETTLE":          73729,
	"ROOM_ROUND_SETTLE":            73730,
	"ROOM_SETTLE_END":              77823,
	"ROOM_TUOGUAN_BEGIN":           77824,
	"ROOM_TUOGUAN_NTF":             77825,
	"ROOM_CANCEL_TUOGUAN_REQ":      77826,
	"ROOM_TUOGUAN_END":             78079,
	"ROOM_CHAT_REQ":                81920,
	"ROOM_CHAT_NTF":                81921,
	"ROOM_DDZ_START_GAME_NTF":      86016,
	"ROOM_DDZ_DEAL_NTF":            86017,
	"ROOM_DDZ_GRAB_LORD_REQ":       86018,
	"ROOM_DDZ_GRAB_LORD_RSP":       86019,
	"ROOM_DDZ_GRAB_LORD_NTF":       86020,
	"ROOM_DDZ_LORD_NTF":            86021,
	"ROOM_DDZ_DOUBLE_REQ":          86022,
	"ROOM_DDZ_DOUBLE_RSP":          86023,
	"ROOM_DDZ_DOUBLE_NTF":          86024,
	"ROOM_DDZ_PLAY_CARD_REQ":       86025,
	"ROOM_DDZ_PLAY_CARD_RSP":       86026,
	"ROOM_DDZ_PLAY_CARD_NTF":       86027,
	"ROOM_DDZ_GAME_OVER_NTF":       86028,
	"ROOM_DDZ_TUOGUAN_REQ":         86029,
	"ROOM_DDZ_TUOGUAN_RSP":         86030,
	"ROOM_DDZ_TUOGUAN_NTF":         86031,
	"ROOM_DDZ_RESUME_REQ":          86032,
	"ROOM_DDZ_RESUME_RSP":          86033,
	"ROOM_END":                     131071,
}

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}
func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}
func (x *MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgID_value, data, "MsgID")
	if err != nil {
		return err
	}
	*x = MsgID(value)
	return nil
}
func (MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("msgid.MsgID", MsgID_name, MsgID_value)
}

func init() { proto.RegisterFile("msgid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 930 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0x4d, 0x92, 0xdc, 0x34,
	0x14, 0xc7, 0x0b, 0xca, 0x2e, 0x40, 0x21, 0x83, 0xa2, 0xc9, 0xc7, 0x04, 0x02, 0x4c, 0xb1, 0x64,
	0xc1, 0x19, 0xa2, 0x6e, 0x6b, 0xdc, 0x82, 0x6e, 0x49, 0x6d, 0xc9, 0x93, 0x61, 0x36, 0xde, 0x50,
	0x95, 0x62, 0x41, 0x85, 0x22, 0x1c, 0x40, 0xdd, 0x33, 0x49, 0x08, 0x49, 0x48, 0xa0, 0x60, 0x8d,
	0xdb, 0x14, 0x07, 0xe0, 0x04, 0x74, 0xc2, 0x09, 0xe0, 0x02, 0x4c, 0x6e, 0x10, 0x4e, 0x10, 0x4a,
	0xcf, 0x1f, 0x52, 0x3c, 0xed, 0x59, 0xfa, 0xff, 0x7b, 0x1f, 0x7a, 0x4f, 0x4f, 0x7e, 0xe8, 0xdc,
	0x57, 0xb7, 0x6f, 0x7e, 0xf9, 0xc5, 0x27, 0x5f, 0x7f, 0x73, 0xeb, 0xdb, 0x5b, 0x24, 0x86, 0x8f,
	0x8f, 0xff, 0xd8, 0x46, 0xf1, 0xec, 0xf6, 0x4d, 0x9e, 0x10, 0x82, 0xb6, 0xa6, 0x32, 0xe5, 0xa2,
	0xa0, 0xb9, 0x99, 0x14, 0x19, 0x9b, 0xe3, 0xd7, 0xfa, 0x9a, 0x56, 0xf8, 0x75, 0x42, 0xd0, 0xf9,
	0x94, 0x1a, 0xe6, 0xcd, 0x16, 0xbb, 0x3d, 0x4d, 0x2b, 0xbc, 0xdc, 0x25, 0x3b, 0x68, 0x1b, 0xb4,
	0x09, 0xa3, 0x99, 0x29, 0x46, 0x8c, 0x1a, 0xb0, 0x3e, 0xda, 0x4c, 0xb4, 0xc2, 0xc7, 0xbb, 0x64,
	0x0b, 0xbd, 0x35, 0xa3, 0x66, 0xdc, 0xc4, 0xbd, 0x1e, 0x7c, 0xbb, 0x98, 0xd7, 0x09, 0x46, 0x28,
	0x93, 0x72, 0x56, 0x8c, 0x58, 0xca, 0x05, 0xb6, 0x36, 0x22, 0x97, 0xd0, 0x3b, 0xb5, 0x42, 0x35,
	0x6b, 0xe4, 0x85, 0x8d, 0xc8, 0x45, 0xb4, 0x05, 0x72, 0x7d, 0x7a, 0x17, 0x6d, 0x79, 0x5a, 0xd5,
	0x0a, 0x1f, 0xd9, 0x88, 0xec, 0x20, 0x02, 0xea, 0xa7, 0x92, 0x8b, 0x22, 0x61, 0xfa, 0x33, 0xb0,
	0x3f, 0xde, 0x4c, 0xb4, 0xc2, 0x77, 0x6c, 0x44, 0xde, 0x43, 0x97, 0x80, 0x80, 0x38, 0xce, 0x18,
	0x35, 0x2c, 0x29, 0x84, 0xd9, 0xc3, 0x77, 0x03, 0x37, 0x80, 0xf3, 0x9c, 0xd7, 0x85, 0xdf, 0xeb,
	0xbb, 0x25, 0x5c, 0xcf, 0xb8, 0xd6, 0xe0, 0xf6, 0x9d, 0x8d, 0xc8, 0x35, 0x74, 0x39, 0x88, 0x29,
	0x85, 0xe1, 0x22, 0x67, 0xe0, 0x7a, 0x7f, 0x98, 0x6a, 0x85, 0xbf, 0x0f, 0xe8, 0x3e, 0xd7, 0xdc,
	0xc8, 0x2c, 0xa8, 0xfb, 0xc1, 0x30, 0xd5, 0x0a, 0x3f, 0xb4, 0x11, 0xf9, 0x10, 0x5d, 0xf5, 0x91,
	0x05, 0x63, 0x49, 0x91, 0x31, 0x9d, 0xcf, 0xea, 0xd4, 0x8f, 0xce, 0x34, 0xd0, 0x0a, 0xff, 0xb0,
	0xb9, 0x60, 0xad, 0xf0, 0x63, 0x1b, 0x91, 0x0f, 0xd0, 0x0e, 0x10, 0x35, 0xa5, 0x9f, 0x33, 0x97,
	0x78, 0x4c, 0x0d, 0x97, 0xf5, 0xc9, 0x9e, 0x9c, 0xc5, 0xb5, 0xc2, 0x3f, 0x06, 0xdc, 0x47, 0x66,
	0xc2, 0xb0, 0x0c, 0x7a, 0xf6, 0x93, 0x8d, 0xc8, 0xfb, 0xe8, 0x0a, 0xf0, 0xf1, 0x84, 0x8a, 0x94,
	0x35, 0x61, 0x34, 0x84, 0xff, 0xf9, 0x0c, 0xac, 0x15, 0xfe, 0x25, 0xe8, 0x4b, 0x93, 0x3d, 0xe5,
	0xfb, 0x2c, 0x57, 0xe0, 0x6c, 0x97, 0x83, 0x54, 0x2b, 0xbc, 0x18, 0xa4, 0xee, 0x5c, 0xcb, 0x65,
	0x44, 0xb6, 0xd1, 0x79, 0x3f, 0x96, 0x4c, 0x24, 0xf8, 0xe5, 0xaf, 0x7e, 0x56, 0x53, 0x3a, 0x6b,
	0x67, 0xd5, 0x96, 0x11, 0xb9, 0x8a, 0xb6, 0x41, 0xd6, 0xc6, 0x3d, 0x07, 0x80, 0x2e, 0xcc, 0xa2,
	0xf4, 0x03, 0x7b, 0xc0, 0x15, 0xe5, 0x75, 0xf0, 0x40, 0xdd, 0xa3, 0xad, 0x7a, 0x54, 0xfa, 0xd9,
	0x9a, 0xe4, 0x54, 0x68, 0x2a, 0x0e, 0x5d, 0xc9, 0x50, 0xcb, 0x6a, 0x10, 0x6a, 0x85, 0xab, 0x32,
	0x22, 0x1f, 0xa1, 0x6b, 0xa7, 0xe1, 0x1e, 0x17, 0x5c, 0x4f, 0x20, 0xfa, 0x6f, 0x65, 0x44, 0x2e,
	0x23, 0x5c, 0x5f, 0x04, 0x17, 0xe9, 0x3c, 0xaf, 0x4f, 0xb8, 0xde, 0xa0, 0xbb, 0x84, 0x4f, 0x37,
	0xe9, 0x5a, 0xe1, 0x67, 0xa5, 0xbf, 0x91, 0x56, 0x0f, 0xd2, 0xfc, 0x15, 0x9c, 0x73, 0x3c, 0xc9,
	0x15, 0xe5, 0x37, 0x98, 0x38, 0xc8, 0x05, 0xc0, 0x93, 0xc0, 0xf7, 0x80, 0x8b, 0xd4, 0x55, 0x4e,
	0xc7, 0xdd, 0x2c, 0xbd, 0x28, 0x7d, 0xcf, 0x15, 0x6b, 0x6a, 0x5b, 0xac, 0x7a, 0x22, 0x34, 0x30,
	0x10, 0xd3, 0xb6, 0x0b, 0xab, 0xbe, 0xe8, 0x2c, 0xab, 0x55, 0x44, 0x2e, 0xa0, 0x73, 0x4d, 0x6b,
	0xc0, 0xee, 0xe9, 0xab, 0x92, 0xb3, 0x7a, 0x16, 0x48, 0x73, 0x0e, 0x56, 0xcf, 0x57, 0x11, 0x21,
	0xe8, 0xed, 0xa6, 0x90, 0x5a, 0x7b, 0xd1, 0xd3, 0x9c, 0xeb, 0x7f, 0x2b, 0x7f, 0x97, 0x87, 0xbc,
	0xad, 0xd4, 0x56, 0x7e, 0x52, 0xea, 0x36, 0xd4, 0x7f, 0xc3, 0xd3, 0x32, 0x94, 0x53, 0xf9, 0x18,
	0x33, 0xd9, 0xcd, 0x43, 0xa0, 0x8e, 0xf2, 0x49, 0x4e, 0x41, 0x3d, 0xae, 0xfc, 0x5b, 0xbe, 0x41,
	0xb9, 0x29, 0xe6, 0x9c, 0x8a, 0xd4, 0xd5, 0xdb, 0xd4, 0x52, 0x56, 0x11, 0xb9, 0x82, 0x2e, 0x80,
	0x81, 0xe1, 0x22, 0xe5, 0x62, 0x4f, 0x02, 0x58, 0x55, 0xc1, 0x5b, 0xa2, 0x99, 0x91, 0x52, 0xb4,
	0x37, 0xe7, 0xce, 0xf6, 0x7b, 0xe0, 0xe7, 0xe6, 0x57, 0xee, 0x37, 0x4f, 0x74, 0x5d, 0x45, 0xe4,
	0x5d, 0x74, 0x11, 0x40, 0xf3, 0xcf, 0x80, 0xf9, 0x76, 0x4e, 0x7f, 0x0f, 0x31, 0xad, 0xf0, 0x3f,
	0x55, 0x78, 0x1f, 0xb3, 0xe6, 0x09, 0xfd, 0xe9, 0xb3, 0x68, 0x66, 0xcc, 0xb4, 0x7b, 0x44, 0x6b,
	0xff, 0x88, 0xb8, 0xd0, 0x86, 0x0a, 0xd3, 0x18, 0xe0, 0xc5, 0xda, 0xfb, 0x64, 0x32, 0x17, 0x49,
	0x0b, 0x96, 0x6b, 0xdf, 0xce, 0x26, 0x18, 0xe4, 0xf8, 0xd7, 0xff, 0xcd, 0x4c, 0x2e, 0xd3, 0x9c,
	0x8a, 0x36, 0xc9, 0x89, 0x1f, 0xea, 0x96, 0xc0, 0x33, 0x3d, 0x09, 0x5b, 0x23, 0xc6, 0x6c, 0xda,
	0x61, 0x58, 0x3b, 0x1b, 0xdc, 0x20, 0xd1, 0x73, 0x5f, 0xe1, 0x78, 0xd2, 0xec, 0x46, 0x6b, 0xe3,
	0x57, 0x45, 0x48, 0x60, 0x63, 0xff, 0x6a, 0x92, 0xc3, 0xfe, 0x6f, 0xc2, 0x96, 0x71, 0x57, 0xa1,
	0xc3, 0x09, 0xa3, 0xd3, 0xe6, 0xff, 0x11, 0xfb, 0xa5, 0x91, 0x1c, 0x16, 0x69, 0x46, 0x47, 0xc5,
	0x54, 0x66, 0x49, 0x7d, 0xae, 0x61, 0xea, 0xd6, 0xe2, 0x20, 0x85, 0x39, 0xea, 0xa5, 0xec, 0xc0,
	0x9d, 0x32, 0xee, 0x2e, 0x02, 0xce, 0x22, 0xf3, 0xd1, 0xb4, 0xbe, 0xed, 0xbb, 0x03, 0x48, 0x2b,
	0x7c, 0x6f, 0x33, 0x82, 0xb5, 0xd8, 0x3b, 0x87, 0xfb, 0xd9, 0xba, 0x01, 0xac, 0x6b, 0xb8, 0x3f,
	0x4c, 0xdd, 0x5a, 0x1c, 0xa4, 0x2e, 0xf2, 0x83, 0x7e, 0x85, 0xae, 0x9f, 0xdd, 0xdc, 0x3e, 0x2c,
	0xe3, 0x6e, 0x36, 0x1d, 0x0d, 0x6f, 0xf4, 0xd1, 0x10, 0x73, 0xcb, 0x70, 0x80, 0xb9, 0x98, 0x8f,
	0x7b, 0x65, 0x06, 0x4b, 0xf6, 0xc9, 0x00, 0x72, 0x4b, 0xb0, 0x8c, 0xc9, 0x16, 0x7a, 0x13, 0x10,
	0xcc, 0xcd, 0xcb, 0x37, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x45, 0x35, 0xbc, 0x75, 0xc8, 0x09,
	0x00, 0x00,
}
