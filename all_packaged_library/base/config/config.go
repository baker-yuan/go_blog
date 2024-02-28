package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	baseConf  *BaseConf
	httpConf  *HttpConf
	logConf   *LogConf
	mySQLConf map[string]*MySQLConf
	redisConf map[string]*RedisConf
)

// AppConfig app.yaml
type AppConfig struct {
	Base  *BaseConf             `yaml:"base" mapstructure:"base"`
	Http  *HttpConf             `yaml:"http" mapstructure:"http"`
	Log   *LogConf              `yaml:"log" mapstructure:"log"`
	Mysql map[string]*MySQLConf `yaml:"mysql" mapstructure:"mysql"`
	Redis map[string]*RedisConf `yaml:"redis" mapstructure:"redis"`
}

// RedisConf redis配置
type RedisConf struct {
	ProxyList    []string `yaml:"proxy_list" mapstructure:"proxy_list"`       // redis集群地址
	Password     string   `yaml:"password" mapstructure:"password"`           // 密码
	Db           int      `yaml:"db" mapstructure:"db"`                       // 指定数据库
	ConnTimeout  int      `yaml:"conn_timeout" mapstructure:"conn_timeout"`   // 连接超时
	ReadTimeout  int      `yaml:"read_timeout" mapstructure:"read_timeout"`   // 读取超时
	WriteTimeout int      `yaml:"write_timeout" mapstructure:"write_timeout"` // 写入超时
}

func GetMySQLConf() map[string]*MySQLConf {
	return mySQLConf
}

func GetLogConf() *LogConf {
	return logConf
}
func GetBaseConf() *BaseConf {
	return baseConf
}

func GetHttpConf() *HttpConf {
	return httpConf
}

// Init 配置初始化
func Init(confEnvPath string) {
	// 读取文件
	bts, err := os.ReadFile(confEnvPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[INFO] load config \n%s\n", string(bts))
	// 解析yaml
	viperParse := viper.New()
	viperParse.SetConfigType("yaml")
	viperParse.ReadConfig(bytes.NewBuffer(bts))
	// 反序列化
	var app AppConfig
	err = viperParse.Unmarshal(&app)
	if err != nil {
		panic(err)
	}
	// 赋值
	baseConf = app.Base
	httpConf = app.Http
	logConf = app.Log
	mySQLConf = app.Mysql
	redisConf = app.Redis
}
