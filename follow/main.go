package main

import (
	"flag"

	"gitee.com/baker-yuan/go-blog/all_packaged_library/base"
)

var (
	// config 对应配置文件夹
	// 开发 ./conf/dev/
	// 线上 ./conf/prod/
	config = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	base.Init(config)
}
