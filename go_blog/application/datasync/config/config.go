package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config 配置
	Config struct {
		BinLog   `yaml:"bin_log"`
		Mysql    `yaml:"mysql"`
		Consumer `yaml:"consumer"`
	}
	// BinLog binlog配置
	BinLog struct {
		PositionName string `yaml:"position_name"` // 在binlog文件的名称
		PositionPos  uint32 `yaml:"position_pos"`  // 在binlog文件中的位置（偏移量）
	}
	// Mysql mysql配置
	Mysql struct {
		// mysql配置
		Addr     string `yaml:"addr"`     // mysql地址
		User     string `yaml:"user"`     // mysql用户名
		Password string `yaml:"password"` // mysql密码
		// 监听库和表
		DbName string   `yaml:"db_name"` // 监听的数据
		TbName []string `yaml:"tb_name"` // 监听的表
	}

	// Consumer 消费者
	Consumer struct {
		BufferSize uint32 `yaml:"buffer_size"` // 通道的缓冲大小
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
