package config

import (
	toml "github.com/pelletier/go-toml"
	"github.com/pzqf/zEngine/zLog"
	"log"
	"path/filepath"
)

type ServerConfig struct {
	Port           int   `toml:"port" json:"port"`
	MaxClientCount int32 `toml:"max_client_count" json:"max_client_count"`
}

type Config struct {
	Server ServerConfig `toml:"server" json:"server"`
	Logger zLog.Config  `toml:"log" json:"log"`
}

func Load(configFile string) (*Config, error) {
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

var GConfig *Config

func InitDefaultConfig(configFile string) error {
	cfg, err := Load(configFile)
	if err != nil {
		return err
	}
	GConfig = cfg
	return nil
}
