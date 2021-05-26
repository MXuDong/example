package server

import (
	"github.com/sirupsen/logrus"
	"github.io/MXuDong/example/config"
	"github.io/MXuDong/example/pkg/util/nets"
	"strconv"
)

// TcpServerStart will listen tcp on target port
func TcpServerStart() error {
	if !config.Config.ServerConfig.EnableTcpServer {
		return nil
	}

	tcpServer := nets.TcpServer{}
	err := tcpServer.Listen(config.Config.ServerConfig.TcpNetWork,
		config.Config.ServerConfig.TcpAddress+":"+strconv.Itoa(config.Config.ServerConfig.TcpServerPort),
		int64(config.Config.ServerConfig.MaxHandlerCount))
	tcpServer.RegisterHandler(nets.DefaultTcpHandler)
	if err != nil {
		return err
	}
	logrus.Infof("Tcp server already start on %v:%v",
		config.Config.ServerConfig.TcpAddress,
		strconv.Itoa(config.Config.ServerConfig.TcpServerPort))
	err = tcpServer.Do(func(err error) {
		logrus.Warnf("Tcp handler error : %v", err)
	})

	if err != nil {
	}
	return err
}
