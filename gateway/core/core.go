package core

import (
	"fmt"
	"steve/structs"
	"steve/structs/net"
	"steve/structs/service"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type gatewayCore struct {
	e *structs.Exposer

	dog net.WatchDog
}

// NewService 创建服务
func NewService() service.Service {
	return new(gatewayCore)
}

func (c *gatewayCore) Init(e *structs.Exposer, param ...string) error {
	c.e = e
	return nil
}

func (c *gatewayCore) Start() error {
	return c.startWatchDog()
}

func (c *gatewayCore) startWatchDog() error {
	listenIP := viper.GetString(ListenClientAddr)
	listenPort := viper.GetInt(ListenClientPort)

	logEntry := logrus.WithFields(logrus.Fields{
		"listen_ip":   listenIP,
		"listen_port": listenPort,
	})

	mo := &messageObserver{
		core: c,
	}
	co := &connectObserver{}

	// TODO  id 分配器
	c.dog = c.e.WatchDogFactory.NewWatchDog(nil, mo, co)
	if c.dog == nil {
		logEntry.Error("创建 watchdog 失败")
		return fmt.Errorf("创建 watchdog 失败")
	}
	logEntry.Info("准备监听")

	addr := fmt.Sprintf("%s:%d", listenIP, listenPort)
	return c.dog.Start(addr, net.TCP)
}