package scxz

import (
	"steve/majong/global"
	"steve/majong/interfaces"
	"steve/majong/interfaces/facade"
	"steve/majong/utils"
	majongpb "steve/server_pb/majong"

	"github.com/golang/protobuf/proto"
)

// QiangGangHuSettleState 枪杠胡结算状态
type QiangGangHuSettleState struct {
}

var _ interfaces.MajongState = new(GangSettleState)

// ProcessEvent 处理事件
// 枪杠胡逻辑执行完后，进入枪杠胡结算状态
// 1.处理结算完成事件，返回摸牌状态
// 2.处理玩家认输事件，返回游戏结束状态
func (s *QiangGangHuSettleState) ProcessEvent(eventID majongpb.EventID, eventContext []byte, flow interfaces.MajongFlow) (newState majongpb.StateID, err error) {
	s.setMopaiPlayer(flow)
	if eventID == majongpb.EventID_event_settle_finish {
		message := &majongpb.SettleFinishEvent{}
		err := proto.Unmarshal(eventContext, message)
		if err != nil {
			return majongpb.StateID_state_gang_settle, global.ErrInvalidEvent
		}
		return s.settleOver(flow, message)
	}
	return majongpb.StateID(majongpb.StateID_state_gang_settle), global.ErrInvalidEvent
}

// OnEntry 进入状态
func (s *QiangGangHuSettleState) OnEntry(flow interfaces.MajongFlow) {
	s.doQiangGangHuSettle(flow)
}

// OnExit 退出状态
func (s *QiangGangHuSettleState) OnExit(flow interfaces.MajongFlow) {
}

// setMopaiPlayer 设置摸牌玩家
func (s *QiangGangHuSettleState) setMopaiPlayer(flow interfaces.MajongFlow) {
	mjContext := flow.GetMajongContext()
	mjContext.MopaiPlayer = mjContext.GetLastGangPlayer()
	mjContext.MopaiType = majongpb.MopaiType_MT_GANG
}

// doQiangGangHuSettle 抢杠胡结算
func (s *QiangGangHuSettleState) doQiangGangHuSettle(flow interfaces.MajongFlow) {
	mjContext := flow.GetMajongContext()

	allPlayers := make([]uint64, 0)
	for _, player := range mjContext.Players {
		if player.State == majongpb.PlayerState_normal {
			allPlayers = append(allPlayers, player.GetPalyerId())
		}
	}

	cardValues := make(map[uint64]uint32, 0)
	cardTypes := make(map[uint64][]majongpb.CardType, 0)
	genCount := make(map[uint64]uint32, 0)
	gameID := int(mjContext.GetGameId())

	huPlayers := mjContext.GetLastHuPlayers()
	for _, huPlayerID := range huPlayers {
		huPlayer := utils.GetPlayerByID(mjContext.Players, huPlayerID)
		cardParams := interfaces.CardCalcParams{
			HandCard: huPlayer.HandCards,
			PengCard: utils.TransPengCard(huPlayer.PengCards),
			GangCard: utils.TransGangCard(huPlayer.GangCards),
			HuCard:   mjContext.GetGangCard(),
			GameID:   gameID,
		}
		calculator := global.GetCardTypeCalculator()
		cardType, gen := calculator.Calculate(cardParams)
		cardValue, _ := calculator.CardTypeValue(gameID, cardType, gen)

		cardTypes[huPlayerID] = cardType
		cardValues[huPlayerID] = cardValue
		genCount[huPlayerID] = gen
	}

	params := interfaces.HuSettleParams{
		HuPlayers:  huPlayers,
		SrcPlayer:  mjContext.GetLastGangPlayer(),
		AllPlayers: allPlayers,
		SettleType: majongpb.SettleType_settle_dianpao,
		HuType:     majongpb.HuType_hu_qiangganghu,
		CardTypes:  cardTypes,
		CardValues: cardValues,
		GenCount:   genCount,
		SettleID:   mjContext.CurrentSettleId,
	}
	settleInfos := facade.SettleHu(global.GetGameSettlerFactory(), int(mjContext.GetGameId()), params)
	maxSID := uint64(0)
	for _, settleInfo := range settleInfos {
		mjContext.SettleInfos = append(mjContext.SettleInfos, settleInfo)
		if settleInfo.Id > maxSID {
			maxSID = settleInfo.Id
		}
	}
	mjContext.CurrentSettleId = maxSID
}

//settleOver 结算完成
func (s *QiangGangHuSettleState) settleOver(flow interfaces.MajongFlow, message *majongpb.SettleFinishEvent) (majongpb.StateID, error) {
	mjContext := flow.GetMajongContext()
	playerIds := message.GetPlayerId()
	if len(playerIds) != 0 {
		for _, pid := range playerIds {
			player := utils.GetMajongPlayer(pid, mjContext)
			if player == nil {
				return majongpb.StateID_state_gang_settle, global.ErrInvalidEvent
			}
			player.State = majongpb.PlayerState_give_up
		}
		return majongpb.StateID_state_gameover, nil
	}
	return majongpb.StateID(majongpb.StateID_state_mopai), nil
}
