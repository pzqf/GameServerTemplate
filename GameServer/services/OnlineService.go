package services

import (
	"ServerTemplate/GameServer/config"
	"context"
	"fmt"
	"github.com/pzqf/zEngine/zEtcd"
	"github.com/pzqf/zEngine/zLog"
	"github.com/pzqf/zEngine/zObject"
	"time"
)

type OnlineService struct {
	zObject.BaseObject
	etcdCli   *zEtcd.Client
	ctxCancel context.CancelFunc
	ctx       context.Context
	etcdKey   string
}

func NewOnlineService() *OnlineService {
	a := &OnlineService{}
	a.SetId(ServiceIdOnlineService)
	return a
}

func (os *OnlineService) Init() error {
	os.ctx, os.ctxCancel = context.WithCancel(context.Background())

	cli, err := zEtcd.NewEtcdClient(&zEtcd.ClientConfig{
		Endpoints: config.GConfig.EtcdServer,
	})

	os.etcdKey = fmt.Sprintf("/server/game_server/online/%d", config.GConfig.Server.Id)

	uploadInfo := "online"

	_, err = cli.PutWithTTL(context.Background(), os.etcdKey, uploadInfo, int64(config.GConfig.Server.Heartbeat*3))
	if err != nil {
		return err
	}

	os.etcdCli = cli

	return nil
}

func (os *OnlineService) Close() error {
	zLog.Info("online service close")
	os.ctxCancel()
	return nil
}

func (os *OnlineService) Serve() {
	go func() {
		defer func() {
			_ = os.etcdCli.Delete(context.Background(), os.etcdKey)
			_ = os.etcdCli.Close()
		}()
		for {
			select {
			case <-time.After(time.Duration(config.GConfig.Server.Heartbeat) * time.Second):
				uploadInfo := "online"
				_, err := os.etcdCli.PutWithTTL(context.Background(), os.etcdKey, uploadInfo, int64(config.GConfig.Server.Heartbeat*3))
				if err != nil {
					return
				}
			case <-os.ctx.Done():
				return
			}
		}
	}()
}
