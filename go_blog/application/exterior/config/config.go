package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config 配置
	Config struct {
		Oss *Oss `yaml:"oss"`
	}
	// Oss 文件上传
	Oss struct {
		// 你的AccessKeyID
		AccessKeyID string `yaml:"access_key_id"`
		// 你的AccessKeySecret
		AccessKeySecret string `yaml:"access_key_secret"`
		// 你的OSS Bucket名称
		BucketName string `yaml:"bucket_name"`
		// 你的OSS Endpoint
		Endpoint string `yaml:"endpoint"`
		// 回调
		CallbackUrl string `yaml:"callback_url"`
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
