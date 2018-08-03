package cache

import "fmt"

// HallPlayer 大厅玩家
type HallPlayer struct {
	PlayerID uint64 `protobuf:"varint,1,opt,name=playerID" json:"playerID,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	Gender   uint64 `protobuf:"bytes,4,opt,name=gender" json:"gender,omitempty"`
	Name     string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Phone    string `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
	IDdCard  string `protobuf:"bytes,7,opt,name=idCard" json:"idCard,omitempty"`
	Coin     uint64 `protobuf:"varint,8,opt,name=coin" json:"coin,omitempty"`
	GameID   uint64 `protobuf:"varint,9,opt,name=gameID" json:"gameID,omitempty"`
	State    uint64 `protobuf:"varint,10,opt,name=state" json:"state,omitempty"`
}

// RobotPlayer 机器人玩家
type RobotPlayer struct {
	PlayerID      uint64            `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	NickName      string            `protobuf:"bytes,2,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Avatar        string            `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	Coin          uint64            `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	State         uint64            `protobuf:"varint,5,opt,name=state" json:"state,omitempty"`
	GameIDWinRate map[uint64]uint64 `protobuf:"bytes,6,rep,name=game_id_win_rate,json=gameIdWinRate" json:"game_id_win_rate,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

// key formats
const (
	// AccountPlayerKey 账号关联的玩家
	AccountPlayerKey = "account:player:%v"

	// playerTokenKeyFmt
	playerTokenKeyFmt = "playertoken:%d"
)

// Player redis字段
const (
	// NickNameField ...昵称
	NickNameField = "nickname"
	// AvatarField ...头像
	AvatarField = "avatar"
	// GenderField  ...性别
	GenderField = "gender"
	// NameField  ...姓名
	NameField = "name"
	// PhoneField  ...联系电话
	PhoneField = "phone"
	// PlayerStateField ...玩家状态
	PlayerStateField = "game_state"
	// GameIDField ...正在进行的游戏id
	GameIDField = "game_id"
	// GateAddrField ...网关服地址
	GateAddrField = "gate_addr"
	// MatchAddrField ...匹配服地址
	MatchAddrField = "match_addr"
	// RoomAddrField ...房间服地址
	RoomAddrField = "room_addr"
)

// FmtAccountPlayerKey 账号所关联玩家 key
func FmtAccountPlayerKey(accountID uint64) string {
	return fmt.Sprintf(AccountPlayerKey, accountID)
}

// FmtPlayerIDKey 玩家ID key
func FmtPlayerIDKey(playerID uint64) string {
	return fmt.Sprintf("player:%v", playerID)
}

// FmtPlayerTokenKey format player's token key
func FmtPlayerTokenKey(playerID uint64) string {
	return fmt.Sprintf(playerTokenKeyFmt, playerID)
}
