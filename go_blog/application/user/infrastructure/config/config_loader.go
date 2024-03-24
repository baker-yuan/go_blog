package config

import (
	"github.com/baker-yuan/go-blog/common/jwt"
	"github.com/spf13/viper"
)

// AppConfig 保存应用程序配置
type AppConfig struct {
	JWT jwt.TokenGeneratorConfig `mapstructure:"jwt"` // JWT相关配置
}

// LoadConfig 从指定的文件路径加载应用程序配置。
func LoadConfig(configFilePath string) (*AppConfig, error) {
	viper.SetConfigFile(configFilePath) // 设置配置文件路径
	viper.AutomaticEnv()                // 使用环境变量覆盖配置
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	// 解析配置到结构
	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
