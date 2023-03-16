package config

// BaseConf 基础配置
type BaseConf struct {
	DebugMode    string `yaml:"debug_mode" mapstructure:"debug_mode"`
	TimeLocation string `yaml:"time_location" mapstructure:"time_location"`
	AppName      string `yaml:"app_name" mapstructure:"app_name"`
}
