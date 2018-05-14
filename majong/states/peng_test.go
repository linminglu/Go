package states

import (
	"fmt"
	majongpb "steve/server_pb/majong"
	"testing"

	"steve/majong/interfaces"

	"github.com/Sirupsen/logrus"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// 测试碰状态后接受到出牌消息
func TestPengState_chupai(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	//初始化牌局信息
	flow := interfaces.NewMockMajongFlow(mc)
	mjContext := majongpb.MajongContext{
		Players:      []*majongpb.Player{},
		ActivePlayer: 1,
	}

	// 初始玩家信息
	mjContext.Players = mjContext.Players[0:0]
	for i := 0; i < 4; i++ {
		mjContext.Players = append(mjContext.Players, &majongpb.Player{
			PalyerId:  uint64(i),
			HandCards: []*majongpb.Card{&Card1W, &Card1T, &Card1B, &Card2W,&Card3W},
		})
	}
	// 序列化消息
	chupaiEvent := &majongpb.ChupaiRequestEvent{
		Head:  &majongpb.RequestEventHead{PlayerId: 1},
		Cards: &Card3W,
	}

	eventContext, err := proto.Marshal(chupaiEvent)
	if err != nil {
		fmt.Println(err)
	}
	flow.EXPECT().GetMajongContext().Return(&mjContext).AnyTimes()
	start := "碰状态"
	logrus.WithFields(logrus.Fields{
		"状态":           start,
		"pengPlayerID": mjContext.Players[1].PalyerId,
		"handCards":    mjContext.Players[1].GetHandCards(),
		"OutCards":     mjContext.Players[1].GetOutCards(),
	}).Info("前")
	// 碰状态接受到出牌消息
	p := new(PengState)
	newStateID, err := p.ProcessEvent(majongpb.EventID_event_chupai_request, eventContext, flow)
	if newStateID == majongpb.StateID_state_chupai {
		start = "出牌状态"
	}
	logrus.WithFields(logrus.Fields{
		"状态":           start,
		"pengPlayerID": mjContext.Players[1].PalyerId,
		"handCards":    mjContext.Players[1].GetHandCards(),
		"OutCards":     mjContext.Players[1].GetOutCards(),
	}).Info("后")
	assert.Nil(t, err)
	// 手牌中是否删除了1W
	assert.Equal(t, mjContext.Players[1].HandCards, []*majongpb.Card{&Card1W, &Card1T, &Card1B, &Card2W})
	// 玩家出牌数组中是否添加1W
	assert.Equal(t, mjContext.Players[1].GetOutCards()[len(mjContext.Players[1].GetOutCards())-1], &Card3W)
	// 返回是否是出牌状态ID
	assert.Equal(t, majongpb.StateID_state_chupai, newStateID)
}
