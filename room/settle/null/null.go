package null

import (
	"steve/room/interfaces"
	server_pb "steve/server_pb/majong"
)

type nullSettler struct{}

func (s *nullSettler) Settle(desk interfaces.Desk, mjContext server_pb.MajongContext) {
}

func (s *nullSettler) RoundSettle(desk interfaces.Desk, mjContext server_pb.MajongContext) {
}

// NewNullSettle 初始化麻将结算
func NewNullSettle() *nullSettler {
	return &nullSettler{}
}