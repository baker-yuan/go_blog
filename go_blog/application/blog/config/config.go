package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config 配置
	Config struct {
		NewBlogLimitNum    uint32 `yaml:"new_blog_limit_num"`
		RandomBlogLimitNum uint32 `yaml:"random_blog_limit_num"`
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
