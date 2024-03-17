package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config 配置
	Config struct {
		Global `yaml:"global"`
		Http   `yaml:"http"`
	}
	// Global 全局配置
	Global struct {
		Namespace string `yaml:"namespace"` // 命名空间
		EnvName   string `yaml:"env_name"`  // 环境名称
	}
	// Http http监听
	Http struct {
		Addr string `yaml:"addr"` // 监听地址
	}
)

// NewConfig 返回程序配置
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
