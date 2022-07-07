package SensitiveWords

import (
	"ServerTemplate/GameServer/config"
	"context"
	"strings"
	"time"

	"github.com/pzqf/zEngine/zEtcd"
	"github.com/pzqf/zUtil/zKeyWordFilter"
)

var filter *zKeyWordFilter.DeaFilter

func Init() error {
	filter = zKeyWordFilter.NewFilter()

	cli, err := zEtcd.NewEtcdClient(&zEtcd.ClientConfig{
		Endpoints: config.GConfig.EtcdServer,
	})
	if err != nil {
		return err
	}
	key := "/keyword"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	configContent, err := cli.GetOne(ctx, key)
	cancel()
	if err != nil {
		return err
	}

	splitContent(configContent)

	watcher, eventCh := cli.Watch(context.Background(), key, false)
	go func() {
		for e := range eventCh {
			switch e.Event {
			case zEtcd.EventCreate, zEtcd.EventModify:
				configContent, err = cli.GetOne(context.Background(), key)
				if err != nil {
					continue
				}
				splitContent(configContent)
			case zEtcd.EventDelete:
				continue
			case zEtcd.EventWatchCancel:
				_ = watcher.Close()
				return
			}
		}
	}()

	return nil
}

func splitContent(configContent string) {
	configContent = strings.Replace(configContent, "\r", "", -1)
	wordList := strings.Split(configContent, "\n")
	for _, word := range wordList {
		filter.AddWord(word)
	}
}

// Filter 敏感词过滤
func Filter(str string) string {
	return filter.Filter(str)
}
