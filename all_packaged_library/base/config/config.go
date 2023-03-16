package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	baseConf  *BaseConf
	httpConf  *HttpConf
	logConf   *LogConf
	mySQLConf map[string]*MySQLConf
	redisConf map[string]*RedisConf
)
var (
	env string // 配置环境名 比如：dev prod test
)

// AppConfig app.yaml
type AppConfig struct {
	Base  *BaseConf             `yaml:"base" mapstructure:"base"`
	Http  *HttpConf             `yaml:"http" mapstructure:"http"`
	Log   *LogConf              `yaml:"log" mapstructure:"log"`
	Mysql map[string]*MySQLConf `yaml:"mysql" mapstructure:"mysql"`
	Redis map[string]*RedisConf `yaml:"redis" mapstructure:"redis"`
}

// HttpConf http
type HttpConf struct {
	Addr           string `yaml:"addr" mapstructure:"addr"`                         //  监听地址, default ":8700"
	ReadTimeout    int    `yaml:"read_timeout" mapstructure:"read_timeout"`         // 读取超时时长
	WriteTimeout   int    `yaml:"write_timeout" mapstructure:"write_timeout"`       // 写入超时时长
	MaxHeaderBytes int    `yaml:"max_header_bytes" mapstructure:"max_header_bytes"` // 最大的header大小，二进制位长度
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

// GetConfEnv 获取配置环境名
func GetConfEnv() string {
	return env
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

// DoInit 配置初始化
// confEnvPath 配置文件夹
func doInit(confEnvPath string, confEnv string) {
	// 读取文件
	bts, err := os.ReadFile(confEnvPath + "/" + "app.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("[INFO] load app.yaml \n%s\n", string(bts))
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
	env = confEnv
}

// Init 解析配置文件目录
//
// 配置文件必须放到一个文件夹中
// 如：config=conf/dev/app.yaml 	ConfEnvPath=conf/dev	ConfEnv=dev
func Init(config string) {
	path := strings.Split(config, "/")
	confEnvPath := strings.Join(path[:len(path)-1], "/")
	confEnv := path[len(path)-2]
	doInit(confEnvPath, confEnv)
}
