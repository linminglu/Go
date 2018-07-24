package utils

import (
	"steve/gutils"
	"steve/majong/interfaces"
	majongpb "steve/entity/majong"

	"github.com/Sirupsen/logrus"
)

//GetNextXpPlayerByID 获取下一个行牌玩家
func GetNextXpPlayerByID(srcPlayerID uint64, players []*majongpb.Player, mjContext *majongpb.MajongContext) (nextPalyer *majongpb.Player) {
	curPlayerID, i := srcPlayerID, 0
	for i < len(mjContext.Players) {
		nextPalyer = GetNextPlayerByID(players, curPlayerID)
		// 当前下个玩家可以继续，退出循环
		if gutils.IsPlayerContinue(nextPalyer.GetXpState(), mjContext) {
			break
		}
		curPlayerID = nextPalyer.GetPalyerId()
	}
	logrus.WithFields(logrus.Fields{"playerID": nextPalyer.GetPalyerId(),
		"playerStatus": nextPalyer.GetXpState()}).Info("获取下个正常状态的玩家")
	return nextPalyer
}

//GetCanXpPlayers 获取能行牌玩家数组
func GetCanXpPlayers(players []*majongpb.Player, mjContext *majongpb.MajongContext) []*majongpb.Player {
	newPlalyers := make([]*majongpb.Player, 0)
	for _, player := range players {
		// 不是正常行牌的玩家，不能检查胡，碰，杠，摸牌。。。
		if !gutils.IsPlayerContinue(player.GetXpState(), mjContext) {
			logrus.WithFields(logrus.Fields{"PlayerIDs": player.GetPalyerId(), "PlayerState": player.GetXpState()}).Info("不正常玩家")
			continue
		}
		newPlalyers = append(newPlalyers, player)
	}
	return newPlalyers
}

//IsAreThereEnoughpeople 判断是否有足够多的人数
func IsAreThereEnoughpeople(players []*majongpb.Player, mjContext *majongpb.MajongContext) bool {
	count := 0
	for _, player := range players {
		if gutils.IsPlayerContinue(player.GetXpState(), mjContext) {
			count++
		}
	}
	logrus.WithFields(logrus.Fields{"NormalPlayerConut": count}).Infoln("正常状态玩家数量")
	return count <= 1
}

//SettleOver 结算完成
func SettleOver(flow interfaces.MajongFlow, message *majongpb.SettleFinishEvent) {
	logEntry := logrus.WithFields(
		logrus.Fields{
			"func_name": "settleOver",
		},
	)
	mjContext := flow.GetMajongContext()
	playerIds := message.GetPlayerId()
	for _, pid := range playerIds {
		player := GetMajongPlayer(pid, mjContext)
		if player == nil {
			logEntry.WithField("playerID: ", pid).Errorln("找不到玩家")
			continue
		}
		player.XpState = player.GetXpState() | majongpb.XingPaiState_give_up
	}

}

// IsGameOverReturnState 判断游戏是否结束返回状态
func IsGameOverReturnState(mjContext *majongpb.MajongContext) majongpb.StateID {
	// 正常玩家<=1,游戏结束
	if IsAreThereEnoughpeople(mjContext.GetPlayers(), mjContext) {
		return majongpb.StateID_state_gameover
	}
	return majongpb.StateID_state_mopai
}

// GetRealHuCardPlayer 获取亮实胡牌的玩家
func GetRealHuCardPlayer(huPlayers []uint64, lostPlayer uint64, mjContext *majongpb.MajongContext) (isRealPlayerID uint64) {
	nextPlayerID := lostPlayer
	for i := 0; i < len(mjContext.GetPlayers()); i++ {
		p := GetNextXpPlayerByID(nextPlayerID, mjContext.GetPlayers(), mjContext)
		if Contains(huPlayers, p.GetPalyerId()) {
			isRealPlayerID = p.GetPalyerId()
			break
		}
		nextPlayerID = p.GetPalyerId()
	}
	logrus.WithFields(logrus.Fields{
		"RealHucardPlayer": isRealPlayerID,
	}).Infoln("亮实牌的玩家")
	return
}

// Contains 目标ID是否包含在数组内
func Contains(IDs []uint64, srcID uint64) bool {
	for _, ID := range IDs {
		if ID == srcID {
			return true
		}
	}
	return false
}
