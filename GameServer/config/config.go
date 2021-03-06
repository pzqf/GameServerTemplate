package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/pelletier/go-toml"
	"github.com/pzqf/zEngine/zEtcd"
	"github.com/pzqf/zEngine/zLog"
)

type ServerConfig struct {
	Id        int    `toml:"id" json:"id"`
	Name      string `toml:"name" json:"name"`
	Heartbeat int    `toml:"heartbeat" json:"heartbeat"`
}

type TcpServerConfig struct {
	Port           int   `toml:"port" json:"port"`
	MaxClientCount int32 `toml:"max_client_count" json:"max_client_count"`
}

type HttpServerConfig struct {
	Addr string `toml:"addr" json:"addr"`
}

type Config struct {
	Server     ServerConfig     `toml:"server" json:"server"`
	TcpServer  TcpServerConfig  `toml:"tcp_server" json:"tcp_server"`
	HttpServer HttpServerConfig `toml:"http_server" json:"http_server"`
	Logger     zLog.Config      `toml:"log" json:"log"`
	EtcdServer []string
}

func LoadFile(configFile string) (*Config, error) {
	filePath, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	log.Println("Parse config file! path: ", filePath)
	tree, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := tree.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func Load(content string) (*Config, error) {
	tree, err := toml.Load(content)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := tree.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

var GConfig *Config

func InitDefaultConfig(configFile string) error {
	cfg, err := LoadFile(configFile)
	if err != nil {
		return err
	}
	GConfig = cfg
	return nil
}

func InitDefaultConfigByEtcd(serverId int, etcdAddress []string) error {
	cli, err := zEtcd.NewEtcdClient(&zEtcd.ClientConfig{
		Endpoints: etcdAddress,
	})
	if err != nil {
		return err
	}
	key := fmt.Sprintf("/config/game_server/%d", serverId)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	configContent, err := cli.GetOne(ctx, key)
	cancel()
	if err != nil {
		return err
	}

	cfg, err := Load(configContent)
	if err != nil {
		return err
	}
	if cfg.Server.Id != serverId {
		return errors.New("server id not match")
	}
	cfg.EtcdServer = etcdAddress
	GConfig = cfg

	watcher, eventCh := cli.Watch(context.Background(), key, false)
	go func() {
		for e := range eventCh {
			switch e.Event {
			case zEtcd.EventCreate, zEtcd.EventModify:
				configContent, err = cli.GetOne(context.Background(), key)
				if err != nil {
					continue
				}
				cfg, err = Load(configContent)
				if err != nil {
					continue
				}
				GConfig = cfg
				log.Println("config reload", e.Data, "maybe need restart")
			case zEtcd.EventDelete:
				log.Println("delete", e.Data)

				continue
			case zEtcd.EventWatchCancel:
				_ = watcher.Close()
				return
			}
		}
	}()

	return nil
}
