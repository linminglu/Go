package core

/*
	功能： 服务配置中心，定义RPC服务关联，Client消息分派。

 */
import (
	"steve/server_pb/gold"
	"steve/gold/server"
	"steve/client_pb/msgid"
)


/////////////////////////////////////////[1.配置线程模型]////////////////////////////////////
// 是否采用单线程运行所有协程(goroutime)
var bSingleThread = true

/////////////////////////////[2.定义RPC服务实现]//////////////////////////////////////////
// PB文件中定义的RPC服务接口,如果不提供RPC服务，设置为nil
var pbService interface{} = gold.RegisterGoldServer
// PB定义的RPC服务的实现类,如果不提供RPC服务，设置为nil
var pbServerImp interface{} = &server.GoldServer{}

/////////////////////////////[3.处理Client消息]////////////////////////////////////////////
// 如果需要处理Client消息，需要在GateWay配置消息ID开始~ 结束区间 关联到当前服务名,GateWay才会把消息转发到此服务.

// 添加从GateWay过来的Client消息处理
//msgid.MsgID_HALL_GET_PLAYER_INFO_REQ:msg.ProcessMatchReq,
var mapMsg  = map[msgid.MsgID] interface{} {
	//msgid.MsgID_HALL_GET_PLAYER_INFO_REQ:msg.ProcessMatchReq,
}

/////////////////////////////[4.通过nsq发布和订阅消息]////////////////////////////////////////////
// 4.1发布消息
/*
exposer := structs.GetGlobalExposer()
if err := exposer.Publisher.Publish("player_login", messageData); err != nil {
entry.WithError(err).Errorln("发布消息失败")
}
*/

// 4.2订阅消息
/*
	exposer := structs.GetGlobalExposer()
	if err := exposer.Subscriber.Subscribe("player_login", "match", &playerLoginHandler{}); err != nil {
		logrus.WithError(err).Panicln("订阅登录消息失败")
	}
 */

/////////////////////////////[5.服务初始化配置]////////////////////////////////////////////
// 比如从DB或文件加载配置
func InitServer() error {

	return nil
}




