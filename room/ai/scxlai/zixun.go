package scxlai

import (
	"fmt"
	"steve/gutils"
	"steve/room/interfaces"
	"steve/server_pb/majong"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type zixunStateAI struct {
	maxDingqueTime time.Duration // 最大定缺时间
}

// // 注册 AI
// func init() {
// 	g := global.GetDeskAutoEventGenerator()
// 	g.RegisterAI(gGameID, majong.StateID_state_zixun, &zixunStateAI{})
// }

// GenerateAIEvent 生成 AI 事件
// 首先判断玩家是否时当前可以操作的玩家
// 是的话,判断当前玩家是否可以执行自动事件
// 可以的话,根据玩家状态生成不同的自动事件
// 1,玩家是碰自询:
//	 			之前胡过,自动事件:出最右的一张牌
//				之前没有胡过,自动事件:出最右的一张牌
// 2,玩家是摸牌自询:
//	 			之前胡过,自动事件:
//								可胡,等待三秒,然后自动胡牌
//								不可胡,无需等待,直接出牌
//				之前没有胡过,自动事件:
// 								1,出摸到的那张牌
//								2,如果是庄家首次出牌,出最右侧的牌
func (h *zixunStateAI) GenerateAIEvent(params interfaces.AIEventGenerateParams) (result interfaces.AIEventGenerateResult, err error) {
	result, err = interfaces.AIEventGenerateResult{
		Events: []interfaces.AIEvent{},
	}, nil
	var aiEvent interfaces.AIEvent
	mjContext := params.MajongContext
	player := gutils.GetMajongPlayer(params.PlayerID, mjContext)
	handCards := player.GetHandCards()
	if h.getZixunPlayer(mjContext) != params.PlayerID {
		return result, fmt.Errorf("当前玩家不允许进行自动操作")
	}
	if len(handCards) < 2 {
		return result, fmt.Errorf("手牌数量少于2")
	}
	switch mjContext.GetZixunType() {
	case majong.ZixunType_ZXT_PENG:
		{
			aiEvent = h.chupai(player, handCards[len(handCards)-1])
		}
	case majong.ZixunType_ZXT_NORMAL:
		{
			zxRecord := player.GetZixunRecord()
			canHu := zxRecord.GetEnableZimo()
			if len(player.GetHuCards()) > 0 && canHu && !gutils.CheckHasDingQueCard(handCards, player.GetDingqueColor()) {
				aiEvent = h.hu(player)
			} else {
				if player.GetMopaiCount() == 0 {
					aiEvent = h.chupai(player, handCards[len(handCards)-1])
				} else {
					aiEvent = h.chupai(player, mjContext.GetLastMopaiCard())
				}
			}
		}
	default:
		return
	}
	result.Events = append(result.Events, aiEvent)
	return
}

func (h *zixunStateAI) chupai(player *majong.Player, card *majong.Card) interfaces.AIEvent {
	eventContext := majong.ChupaiRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPalyerId(),
		},
		Cards: card,
	}
	data, err := proto.Marshal(&eventContext)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"func_name": "zixunStateAI.chupai",
			"player_id": player.GetPalyerId(),
		}).Errorln("事件序列化失败")
	}
	return interfaces.AIEvent{
		ID:      majong.EventID_event_chupai_request,
		Context: data,
	}
}

func (h *zixunStateAI) hu(player *majong.Player) interfaces.AIEvent {
	eventContext := majong.HuRequestEvent{
		Head: &majong.RequestEventHead{
			PlayerId: player.GetPalyerId(),
		},
	}
	data, err := proto.Marshal(&eventContext)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"func_name": "zixunStateAI.hu",
			"player_id": player.GetPalyerId(),
		}).Errorln("事件序列化失败")
	}
	return interfaces.AIEvent{
		ID:      majong.EventID_event_hu_request,
		Context: data,
	}
}

func (h *zixunStateAI) getZixunPlayer(mjContext *majong.MajongContext) uint64 {
	zxType := mjContext.GetZixunType()
	if zxType == majong.ZixunType_ZXT_PENG {
		return mjContext.GetLastPengPlayer()
	}
	return mjContext.GetLastMopaiPlayer()
}