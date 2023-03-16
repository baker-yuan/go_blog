package config

// LogConf 日志
type LogConf struct {
	Level string `yaml:"level" mapstructure:"level"`
	Path  string `yaml:"path" mapstructure:"path"`
}
