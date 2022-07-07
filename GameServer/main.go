package main

import (
	"ServerTemplate/GameServer/GameConfig"
	"ServerTemplate/GameServer/SensitiveWords"
	"ServerTemplate/GameServer/config"
	"ServerTemplate/GameServer/handler"
	"ServerTemplate/GameServer/services"
	"flag"
	"log"
	"strings"

	"github.com/pzqf/zEngine/zLog"
	"github.com/pzqf/zEngine/zNet"
	"github.com/pzqf/zEngine/zService"
	"github.com/pzqf/zEngine/zSignal"
	"go.uber.org/zap"
)

var sm = zService.ServiceManager{}

func Init(serverId int, etcdAddress []string) error {
	//获取程序配置及初始化日志================================
	//程序配置项，可从配置文件或ETCD获取
	err := config.InitDefaultConfigByEtcd(serverId, etcdAddress)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("config load from etcd server success ")

	//初始化日志
	err = zLog.InitLogger(&config.GConfig.Logger)
	if err != nil {
		log.Println(err)
		return err
	}

	zLog.Info(`server start....`)

	//初始化游戏配置数据
	GameConfig.LoadGameConfig("D:\\37818\\Documents")

	//初始化各模块==========================================
	//敏感词，来源于etcd
	err = SensitiveWords.Init()
	if err != nil {
		zLog.Error("Sensitive words error ", zap.Error(err))
		return err
	}

	//网络前置初始化=========================================
	zNet.SetLogPrintFunc(func(v ...any) {
		zLog.Info("zNet info", zap.Any("info", v))
	})

	err = handler.Init()
	if err != nil {
		zLog.Error("RegisterHandler error %d", zap.Error(err))
		return err
	}

	//初始化各服务=============================================
	if err = sm.AddService(services.NewTcpService()); err != nil {
		zLog.Error("add service TcpService failed ", zap.Error(err))
		return err
	}
	if err = sm.AddService(services.NewHttpService()); err != nil {
		zLog.Error("add service HttpService failed ", zap.Error(err))
		return err
	}
	if err = sm.AddService(services.NewRpcService()); err != nil {
		zLog.Error("add service RpcService failed ", zap.Error(err))
		return err
	}
	if err = sm.AddService(services.NewOnlineService()); err != nil {
		zLog.Error("add service OnlineService failed ", zap.Error(err))
		return err
	}

	sm.InitServices()
	return nil
}

func Serve() {
	sm.ServeServices()

	s, err := sm.GetService(services.ServiceIdOnlineService)
	if err == nil {
		_ = s.(*services.OnlineService).Update()
	}
}

func Stop() {
	zLog.Info("server will be shutdown")
	sm.CloseServices()
	zLog.Info("server exit")
}

func main() {
	log.Println("server start...")
	etcdAddr := flag.String("e", "127.0.0.1:2379", "etcd server addr, cluster by ','")
	serverId := flag.Int("i", 0, "server id")
	flag.Parse()

	if *serverId <= 0 {
		log.Println("server id is 0")
		return
	}

	err := Init(*serverId, strings.Split(*etcdAddr, ","))
	if err != nil {
		return
	}
	Serve()
	zSignal.GracefulExit()
	Stop()
}
