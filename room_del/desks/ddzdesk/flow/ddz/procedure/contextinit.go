package procedure

import (
	"github.com/Sirupsen/logrus"
	"steve/client_pb/room"
	"steve/entity/poker"
	"steve/entity/poker/ddz"
)

// CreateInitDDZContext 创建初始斗地主现场
func CreateInitDDZContext(players []uint64) *ddz.DDZContext {
	return &ddz.DDZContext{
		GameId:            int32(room.GameId_GAMEID_DOUDIZHU),
		CurState:          ddz.StateID_state_init,
		Players:           createDDZPlayers(players),
		FirstGrabPlayerId: 0,
		GrabbedCount:      0,
		AllAbandonCount:   0,
		TotalGrab:         0,
		TotalDouble:       1,
		CurCardType:       poker.CardType_CT_NONE,
		PassCount:         0,
		TotalBomb:         1,
		Spring:            true,
		AntiSpring:        true,
	}
}

// 根据玩家的playerID创建出斗地主Player
func createDDZPlayers(players []uint64) []*ddz.Player {
	logrus.WithField("players", players).Debug("创建斗地主玩家")
	result := make([]*ddz.Player, 0, len(players))
	for _, playerID := range players {
		result = append(result, &ddz.Player{
			PlayerId: playerID,
		})
	}
	return result
}