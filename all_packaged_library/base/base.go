package base

import (
	"gitee.com/baker-yuan/go-blog/all_packaged_library/base/config"
	"gitee.com/baker-yuan/go-blog/all_packaged_library/base/db"
	"gitee.com/baker-yuan/go-blog/all_packaged_library/base/log"
)

// Init 初始化配置，外部调用
// 公共初始化函数：支持两种方式设置配置文件
// 1、函数传入配置文件 Init("./conf/dev/")
// 2、如果配置文件为空，会从命令行中读取 -config ./conf/dev/
func Init(configPath string) {
	config.Init(configPath)
	db.Init()
	log.InitLog()
}
