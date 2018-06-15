package interfaces

import (
	"steve/server_pb/majong"
	"time"
)

// AutoEventGenerateParams 生成自动事件的参数
type AutoEventGenerateParams struct {
	MajongContext  *majong.MajongContext
	CurTime        time.Time
	StateTime      time.Time
	RobotLv        map[uint64]int
	TuoGuanPlayers []uint64
}

// AutoEventGenerateResult 自动事件生成结果
type AutoEventGenerateResult struct {
	Events []Event
}

// DeskAutoEventGenerator 牌桌自动事件产生器
type DeskAutoEventGenerator interface {
	GenerateV2(params *AutoEventGenerateParams) AutoEventGenerateResult
	RegisterAI(gameID int, stateID majong.StateID, AI MajongAI) // 注册 AI
}