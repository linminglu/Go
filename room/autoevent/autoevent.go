package autoevent

import (
	"steve/room/interfaces"
	"steve/room/interfaces/global"
	"steve/server_pb/majong"
	"time"
)

type autoEventGenerator struct {
	majongAIs map[int](map[majong.StateID]interfaces.MajongAI)
}

func init() {
	global.SetDeskAutoEventGenerator(&autoEventGenerator{
		majongAIs: map[int](map[majong.StateID]interfaces.MajongAI){},
	})
}

// getAI 根据状态和游戏 ID 获取 AI 对象
func (aeg *autoEventGenerator) getAI(mjContext *majong.MajongContext) interfaces.MajongAI {
	gameID := mjContext.GetGameId()
	gameAIs, ok := aeg.majongAIs[int(gameID)]
	if !ok {
		return nil
	}
	state := mjContext.GetCurState()
	AI, ok := gameAIs[state]
	if !ok {
		return nil
	}
	return AI
}

// getStateDuration 获取状态超时时间， 暂时返回 10 秒
func (aeg *autoEventGenerator) getStateDuration() time.Duration {
	return time.Second * 10 // TODO
}

// addAIEvents 将 AI 产生的事件添加到结果中
func (aeg *autoEventGenerator) addAIEvents(result *interfaces.AutoEventGenerateResult, aiResult *interfaces.AIEventGenerateResult, playerID uint64, eventType interfaces.EventType) {
	for _, aiEvent := range aiResult.Events {
		result.Events = append(result.Events, interfaces.Event{
			ID:        aiEvent.ID,
			Context:   aiEvent.Context,
			PlayerID:  playerID,
			EventType: eventType,
		})
	}
}

// handlePlayerAI 处理玩家 AI
func (aeg *autoEventGenerator) handlePlayerAI(result *interfaces.AutoEventGenerateResult, AI interfaces.MajongAI,
	player *majong.Player, mjContext *majong.MajongContext, aiType interfaces.AIType, robotLv int) {
	playerID := player.GetPalyerId()
	aiResult, err := AI.GenerateAIEvent(interfaces.AIEventGenerateParams{
		MajongContext: mjContext,
		PlayerID:      playerID,
		AIType:        aiType,
		RobotLv:       robotLv,
	})
	if err == nil {
		aeg.addAIEvents(result, &aiResult, playerID, interfaces.OverTimeEvent)
	}
}

// handlePlayerTuoGuan 处理玩家托管
func (aeg *autoEventGenerator) handlePlayerTuoGuan(result *interfaces.AutoEventGenerateResult, AI interfaces.MajongAI,
	player *majong.Player, mjContext *majong.MajongContext) {
	aeg.handlePlayerAI(result, AI, player, mjContext, interfaces.TuoGuangAI, 0)
}

// handlePlayerOverTime 处理玩家超时
func (aeg *autoEventGenerator) handlePlayerOverTime(result *interfaces.AutoEventGenerateResult, AI interfaces.MajongAI,
	player *majong.Player, mjContext *majong.MajongContext) {
	aeg.handlePlayerAI(result, AI, player, mjContext, interfaces.OverTimeAI, 0)
}

// handleOverTime 处理超时
func (aeg *autoEventGenerator) handleOverTime(AI interfaces.MajongAI, curTime time.Time, stateTime time.Time, mjContext *majong.MajongContext) (
	finish bool, result interfaces.AutoEventGenerateResult) {

	finish, result = false, interfaces.AutoEventGenerateResult{
		Events: []interfaces.Event{},
	}
	duration := aeg.getStateDuration()
	if duration == 0 || curTime.Sub(stateTime) < duration {
		return
	}
	players := mjContext.GetPlayers()
	for _, player := range players {
		aeg.handlePlayerOverTime(&result, AI, player, mjContext)
	}
	finish = true
	return
}

// isTuoGuan 玩家是否托管
func (aeg *autoEventGenerator) isTuoGuan(playerID uint64, tuoGuanPlayers []uint64) bool {
	for _, player := range tuoGuanPlayers {
		if playerID == player {
			return true
		}
	}
	return false
}

// handleTuoGuan 执行所有玩家的托管
func (aeg *autoEventGenerator) handleTuoGuan(tuoGuanPlayers []uint64, AI interfaces.MajongAI, curTime time.Time, stateTime time.Time, mjContext *majong.MajongContext) interfaces.AutoEventGenerateResult {
	result := interfaces.AutoEventGenerateResult{
		Events: []interfaces.Event{},
	}
	tuoguanOprTime := 1 * time.Second
	if curTime.Sub(stateTime) < tuoguanOprTime {
		return result
	}
	players := mjContext.GetPlayers()
	for _, player := range players {
		playerID := player.GetPalyerId()
		if aeg.isTuoGuan(playerID, tuoGuanPlayers) {
			aeg.handlePlayerTuoGuan(&result, AI, player, mjContext)
		}
	}
	return result
}

// GenerateV2 利用 AI 生成自动事件
func (aeg *autoEventGenerator) GenerateV2(params *interfaces.AutoEventGenerateParams) (result interfaces.AutoEventGenerateResult) {
	mjContext := params.MajongContext
	AI := aeg.getAI(mjContext)
	if AI == nil {
		return
	}
	if overTime, result := aeg.handleOverTime(AI, params.CurTime, params.StateTime, mjContext); overTime {
		return result
	}
	result = aeg.handleTuoGuan(params.TuoGuanPlayers, AI, params.CurTime, params.StateTime, mjContext)

	players := mjContext.GetPlayers()
	for _, player := range players {
		playerID := player.GetPalyerId()
		if lv, exist := params.RobotLv[playerID]; exist {
			aeg.handlePlayerAI(&result, AI, player, mjContext, interfaces.RobotAI, lv)
		}
	}
	return result
}

func (aeg *autoEventGenerator) RegisterAI(gameID int, stateID majong.StateID, AI interfaces.MajongAI) {
	if _, exist := aeg.majongAIs[gameID]; !exist {
		aeg.majongAIs[gameID] = make(map[majong.StateID]interfaces.MajongAI)
	}
	aeg.majongAIs[gameID][stateID] = AI
}