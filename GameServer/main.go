package main

import (
	"ServerTemplate/GameServer/config"
	"ServerTemplate/GameServer/services"
	"flag"
	"github.com/pzqf/zEngine/zLog"
	"github.com/pzqf/zEngine/zService"
	"github.com/pzqf/zEngine/zSignal"
	"go.uber.org/zap"
	"log"
)

func main() {
	log.Println("server start...")
	//初始化程序配置
	configFile := flag.String("c", "config.toml", "set configuration `file`")
	flag.Parse()

	err := config.InitDefaultConfig(*configFile)
	if err != nil {
		log.Println(err)
		return
	}

	//初始化日志
	err = zLog.InitLogger(&config.GConfig.Logger)
	if err != nil {
		log.Println(err)
		return
	}

	zLog.Info(`server start....`)

	//初始化配置数据

	//初始化各模块
	/*
		zKeyWordFilter.InitDefaultFilter()
		err = zKeyWordFilter.ParseFromFile(`keyword.txt`)
		if err != nil {
			zLog.Error("KeyWordFilter.ParseFromFile error ", zap.Error(err))
			return
		}
	*/
	//初始化各服务
	sm := zService.ServiceManager{}

	if err = sm.AddService(services.NewTcpService()); err != nil {
		zLog.Error("add service TcpService failed ", zap.Error(err))
		return
	}
	if err = sm.AddService(services.NewHttpService()); err != nil {
		zLog.Error("add service HttpService failed ", zap.Error(err))
		return
	}
	if err = sm.AddService(services.NewRpcService()); err != nil {
		zLog.Error("add service RpcService failed ", zap.Error(err))
		return
	}

	sm.InitServices()
	sm.ServeServices()

	zSignal.GracefulExit()
	zLog.Info("server will be shutdown")
	sm.CloseServices()
	zLog.Info("server exit")
}
