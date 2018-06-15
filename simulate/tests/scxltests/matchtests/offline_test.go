package matchtests

import (
	"steve/simulate/config"
	"steve/simulate/connect"
	"steve/simulate/global"
	"steve/simulate/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test_OfflineMatch 测试离线时不会被匹配到
// 1. 登录玩家1，发送加入房间请求
// 2. 玩家1断开连接
// 3. 登录4个玩家，分别发送加入房间请求
// 预期：
//  后4个玩家都收到了创建房间通知和游戏开始通知
func Test_OfflineMatch(t *testing.T) {
	client1 := connect.NewTestClient(config.ServerAddr, config.ClientVersion)
	assert.NotNil(t, client1)
	player1, err := utils.LoginUser(client1, global.AllocUserName())
	utils.ApplyJoinDesk(player1)
	assert.Nil(t, err)
	client1.Stop()
	time.Sleep(time.Millisecond * 200) // 等200毫秒，确保连接断开

	params := global.NewCommonStartGameParams()
	deskData, err := utils.StartGame(params)
	assert.NotNil(t, deskData)
	assert.Nil(t, err)
}