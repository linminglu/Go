package scxlai

import (
	"steve/client_pb/room"
	"steve/gutils"
	"steve/room2/ai"
	"steve/server_pb/majong"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type fapaiStateAI struct {
}

// GenerateAIEvent 生成 发牌AI 事件
// 无论是超时、托管还是机器人，发牌发牌动画完成产生相应的事件
func (f *fapaiStateAI) GenerateAIEvent(params ai.AIEventGenerateParams) (result ai.AIEventGenerateResult, err error) {
	result, err = ai.AIEventGenerateResult{
		Events: []ai.AIEvent{},
	}, nil
	if params.AIType != ai.RobotAI { //不是机器人不能发送动画完成请求
		return
	}
	mjContext := params.MajongContext
	crPlayerIDs := mjContext.GetTempData().GetCartoonReqPlayerIDs()
	if len(crPlayerIDs) == len(mjContext.GetPlayers()) { //所有玩家都发送过动画完成请求
		return
	}
	player := gutils.GetMajongPlayer(params.PlayerID, mjContext)
	for _, playerID := range crPlayerIDs {
		if playerID == player.GetPalyerId() { // 当前玩家已经发送过
			return
		}
	}
	// 发送动画完成请求
	if event := CartoonFinsh(player, int32(room.CartoonType_CTNT_FAPAI)); event != nil {
		result.Events = append(result.Events, *event)
	}
	return
}

//CartoonFinsh 动画完成请求事件
func CartoonFinsh(player *majong.Player, cartoonType int32) *ai.AIEvent {
	event := majong.CartoonFinishRequestEvent{
		CartoonType: cartoonType,
		PlayerId:    player.GetPalyerId(),
	}
	data, err := proto.Marshal(&event)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"func_name":   "ai.CartoonFinsh",
			"player_id":   player.GetPalyerId(),
			"cartoonType": cartoonType,
		}).Errorln("事件序列化失败")
		return nil
	}
	logrus.WithFields(logrus.Fields{"func_name": "ai.CartoonFinsh", "player_id": player.GetPalyerId(), "cartoonType": cartoonType}).Errorln("机器人发送动画完成请求事件")
	return &ai.AIEvent{
		ID:      int32(majong.EventID_event_cartoon_finish_request),
		Context: data,
	}
}
