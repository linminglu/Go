// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_ddz.proto

package room

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DDZStage:当前游戏状态(恢复对局用)
type DDZStage int32

const (
	DDZStage_DDZ_STAGE_NONE    DDZStage = 0
	DDZStage_DDZ_STAGE_DEAL    DDZStage = 1
	DDZStage_DDZ_STAGE_CALL    DDZStage = 2
	DDZStage_DDZ_STAGE_GRAB    DDZStage = 3
	DDZStage_DDZ_STAGE_DOUBLE  DDZStage = 4
	DDZStage_DDZ_STAGE_PLAYING DDZStage = 5
	DDZStage_DDZ_STAGE_OVER    DDZStage = 6
)

var DDZStage_name = map[int32]string{
	0: "DDZ_STAGE_NONE",
	1: "DDZ_STAGE_DEAL",
	2: "DDZ_STAGE_CALL",
	3: "DDZ_STAGE_GRAB",
	4: "DDZ_STAGE_DOUBLE",
	5: "DDZ_STAGE_PLAYING",
	6: "DDZ_STAGE_OVER",
}
var DDZStage_value = map[string]int32{
	"DDZ_STAGE_NONE":    0,
	"DDZ_STAGE_DEAL":    1,
	"DDZ_STAGE_CALL":    2,
	"DDZ_STAGE_GRAB":    3,
	"DDZ_STAGE_DOUBLE":  4,
	"DDZ_STAGE_PLAYING": 5,
	"DDZ_STAGE_OVER":    6,
}

func (x DDZStage) Enum() *DDZStage {
	p := new(DDZStage)
	*p = x
	return p
}
func (x DDZStage) String() string {
	return proto.EnumName(DDZStage_name, int32(x))
}
func (x *DDZStage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DDZStage_value, data, "DDZStage")
	if err != nil {
		return err
	}
	*x = DDZStage(value)
	return nil
}
func (DDZStage) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type CardType int32

const (
	CardType_CT_NONE     CardType = 0
	CardType_CT_SINGLE   CardType = 1
	CardType_CT_PAIR     CardType = 2
	CardType_CT_SHUNZI   CardType = 3
	CardType_CT_PAIRS    CardType = 4
	CardType_CT_TRIPLE   CardType = 5
	CardType_CT_3AND1    CardType = 6
	CardType_CT_3AND2    CardType = 7
	CardType_CT_TRIPLES  CardType = 8
	CardType_CT_3SAND1S  CardType = 9
	CardType_CT_3SAND2S  CardType = 10
	CardType_CT_4SAND1S  CardType = 11
	CardType_CT_4SAND2S  CardType = 12
	CardType_CT_BOMB     CardType = 13
	CardType_CT_KINGBOMB CardType = 14
)

var CardType_name = map[int32]string{
	0:  "CT_NONE",
	1:  "CT_SINGLE",
	2:  "CT_PAIR",
	3:  "CT_SHUNZI",
	4:  "CT_PAIRS",
	5:  "CT_TRIPLE",
	6:  "CT_3AND1",
	7:  "CT_3AND2",
	8:  "CT_TRIPLES",
	9:  "CT_3SAND1S",
	10: "CT_3SAND2S",
	11: "CT_4SAND1S",
	12: "CT_4SAND2S",
	13: "CT_BOMB",
	14: "CT_KINGBOMB",
}
var CardType_value = map[string]int32{
	"CT_NONE":     0,
	"CT_SINGLE":   1,
	"CT_PAIR":     2,
	"CT_SHUNZI":   3,
	"CT_PAIRS":    4,
	"CT_TRIPLE":   5,
	"CT_3AND1":    6,
	"CT_3AND2":    7,
	"CT_TRIPLES":  8,
	"CT_3SAND1S":  9,
	"CT_3SAND2S":  10,
	"CT_4SAND1S":  11,
	"CT_4SAND2S":  12,
	"CT_BOMB":     13,
	"CT_KINGBOMB": 14,
}

func (x CardType) Enum() *CardType {
	p := new(CardType)
	*p = x
	return p
}
func (x CardType) String() string {
	return proto.EnumName(CardType_name, int32(x))
}
func (x *CardType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CardType_value, data, "CardType")
	if err != nil {
		return err
	}
	*x = CardType(value)
	return nil
}
func (CardType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

// 通用操作结果，客户端直接弹出err_desc提示即可
type Result struct {
	ErrCode          *uint32 `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrDesc          *string `protobuf:"bytes,2,opt,name=err_desc,json=errDesc" json:"err_desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Result) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Result) GetErrDesc() string {
	if m != nil && m.ErrDesc != nil {
		return *m.ErrDesc
	}
	return ""
}

// “Stage切换”可以独立一条消息，也可以嵌入到其他会发生阶段切换的消息体中
type NextStage struct {
	Stage            *DDZStage `protobuf:"varint,1,opt,name=stage,enum=room.DDZStage" json:"stage,omitempty"`
	Time             *uint32   `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NextStage) Reset()                    { *m = NextStage{} }
func (m *NextStage) String() string            { return proto.CompactTextString(m) }
func (*NextStage) ProtoMessage()               {}
func (*NextStage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *NextStage) GetStage() DDZStage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return DDZStage_DDZ_STAGE_NONE
}

func (m *NextStage) GetTime() uint32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

// DDZStartGameNtf 斗地主游戏开始通知
// CMD: 0x15000
type DDZStartGameNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,2,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZStartGameNtf) Reset()                    { *m = DDZStartGameNtf{} }
func (m *DDZStartGameNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZStartGameNtf) ProtoMessage()               {}
func (*DDZStartGameNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DDZStartGameNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZStartGameNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZDealNtf 发牌通知
// CMD: 0x15001
type DDZDealNtf struct {
	Cards            []uint32   `protobuf:"varint,1,rep,name=cards" json:"cards,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,2,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZDealNtf) Reset()                    { *m = DDZDealNtf{} }
func (m *DDZDealNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZDealNtf) ProtoMessage()               {}
func (*DDZDealNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *DDZDealNtf) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DDZDealNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZGrabLordReq 叫/抢地主请求，叫地主和抢地主用同一个请求
// CMD: 0x15002
type DDZGrabLordReq struct {
	Grab             *bool  `protobuf:"varint,1,opt,name=grab" json:"grab,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DDZGrabLordReq) Reset()                    { *m = DDZGrabLordReq{} }
func (m *DDZGrabLordReq) String() string            { return proto.CompactTextString(m) }
func (*DDZGrabLordReq) ProtoMessage()               {}
func (*DDZGrabLordReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *DDZGrabLordReq) GetGrab() bool {
	if m != nil && m.Grab != nil {
		return *m.Grab
	}
	return false
}

// DDZGrabLordNtf 叫/抢地主广播
// CMD: 0x15004
type DDZGrabLordNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Grab             *bool      `protobuf:"varint,2,opt,name=grab" json:"grab,omitempty"`
	NextPlayerId     *uint64    `protobuf:"varint,3,opt,name=next_player_id,json=nextPlayerId" json:"next_player_id,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,4,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZGrabLordNtf) Reset()                    { *m = DDZGrabLordNtf{} }
func (m *DDZGrabLordNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZGrabLordNtf) ProtoMessage()               {}
func (*DDZGrabLordNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *DDZGrabLordNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZGrabLordNtf) GetGrab() bool {
	if m != nil && m.Grab != nil {
		return *m.Grab
	}
	return false
}

func (m *DDZGrabLordNtf) GetNextPlayerId() uint64 {
	if m != nil && m.NextPlayerId != nil {
		return *m.NextPlayerId
	}
	return 0
}

func (m *DDZGrabLordNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZLordNtf 地主广播
// CMD: 0x15005
type DDZLordNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	TotalGrab        *uint32    `protobuf:"varint,2,opt,name=total_grab,json=totalGrab" json:"total_grab,omitempty"`
	Dipai            []uint32   `protobuf:"varint,3,rep,name=dipai" json:"dipai,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,4,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZLordNtf) Reset()                    { *m = DDZLordNtf{} }
func (m *DDZLordNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZLordNtf) ProtoMessage()               {}
func (*DDZLordNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *DDZLordNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZLordNtf) GetTotalGrab() uint32 {
	if m != nil && m.TotalGrab != nil {
		return *m.TotalGrab
	}
	return 0
}

func (m *DDZLordNtf) GetDipai() []uint32 {
	if m != nil {
		return m.Dipai
	}
	return nil
}

func (m *DDZLordNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZDoubleReq 斗地主加倍请求
// CMD: 0x15006
type DDZDoubleReq struct {
	IsDouble         *bool  `protobuf:"varint,1,opt,name=is_double,json=isDouble" json:"is_double,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DDZDoubleReq) Reset()                    { *m = DDZDoubleReq{} }
func (m *DDZDoubleReq) String() string            { return proto.CompactTextString(m) }
func (*DDZDoubleReq) ProtoMessage()               {}
func (*DDZDoubleReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *DDZDoubleReq) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return false
}

// DDZDoubleNtf 加倍广播
// CMD: 0x15008
type DDZDoubleNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	IsDouble         *bool      `protobuf:"varint,2,opt,name=is_double,json=isDouble" json:"is_double,omitempty"`
	TotalDouble      *uint32    `protobuf:"varint,3,opt,name=total_double,json=totalDouble" json:"total_double,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,4,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZDoubleNtf) Reset()                    { *m = DDZDoubleNtf{} }
func (m *DDZDoubleNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZDoubleNtf) ProtoMessage()               {}
func (*DDZDoubleNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *DDZDoubleNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZDoubleNtf) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return false
}

func (m *DDZDoubleNtf) GetTotalDouble() uint32 {
	if m != nil && m.TotalDouble != nil {
		return *m.TotalDouble
	}
	return 0
}

func (m *DDZDoubleNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZPlayCardReq 出牌请求
// CMD: 0x15009
type DDZPlayCardReq struct {
	Cards            []uint32  `protobuf:"varint,1,rep,name=cards" json:"cards,omitempty"`
	CardType         *CardType `protobuf:"varint,2,opt,name=card_type,json=cardType,enum=room.CardType" json:"card_type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *DDZPlayCardReq) Reset()                    { *m = DDZPlayCardReq{} }
func (m *DDZPlayCardReq) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayCardReq) ProtoMessage()               {}
func (*DDZPlayCardReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *DDZPlayCardReq) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DDZPlayCardReq) GetCardType() CardType {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return CardType_CT_NONE
}

// DDZPlayCardRsp 出牌响应
// CMD: 0x1500A
type DDZPlayCardRsp struct {
	Result           *Result `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DDZPlayCardRsp) Reset()                    { *m = DDZPlayCardRsp{} }
func (m *DDZPlayCardRsp) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayCardRsp) ProtoMessage()               {}
func (*DDZPlayCardRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *DDZPlayCardRsp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

// DDZPlayCardNtf 出牌广播
// CMD: 0x1500B
type DDZPlayCardNtf struct {
	PlayerId         *uint64    `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Cards            []uint32   `protobuf:"varint,2,rep,name=cards" json:"cards,omitempty"`
	CardType         *CardType  `protobuf:"varint,3,opt,name=card_type,json=cardType,enum=room.CardType" json:"card_type,omitempty"`
	TotalBomb        *uint32    `protobuf:"varint,4,opt,name=total_bomb,json=totalBomb" json:"total_bomb,omitempty"`
	NextPlayerId     *uint64    `protobuf:"varint,5,opt,name=next_player_id,json=nextPlayerId" json:"next_player_id,omitempty"`
	NextStage        *NextStage `protobuf:"bytes,6,opt,name=next_stage,json=nextStage" json:"next_stage,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DDZPlayCardNtf) Reset()                    { *m = DDZPlayCardNtf{} }
func (m *DDZPlayCardNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayCardNtf) ProtoMessage()               {}
func (*DDZPlayCardNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *DDZPlayCardNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZPlayCardNtf) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *DDZPlayCardNtf) GetCardType() CardType {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return CardType_CT_NONE
}

func (m *DDZPlayCardNtf) GetTotalBomb() uint32 {
	if m != nil && m.TotalBomb != nil {
		return *m.TotalBomb
	}
	return 0
}

func (m *DDZPlayCardNtf) GetNextPlayerId() uint64 {
	if m != nil && m.NextPlayerId != nil {
		return *m.NextPlayerId
	}
	return 0
}

func (m *DDZPlayCardNtf) GetNextStage() *NextStage {
	if m != nil {
		return m.NextStage
	}
	return nil
}

// DDZGameOverNtf 斗地主游戏结束通知
// CMD: 0x1500C
type DDZGameOverNtf struct {
	WinnerId         *uint64              `protobuf:"varint,1,opt,name=winner_id,json=winnerId" json:"winner_id,omitempty"`
	ShowHandTime     *uint32              `protobuf:"varint,2,opt,name=show_hand_time,json=showHandTime" json:"show_hand_time,omitempty"`
	Bills            []*DDZBillPlayerInfo `protobuf:"bytes,3,rep,name=bills" json:"bills,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DDZGameOverNtf) Reset()                    { *m = DDZGameOverNtf{} }
func (m *DDZGameOverNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZGameOverNtf) ProtoMessage()               {}
func (*DDZGameOverNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *DDZGameOverNtf) GetWinnerId() uint64 {
	if m != nil && m.WinnerId != nil {
		return *m.WinnerId
	}
	return 0
}

func (m *DDZGameOverNtf) GetShowHandTime() uint32 {
	if m != nil && m.ShowHandTime != nil {
		return *m.ShowHandTime
	}
	return 0
}

func (m *DDZGameOverNtf) GetBills() []*DDZBillPlayerInfo {
	if m != nil {
		return m.Bills
	}
	return nil
}

// BillPlayersInfo 结算玩家账单
type DDZBillPlayerInfo struct {
	PlayerId         *uint64  `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	PlayerName       *string  `protobuf:"bytes,2,opt,name=player_name,json=playerName" json:"player_name,omitempty"`
	Base             *int32   `protobuf:"varint,3,opt,name=base" json:"base,omitempty"`
	Multiple         *int32   `protobuf:"varint,4,opt,name=multiple" json:"multiple,omitempty"`
	Score            *int64   `protobuf:"varint,5,opt,name=score" json:"score,omitempty"`
	CurrentScore     *int64   `protobuf:"varint,6,opt,name=current_score,json=currentScore" json:"current_score,omitempty"`
	Lord             *bool    `protobuf:"varint,7,opt,name=lord" json:"lord,omitempty"`
	OutCards         []uint32 `protobuf:"varint,8,rep,name=out_cards,json=outCards" json:"out_cards,omitempty"`
	HandCards        []uint32 `protobuf:"varint,9,rep,name=hand_cards,json=handCards" json:"hand_cards,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DDZBillPlayerInfo) Reset()                    { *m = DDZBillPlayerInfo{} }
func (m *DDZBillPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*DDZBillPlayerInfo) ProtoMessage()               {}
func (*DDZBillPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *DDZBillPlayerInfo) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *DDZBillPlayerInfo) GetBase() int32 {
	if m != nil && m.Base != nil {
		return *m.Base
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetMultiple() int32 {
	if m != nil && m.Multiple != nil {
		return *m.Multiple
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetCurrentScore() int64 {
	if m != nil && m.CurrentScore != nil {
		return *m.CurrentScore
	}
	return 0
}

func (m *DDZBillPlayerInfo) GetLord() bool {
	if m != nil && m.Lord != nil {
		return *m.Lord
	}
	return false
}

func (m *DDZBillPlayerInfo) GetOutCards() []uint32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func (m *DDZBillPlayerInfo) GetHandCards() []uint32 {
	if m != nil {
		return m.HandCards
	}
	return nil
}

// DDZTuoGuanReq 托管/取消托管请求
// CMD: 0x1500D
type DDZTuoGuanReq struct {
	Tuoguan          *bool  `protobuf:"varint,1,opt,name=tuoguan" json:"tuoguan,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DDZTuoGuanReq) Reset()                    { *m = DDZTuoGuanReq{} }
func (m *DDZTuoGuanReq) String() string            { return proto.CompactTextString(m) }
func (*DDZTuoGuanReq) ProtoMessage()               {}
func (*DDZTuoGuanReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *DDZTuoGuanReq) GetTuoguan() bool {
	if m != nil && m.Tuoguan != nil {
		return *m.Tuoguan
	}
	return false
}

// DDZTuoGuanNtf 托管状态通知
// CMD: 0x1500F
type DDZTuoGuanNtf struct {
	Tuoguan          *bool  `protobuf:"varint,1,opt,name=tuoguan" json:"tuoguan,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DDZTuoGuanNtf) Reset()                    { *m = DDZTuoGuanNtf{} }
func (m *DDZTuoGuanNtf) String() string            { return proto.CompactTextString(m) }
func (*DDZTuoGuanNtf) ProtoMessage()               {}
func (*DDZTuoGuanNtf) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *DDZTuoGuanNtf) GetTuoguan() bool {
	if m != nil && m.Tuoguan != nil {
		return *m.Tuoguan
	}
	return false
}

// TODO 恢复对局（暂定）
// DDZResumeGameReq 恢复对局请求
// CMD: 0x15010
type DDZResumeGameReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DDZResumeGameReq) Reset()                    { *m = DDZResumeGameReq{} }
func (m *DDZResumeGameReq) String() string            { return proto.CompactTextString(m) }
func (*DDZResumeGameReq) ProtoMessage()               {}
func (*DDZResumeGameReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

// DDZResumeGameRsp 恢复对局返回
// CMD: 0x15011
type DDZResumeGameRsp struct {
	Result           *Result      `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	GameInfo         *DDZDeskInfo `protobuf:"bytes,2,opt,name=game_info,json=gameInfo" json:"game_info,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DDZResumeGameRsp) Reset()                    { *m = DDZResumeGameRsp{} }
func (m *DDZResumeGameRsp) String() string            { return proto.CompactTextString(m) }
func (*DDZResumeGameRsp) ProtoMessage()               {}
func (*DDZResumeGameRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

func (m *DDZResumeGameRsp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DDZResumeGameRsp) GetGameInfo() *DDZDeskInfo {
	if m != nil {
		return m.GameInfo
	}
	return nil
}

// DDZDeskInfo 游戏基本信息
type DDZDeskInfo struct {
	Players          []*DDZPlayerInfo `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
	Stage            *NextStage       `protobuf:"bytes,2,opt,name=stage" json:"stage,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DDZDeskInfo) Reset()                    { *m = DDZDeskInfo{} }
func (m *DDZDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*DDZDeskInfo) ProtoMessage()               {}
func (*DDZDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

func (m *DDZDeskInfo) GetPlayers() []*DDZPlayerInfo {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *DDZDeskInfo) GetStage() *NextStage {
	if m != nil {
		return m.Stage
	}
	return nil
}

// DDZPlayerInfo 游戏中的玩家信息
type DDZPlayerInfo struct {
	PlayerInfo       *RoomPlayerInfo `protobuf:"bytes,1,opt,name=player_info,json=playerInfo" json:"player_info,omitempty"`
	OutCards         []uint32        `protobuf:"varint,2,rep,name=out_cards,json=outCards" json:"out_cards,omitempty"`
	HandCards        []uint32        `protobuf:"varint,3,rep,name=hand_cards,json=handCards" json:"hand_cards,omitempty"`
	Lord             *bool           `protobuf:"varint,4,opt,name=lord" json:"lord,omitempty"`
	Tuoguan          *bool           `protobuf:"varint,5,opt,name=tuoguan" json:"tuoguan,omitempty"`
	IsDouble         *bool           `protobuf:"varint,6,opt,name=is_double,json=isDouble" json:"is_double,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DDZPlayerInfo) Reset()                    { *m = DDZPlayerInfo{} }
func (m *DDZPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*DDZPlayerInfo) ProtoMessage()               {}
func (*DDZPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

func (m *DDZPlayerInfo) GetPlayerInfo() *RoomPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *DDZPlayerInfo) GetOutCards() []uint32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func (m *DDZPlayerInfo) GetHandCards() []uint32 {
	if m != nil {
		return m.HandCards
	}
	return nil
}

func (m *DDZPlayerInfo) GetLord() bool {
	if m != nil && m.Lord != nil {
		return *m.Lord
	}
	return false
}

func (m *DDZPlayerInfo) GetTuoguan() bool {
	if m != nil && m.Tuoguan != nil {
		return *m.Tuoguan
	}
	return false
}

func (m *DDZPlayerInfo) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return false
}

func init() {
	proto.RegisterType((*Result)(nil), "room.Result")
	proto.RegisterType((*NextStage)(nil), "room.NextStage")
	proto.RegisterType((*DDZStartGameNtf)(nil), "room.DDZStartGameNtf")
	proto.RegisterType((*DDZDealNtf)(nil), "room.DDZDealNtf")
	proto.RegisterType((*DDZGrabLordReq)(nil), "room.DDZGrabLordReq")
	proto.RegisterType((*DDZGrabLordNtf)(nil), "room.DDZGrabLordNtf")
	proto.RegisterType((*DDZLordNtf)(nil), "room.DDZLordNtf")
	proto.RegisterType((*DDZDoubleReq)(nil), "room.DDZDoubleReq")
	proto.RegisterType((*DDZDoubleNtf)(nil), "room.DDZDoubleNtf")
	proto.RegisterType((*DDZPlayCardReq)(nil), "room.DDZPlayCardReq")
	proto.RegisterType((*DDZPlayCardRsp)(nil), "room.DDZPlayCardRsp")
	proto.RegisterType((*DDZPlayCardNtf)(nil), "room.DDZPlayCardNtf")
	proto.RegisterType((*DDZGameOverNtf)(nil), "room.DDZGameOverNtf")
	proto.RegisterType((*DDZBillPlayerInfo)(nil), "room.DDZBillPlayerInfo")
	proto.RegisterType((*DDZTuoGuanReq)(nil), "room.DDZTuoGuanReq")
	proto.RegisterType((*DDZTuoGuanNtf)(nil), "room.DDZTuoGuanNtf")
	proto.RegisterType((*DDZResumeGameReq)(nil), "room.DDZResumeGameReq")
	proto.RegisterType((*DDZResumeGameRsp)(nil), "room.DDZResumeGameRsp")
	proto.RegisterType((*DDZDeskInfo)(nil), "room.DDZDeskInfo")
	proto.RegisterType((*DDZPlayerInfo)(nil), "room.DDZPlayerInfo")
	proto.RegisterEnum("room.DDZStage", DDZStage_name, DDZStage_value)
	proto.RegisterEnum("room.CardType", CardType_name, CardType_value)
}

func init() { proto.RegisterFile("game_ddz.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1028 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xc6, 0xf9, 0xb5, 0x2b, 0x3f, 0xd3, 0xdb, 0x0c, 0x22, 0xec, 0x6a, 0xc5, 0x60, 0x06, 0x69,
	0xd8, 0xd5, 0x8e, 0x44, 0x16, 0x38, 0x22, 0x25, 0x71, 0x94, 0x8d, 0x08, 0xc9, 0xa8, 0x9d, 0x45,
	0x22, 0x07, 0x2c, 0x27, 0xee, 0x99, 0xb1, 0xb0, 0xdd, 0x59, 0xdb, 0x61, 0x77, 0x38, 0x72, 0x46,
	0x5c, 0x39, 0xec, 0x33, 0xf1, 0x18, 0x1c, 0x78, 0x0b, 0xd4, 0x3f, 0xf6, 0xd8, 0xcb, 0xcc, 0x64,
	0xe7, 0xd6, 0x55, 0x5f, 0x75, 0xf9, 0xab, 0xaa, 0xaf, 0x3a, 0x81, 0xee, 0x85, 0x1b, 0x52, 0xc7,
	0xf3, 0x7e, 0x3b, 0xdd, 0xc6, 0x2c, 0x65, 0xb8, 0x16, 0x33, 0x16, 0x3e, 0x84, 0xb5, 0x9b, 0x50,
	0xe9, 0x31, 0xbf, 0x83, 0x06, 0xa1, 0xc9, 0x2e, 0x48, 0xf1, 0x27, 0xa0, 0xd3, 0x38, 0x76, 0x36,
	0xcc, 0xa3, 0x3d, 0xed, 0x48, 0x3b, 0xe9, 0x90, 0x26, 0x8d, 0xe3, 0x11, 0xf3, 0x68, 0x06, 0x79,
	0x34, 0xd9, 0xf4, 0x2a, 0x47, 0xda, 0x89, 0x21, 0x20, 0x8b, 0x26, 0x1b, 0x73, 0x0c, 0xc6, 0x9c,
	0xbe, 0x49, 0xed, 0xd4, 0xbd, 0xa0, 0xf8, 0x18, 0xea, 0x09, 0x3f, 0x88, 0xfb, 0xdd, 0x7e, 0xf7,
	0x94, 0x7f, 0xee, 0xd4, 0xb2, 0x56, 0x02, 0x26, 0x12, 0xc4, 0x18, 0x6a, 0xa9, 0x1f, 0x52, 0x91,
	0xa9, 0x43, 0xc4, 0xd9, 0xfc, 0x19, 0x0e, 0x64, 0x58, 0x9c, 0x4e, 0xdc, 0x90, 0xce, 0xd3, 0x73,
	0xfc, 0x08, 0x8c, 0x6d, 0xe0, 0x5e, 0xd1, 0xd8, 0xf1, 0x3d, 0x91, 0xb0, 0x46, 0x74, 0xe9, 0x98,
	0x7a, 0xf8, 0x14, 0x20, 0xa2, 0x6f, 0x52, 0x47, 0x7e, 0x8e, 0x67, 0x6a, 0xf5, 0x0f, 0xe4, 0xe7,
	0x72, 0x3a, 0xc4, 0x88, 0xb2, 0xa3, 0x49, 0x00, 0x2c, 0x6b, 0x65, 0x51, 0x37, 0xe0, 0xa9, 0x0f,
	0xa1, 0xbe, 0x71, 0x63, 0x2f, 0xe9, 0x69, 0x47, 0xd5, 0x93, 0x0e, 0x91, 0xc6, 0xbd, 0x73, 0x1e,
	0x43, 0xd7, 0xb2, 0x56, 0x93, 0xd8, 0x5d, 0xcf, 0x58, 0xec, 0x11, 0xfa, 0x8a, 0x57, 0x76, 0x11,
	0xbb, 0x6b, 0xc1, 0x56, 0x27, 0xe2, 0x6c, 0xfe, 0xa5, 0x95, 0xc2, 0xf6, 0x56, 0x96, 0xe5, 0xa8,
	0x5c, 0xe7, 0xc0, 0xc7, 0xd0, 0x15, 0xcc, 0xae, 0x6f, 0x55, 0xc5, 0xad, 0x36, 0xf7, 0x9e, 0xdd,
	0xdc, 0x93, 0xda, 0x5e, 0xfe, 0x7f, 0x6a, 0xa2, 0x29, 0xef, 0xc5, 0xea, 0x31, 0x40, 0xca, 0x52,
	0x37, 0x70, 0x72, 0x6e, 0x1d, 0x62, 0x08, 0x0f, 0x2f, 0x8c, 0x37, 0xd4, 0xf3, 0xb7, 0xae, 0xdf,
	0xab, 0xca, 0x86, 0x0a, 0xe3, 0xde, 0x84, 0x9e, 0x42, 0x9b, 0x0f, 0x89, 0xed, 0xd6, 0x01, 0xe5,
	0xed, 0x7c, 0x04, 0x86, 0x9f, 0x38, 0x9e, 0xb0, 0x55, 0x4f, 0x75, 0x3f, 0x91, 0xb8, 0xf9, 0x56,
	0x2b, 0x44, 0xef, 0xe5, 0x5f, 0x4a, 0x55, 0x29, 0xa7, 0xc2, 0x9f, 0x41, 0x5b, 0x16, 0xa7, 0xf0,
	0xaa, 0x28, 0xaf, 0x25, 0x7c, 0x2a, 0xe4, 0xbe, 0xa5, 0xd8, 0x62, 0xe8, 0x7c, 0x34, 0x23, 0x57,
	0x6a, 0xe3, 0x66, 0xcd, 0x3d, 0x05, 0x83, 0x1f, 0x9c, 0xf4, 0x6a, 0x2b, 0x79, 0xe5, 0x5b, 0xc3,
	0xef, 0x2d, 0xaf, 0xb6, 0x94, 0xe8, 0x1b, 0x75, 0x32, 0xbf, 0x2d, 0x27, 0x4d, 0xb6, 0xf8, 0x18,
	0x1a, 0xb1, 0xd8, 0x5e, 0x51, 0x70, 0xab, 0xdf, 0x96, 0x77, 0xe5, 0x46, 0x13, 0x85, 0x99, 0xff,
	0x68, 0xa5, 0x8b, 0x7b, 0x9b, 0x95, 0x53, 0xad, 0xdc, 0x4a, 0xb5, 0x7a, 0x37, 0xd5, 0x6b, 0xbd,
	0xac, 0x59, 0xb8, 0x16, 0xfd, 0xca, 0xf4, 0x32, 0x64, 0xe1, 0x4d, 0x82, 0xae, 0xef, 0x15, 0x74,
	0x63, 0x6f, 0xd3, 0x7f, 0x57, 0xab, 0xe6, 0x86, 0x74, 0xf1, 0x2b, 0x8d, 0x55, 0x9d, 0xaf, 0xfd,
	0x28, 0x2a, 0xd5, 0x29, 0x1d, 0x53, 0x8f, 0xb3, 0x48, 0x2e, 0xd9, 0x6b, 0xe7, 0xd2, 0x8d, 0x3c,
	0xa7, 0xf0, 0x24, 0xb5, 0xb9, 0xf7, 0x85, 0x1b, 0x79, 0x4b, 0x3f, 0xa4, 0xf8, 0x19, 0xd4, 0xd7,
	0x7e, 0x10, 0x24, 0x42, 0xdb, 0xad, 0xfe, 0xc7, 0xf9, 0xa3, 0x36, 0xf4, 0x83, 0x40, 0x71, 0x8d,
	0xce, 0x19, 0x91, 0x51, 0xe6, 0x1f, 0x15, 0x78, 0xf0, 0x3f, 0xf0, 0xee, 0x7e, 0x7f, 0x0a, 0x2d,
	0x05, 0x46, 0xae, 0x22, 0x61, 0x10, 0x90, 0xae, 0xb9, 0x1b, 0x8a, 0x17, 0x93, 0x3f, 0xd9, 0xa2,
	0xeb, 0x75, 0x22, 0xce, 0xf8, 0x21, 0xe8, 0xe1, 0x2e, 0x48, 0xfd, 0x6d, 0x20, 0xf5, 0x58, 0x27,
	0xb9, 0xcd, 0x07, 0x98, 0x6c, 0x58, 0x4c, 0x45, 0x57, 0xab, 0x44, 0x1a, 0xf8, 0x73, 0xe8, 0x6c,
	0x76, 0x71, 0x4c, 0xa3, 0xd4, 0x91, 0x68, 0x43, 0xa0, 0x6d, 0xe5, 0xb4, 0x45, 0x10, 0x86, 0x5a,
	0xc0, 0x62, 0xaf, 0xd7, 0x94, 0xcf, 0x0f, 0x3f, 0x73, 0xf2, 0x6c, 0x97, 0x3a, 0x52, 0x13, 0xba,
	0xd0, 0x84, 0xce, 0x76, 0xe9, 0x48, 0xc8, 0xe2, 0x31, 0x80, 0xe8, 0x9f, 0x44, 0x0d, 0x81, 0x1a,
	0xdc, 0x23, 0x60, 0xf3, 0x4b, 0xe8, 0x58, 0xd6, 0x6a, 0xb9, 0x63, 0x93, 0x9d, 0x1b, 0xf1, 0x3d,
	0xe8, 0x41, 0x33, 0xdd, 0xb1, 0x8b, 0x9d, 0x1b, 0xa9, 0x95, 0xce, 0xcc, 0x72, 0x28, 0x1f, 0xde,
	0xed, 0xa1, 0x18, 0x90, 0x65, 0xad, 0xb8, 0xcc, 0x43, 0xca, 0xc7, 0x4d, 0xe8, 0x2b, 0xf3, 0xf2,
	0x5d, 0xdf, 0xfb, 0xee, 0x07, 0x3e, 0x05, 0x43, 0xfc, 0x4e, 0xfa, 0xd1, 0x39, 0x53, 0xef, 0xfe,
	0x83, 0x7c, 0xca, 0x16, 0x4d, 0x7e, 0x11, 0xf3, 0xd5, 0x79, 0x0c, 0x3f, 0x99, 0x1b, 0x68, 0x15,
	0x00, 0xfc, 0x0c, 0x9a, 0x72, 0x56, 0x72, 0xb7, 0x5b, 0xfd, 0x0f, 0xf3, 0xcb, 0x05, 0x79, 0x64,
	0x31, 0xf8, 0x8b, 0xec, 0x47, 0xf2, 0x96, 0x5f, 0x18, 0x89, 0x9a, 0x7f, 0x6b, 0xa2, 0x1d, 0x05,
	0x0d, 0x7d, 0x93, 0xcb, 0x44, 0x10, 0x95, 0x15, 0x1d, 0xaa, 0x8a, 0x18, 0x0b, 0x0b, 0x1f, 0x53,
	0xe2, 0xc9, 0xa4, 0x77, 0x3d, 0xbd, 0xca, 0x9d, 0xd3, 0xab, 0xbe, 0x33, 0xbd, 0x5c, 0x0d, 0xb5,
	0x82, 0x1a, 0x0a, 0x53, 0xa9, 0x97, 0xa6, 0x52, 0x7e, 0x64, 0x1b, 0xe5, 0x47, 0xf6, 0xc9, 0x5b,
	0x0d, 0xf4, 0xec, 0x9f, 0x00, 0xc6, 0x62, 0x51, 0x1d, 0x7b, 0x39, 0x98, 0x8c, 0x9d, 0xf9, 0x62,
	0x3e, 0x46, 0x1f, 0x94, 0x7d, 0xd6, 0x78, 0x30, 0x43, 0x5a, 0xd9, 0x37, 0x1a, 0xcc, 0x66, 0xa8,
	0x52, 0xf6, 0x4d, 0xc8, 0x60, 0x88, 0xaa, 0xf8, 0x50, 0xcc, 0x3e, 0xbb, 0xbb, 0x78, 0x39, 0x9c,
	0x8d, 0x51, 0x0d, 0x7f, 0x24, 0x36, 0x51, 0x79, 0xcf, 0x66, 0x83, 0x9f, 0xa6, 0xf3, 0x09, 0xaa,
	0x97, 0x13, 0x2c, 0x7e, 0x1c, 0x13, 0xd4, 0x78, 0xf2, 0xaf, 0x06, 0x7a, 0xf6, 0x8c, 0xe1, 0x16,
	0x34, 0x47, 0xcb, 0x8c, 0x56, 0x07, 0x8c, 0xd1, 0xd2, 0xb1, 0xa7, 0xf3, 0xc9, 0x6c, 0x8c, 0x34,
	0x85, 0x9d, 0x0d, 0xa6, 0x04, 0x55, 0x32, 0xec, 0xc5, 0xcb, 0xf9, 0x6a, 0x8a, 0xaa, 0xb8, 0x0d,
	0xba, 0xc2, 0x6c, 0x54, 0x53, 0xe0, 0x92, 0x4c, 0xcf, 0x66, 0x63, 0x54, 0x57, 0xe0, 0xf3, 0xc1,
	0xdc, 0xfa, 0x0a, 0x35, 0x0a, 0x56, 0x1f, 0x35, 0x71, 0x17, 0x20, 0x0f, 0xb5, 0x91, 0xae, 0xec,
	0xe7, 0x36, 0x0f, 0xb6, 0x91, 0x51, 0xb4, 0xfb, 0x36, 0x02, 0x65, 0x7f, 0xad, 0xf0, 0x56, 0xd1,
	0xee, 0xdb, 0xa8, 0xad, 0x48, 0x0e, 0x17, 0x3f, 0x0c, 0x51, 0x07, 0x1f, 0x40, 0x6b, 0xb4, 0x74,
	0xbe, 0x9f, 0xce, 0x27, 0xc2, 0xd1, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x02, 0xd8, 0xb6,
	0x15, 0x0a, 0x00, 0x00,
}
