package config

import (
	"fmt"
	"os"

	"github.com/baker-yuan/go-blog/go_blog/common/util"
	"github.com/ilyakaznacheev/cleanenv"
	"trpc.group/trpc-go/trpc-go/log"
)

type (
	// Config 配置
	Config struct {
		Oss *Oss `yaml:"oss"`
	}
	// Oss 文件上传
	Oss struct {
		// 你的AccessKeyID
		AccessKeyID string `yaml:"access_key_id" env:"ALI_OSS_ACCESS_KEY_ID"` // 读取环境变量
		// 你的AccessKeySecret
		AccessKeySecret string `yaml:"access_key_secret" env:"ALI_OSS_ACCESS_KEY_SECRET"` // 读取环境变量
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
	// 配置文件读取
	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("load config by yml error: %w", err)
	}
	// 环境变量有设置就覆盖
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("load config by env error: %w", err)
	}
	cfg.Oss.AccessKeyID = os.Getenv("ALI_OSS_ACCESS_KEY_ID")
	cfg.Oss.AccessKeySecret = os.Getenv("ALI_OSS_ACCESS_KEY_SECRET")
	log.Debugf("load cfg: %s", util.StrUtils.Obj2Json(cfg))
	return cfg, nil
}
