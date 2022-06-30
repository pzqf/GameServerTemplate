package services

import (
	"ServerTemplate/GameServer/config"
	"ServerTemplate/GameServer/module"
	"fmt"
	"github.com/pzqf/zEngine/zLog"
	"github.com/pzqf/zEngine/zNet"
	"github.com/pzqf/zEngine/zObject"
	"go.uber.org/zap"
)

type TcpService struct {
	zObject.BaseObject
	tcpServer *zNet.TcpServer
}

func NewTcpService() *TcpService {
	a := &TcpService{}
	a.SetId(ServiceIdTcpServer)
	return a
}

func (ts *TcpService) Init() error {
	netCfg := zNet.Config{
		MaxPacketDataSize: zNet.DefaultPacketDataSize,
		ListenAddress:     fmt.Sprintf(":%d", config.GConfig.TcpServer.Port),
		MaxClientCount:    config.GConfig.TcpServer.MaxClientCount,
		ChanSize:          256,
		HeartbeatDuration: 30,
	}

	ts.tcpServer = zNet.NewTcpServer(&netCfg)
	ts.tcpServer.SetRemoveSessionCallBack(module.GetPlayerManager().OnRemoveSession)

	zLog.Info("tcp server init success", zap.Int("maxClientCount", 10000),
		zap.Int32("MaxNetPacketDataSize", zNet.DefaultPacketDataSize),
	)
	return nil
}

func (ts *TcpService) Close() error {
	ts.tcpServer.Close()
	return nil
}

func (ts *TcpService) Serve() {
	_ = ts.tcpServer.Start()
}
