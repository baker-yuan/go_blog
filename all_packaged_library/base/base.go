package base

import (
	"flag"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/config"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/db"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
)

var (
	// app.yaml配置文件路径
	// 开发 ./conf/dev/
	// 线上 ./conf/prod/
	configPath = flag.String("conf", "", "input config file like ./conf/dev/")
)

// Init 初始化配置，外部调用
// 公共初始化函数：支持两种方式设置配置文件
// 1、函数传入配置文件 Init("./conf/dev/app.yaml")
// 2、如果配置文件为空，会从命令行中读取 -conf=./conf/dev/app.yaml
func Init(path ...string) {
	if len(path) != 0 {
		configPath = &path[0]
	} else {
		flag.Parse()
	}
	// 1、初始化配置
	config.Init(*configPath)
	// 2、初始化日志
	log.InitLog()
	// 3、初始化数据库
	db.Init()
}
