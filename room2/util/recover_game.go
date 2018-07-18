package util

import (
	"steve/client_pb/room"
	server_pb "steve/server_pb/majong"
	"time"
	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"steve/majong/utils"
	"steve/room2/desk"
	"steve/room2/desk/contexts"
)

func getStateCostTime(entryTime int64) (costTime uint32) {
	nowTime := time.Now().Unix()
	if nowTime > entryTime {
		costTime = uint32(nowTime - entryTime)
	}
	return
}

func getOperatePlayerID(mjContext *server_pb.MajongContext) *uint64 {
	state := mjContext.GetCurState()
	var playerID uint64
	switch state {
	case server_pb.StateID_state_chupai, server_pb.StateID_state_hu, server_pb.StateID_state_chupaiwenxun:
		playerID = mjContext.GetLastChupaiPlayer()
	case server_pb.StateID_state_angang, server_pb.StateID_state_gang, server_pb.StateID_state_waitqiangganghu:
		playerID = mjContext.GetLastGangPlayer()
	case server_pb.StateID_state_zimo, server_pb.StateID_state_zixun, server_pb.StateID_state_bugang:
		playerID = mjContext.GetLastMopaiPlayer()
	case server_pb.StateID_state_peng:
		playerID = mjContext.GetLastPengPlayer()
	case server_pb.StateID_state_mopai:
		playerID = mjContext.GetMopaiPlayer()
	case server_pb.StateID_state_qiangganghu:
		if players := mjContext.GetLastHuPlayers(); len(players) > 0 {
			playerID = players[0]
		}
	default: // 前端要求必须有一个有效值
		playerID = mjContext.GetLastMopaiPlayer()
	}
	return &playerID
}

func getGameStage(curState server_pb.StateID) (stage room.GameStage) {
	switch curState {
	case server_pb.StateID_state_huansanzhang:
		stage = room.GameStage_GAMESTAGE_HUANSANZHANG
	case server_pb.StateID_state_dingque:
		stage = room.GameStage_GAMESTAGE_DINGQUE
	default:
		stage = room.GameStage_GAMESTAGE_PLAYCARD
	}
	return
}

func getDoorCard(mjContext *server_pb.MajongContext) *uint32 {
	if mjContext.GetCurState() == server_pb.StateID_state_zixun {
		DoorCard := uint32(mjContext.GetLastMopaiCard().GetPoint())
		return &DoorCard
	}
	return nil
}

func getRecoverPlayerInfo(reqPlayerID uint64, d *desk.Desk) (recoverPlayerInfo []*room.GamePlayerInfo) {
	logEntry := logrus.WithFields(logrus.Fields{
		"func_name": "getRecoverPlayerInfo",
	})
	ctx :=d.GetConfig().Context.(contexts.MjContext).MjContext
	mjContext := &ctx
	deskPlayers := d.GetDeskPlayers()
	for _, deskPlayer := range deskPlayers {
		playerID := deskPlayer.GetPlayerID()
		roomPlayerInfo := translateToRoomPlayer(deskPlayer)
		player := utils.GetMajongPlayer(playerID, mjContext)
		if player == nil {
			logEntry.WithField("palyerID: ", playerID).Errorln("mjContext找不到对应玩家")
			continue
		}
		logEntry.Errorln("原生数据")
		logEntry.Errorln(player)
		svrHandCard := player.GetHandCards()
		handCardCount := uint32(len(svrHandCard))
		gamePlayerInfo := &room.GamePlayerInfo{
			PlayerInfo:    &roomPlayerInfo,
			Color:         ServerColor2ClientColor(player.DingqueColor).Enum(),
			HandCardCount: &handCardCount,
			IsTuoguan:     proto.Bool(deskPlayer.IsTuoguan()),
			IsTing:        proto.Bool(IsTing(player)),
			TingType:      getTingType(player),
		}

		xpState := room.XingPaiState(player.GetXpState())
		gamePlayerInfo.XpState = &xpState
		if (gamePlayerInfo.GetXpState() | room.XingPaiState_XP_STATE_HU) != 0 {
			if len(player.HuCards) != 0 {
				gamePlayerInfo.HuType = HuTypeSvr2Client(player.HuCards[0].GetType())
			}
		}
		gamePlayerInfo.TingCardInfos = TingCardInfoSvr2Client(player.GetTingCardInfo())
		// 手牌组，请求恢复对局玩家才发
		if playerID == reqPlayerID {
			cltHandCard := ServerCards2Numbers(svrHandCard)
			handCardGroup := &room.CardsGroup{
				Cards: cltHandCard,
				Type:  room.CardsGroupType_CGT_HAND.Enum(),
			}
			gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, handCardGroup)
		}
		// 吃牌组
		var chiCardGroups []*room.CardsGroup
		for _, chiCard := range player.GetChiCards() {
			srcPlayerID := chiCard.GetSrcPlayer()
			card := ServerCard2Number(chiCard.GetCard())
			cards := []uint32{card, card + 1, card + 2}
			chiCardGroup := &room.CardsGroup{
				Cards: cards,
				Type:  room.CardsGroupType_CGT_CHI.Enum(),
				Pid:   proto.Uint64(srcPlayerID),
			}
			chiCardGroups = append(chiCardGroups, chiCardGroup)
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, chiCardGroups...)
		// 碰牌组,每一次碰牌填1张还是三张
		var pengCardGroups []*room.CardsGroup
		for _, pengCard := range player.GetPengCards() {
			srcPlayerID := pengCard.GetSrcPlayer()
			cards := []uint32{ServerCard2Number(pengCard.GetCard())}
			pengCardGroup := &room.CardsGroup{
				Cards: append(cards, cards[0], cards[0]),
				Type:  room.CardsGroupType_CGT_PENG.Enum(),
				Pid:   &srcPlayerID,
			}
			pengCardGroups = append(pengCardGroups, pengCardGroup)
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, pengCardGroups...)
		// 杠牌组
		var gangCardGroups []*room.CardsGroup
		for _, gangCard := range player.GetGangCards() {
			groupType := GangTypeSvr2Client(gangCard.GetType())
			srcPlayerID := gangCard.GetSrcPlayer()
			cards := []uint32{ServerCard2Number(gangCard.GetCard())}
			gangCardGroup := &room.CardsGroup{
				Cards: append(cards, cards[0], cards[0], cards[0]),
				Type:  &groupType,
				Pid:   &srcPlayerID,
			}
			gangCardGroups = append(gangCardGroups, gangCardGroup)
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, gangCardGroups...)
		// 胡牌组
		var huCardGroups []*room.CardsGroup
		for _, huCard := range player.GetHuCards() {
			srcPlayerID := huCard.GetSrcPlayer()
			huCardGroup := &room.CardsGroup{
				Cards:  []uint32{ServerCard2Number(huCard.GetCard())},
				Type:   room.CardsGroupType_CGT_HU.Enum(),
				Pid:    &srcPlayerID,
				IsReal: proto.Bool(huCard.GetIsReal()),
			}
			huCardGroups = append(huCardGroups, huCardGroup)
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, huCardGroups...)
		// 花牌组
		huaGroup := &room.CardsGroup{
			Cards: ServerCards2Numbers(player.GetHuaCards()),
			Type:  room.CardsGroupType_CGT_HUA.Enum(),
			Pid:   proto.Uint64(player.GetPalyerId()),
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, huaGroup)
		// 出牌组
		outCardGroup := &room.CardsGroup{
			Cards: ServerCards2Numbers(player.GetOutCards()),
			Type:  room.CardsGroupType_CGT_OUT.Enum(),
			Pid:   &playerID,
		}
		gamePlayerInfo.CardsGroup = append(gamePlayerInfo.CardsGroup, outCardGroup)
		recoverPlayerInfo = append(recoverPlayerInfo, gamePlayerInfo)
	}
	return
}

func getTingType(player *server_pb.Player) *room.TingType {
	state := player.GetTingStateInfo()
	if state.GetIsTianting() {
		return room.TingType_TT_TIAN_TING.Enum()
	}
	return room.TingType_TT_NORMAL_TING.Enum()

}

func getZixunInfo(playerID uint64, mjContext *server_pb.MajongContext) (*bool, *room.RoomZixunNtf) {
	if mjContext.GetCurState() != server_pb.StateID_state_zixun {
		return proto.Bool(false), nil
	}

	if mjContext.GetLastMopaiPlayer() != playerID {
		return proto.Bool(false), nil
	}
	player := utils.GetMajongPlayer(playerID, mjContext)
	return proto.Bool(true), zixunTransform(player.GetZixunRecord())
}

func getWenxunInfo(playerID uint64, mjContext *server_pb.MajongContext) (*bool, *room.RoomChupaiWenxunNtf) {
	if mjContext.GetCurState() != server_pb.StateID_state_chupaiwenxun {
		return proto.Bool(false), nil
	}

	player := utils.GetMajongPlayer(playerID, mjContext)
	enableActions := player.GetPossibleActions()
	if len(enableActions) == 0 || player.GetHasSelected() {
		return proto.Bool(false), nil
	}

	outCard := ServerCard2Number(mjContext.GetLastOutCard())
	wenXunInfo := &room.RoomChupaiWenxunNtf{
		Card: &outCard,
	}
	for _, action := range enableActions {
		switch action {
		case server_pb.Action_action_peng:
			wenXunInfo.EnablePeng = proto.Bool(true)
		case server_pb.Action_action_gang:
			wenXunInfo.EnableMinggang = proto.Bool(true)
		case server_pb.Action_action_hu:
			wenXunInfo.EnableDianpao = proto.Bool(true)
		case server_pb.Action_action_qi:
			wenXunInfo.EnableQi = proto.Bool(true)
		case server_pb.Action_action_chi:
			wenXunInfo.ChiInfo = &room.RoomChiInfo{
				ChiCard: proto.Uint32(ServerCard2Number(mjContext.GetLastOutCard())),
				Cards:   player.GetEnbleChiCards(),
			}
		}
	}
	return proto.Bool(true), wenXunInfo
}

func getQghInfo(playerID uint64, mjContext *server_pb.MajongContext) (*bool, *room.RoomWaitQianggangHuNtf) {
	if mjContext.GetCurState() != server_pb.StateID_state_waitqiangganghu {
		return proto.Bool(false), nil
	}

	player := utils.GetMajongPlayer(playerID, mjContext)
	enableActions := player.GetPossibleActions()
	if len(enableActions) == 0 || player.GetHasSelected() {
		return proto.Bool(false), nil
	}

	outCard := ServerCard2Number(mjContext.GetLastOutCard())
	gangPlayerID := mjContext.GetLastGangPlayer()
	qghInfo := &room.RoomWaitQianggangHuNtf{
		Card:         &outCard,
		SelfCan:      proto.Bool(len(player.GetPossibleActions()) != 0),
		FromPlayerId: &gangPlayerID,
	}
	return proto.Bool(true), qghInfo
}

func zixunTransform(record *server_pb.ZiXunRecord) *room.RoomZixunNtf {
	zixunNtf := &room.RoomZixunNtf{}
	zixunNtf.EnableAngangCards = record.GetEnableAngangCards()
	zixunNtf.EnableBugangCards = record.GetEnableBugangCards()
	zixunNtf.EnableChupaiCards = record.GetEnableChupaiCards()
	zixunNtf.EnableQi = proto.Bool(record.GetEnableQi())
	zixunNtf.EnableZimo = proto.Bool(record.GetEnableZimo())
	zixunNtf.EnableTing = proto.Bool(record.GetEnableTing())
	TingType := TingTypeSvr2Client(record.GetTingType())
	if TingType != nil {
		zixunNtf.TingType = TingType
	}
	huType := HuTypeSvr2Client(record.GetHuType())
	if huType != nil {
		zixunNtf.HuType = huType
	}
	zixunNtf.CanTingCardInfo = CanTingCardInfoSvr2Client(record.GetCanTingCardInfo())

	return zixunNtf
}
