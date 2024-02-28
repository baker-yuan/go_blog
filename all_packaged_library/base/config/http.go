package config

// HttpConf http
type HttpConf struct {
	Addr           uint32 `yaml:"addr" mapstructure:"addr"`                         // 监听地址, default ":8700"
	ReadTimeout    int    `yaml:"read_timeout" mapstructure:"read_timeout"`         // 读取超时时长
	WriteTimeout   int    `yaml:"write_timeout" mapstructure:"write_timeout"`       // 写入超时时长
	MaxHeaderBytes int    `yaml:"max_header_bytes" mapstructure:"max_header_bytes"` // 最大的header大小，二进制位长度
}
