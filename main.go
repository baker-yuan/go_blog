package main

//
// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"syscall"
//
// 	"github.com/baker-yuan/go-blog/all_packaged_library/lib"
// 	"github.com/baker-yuan/go-blog/router"
// )
//
// var (
// 	// config 对应配置文件夹
// 	// 开发 ./conf/dev/
// 	// 线上 ./conf/prod/
// 	config = flag.String("config", "", "input config file like ./conf/dev/")
// )
//
// func main() {
// 	// 初始化配置
// 	flag.Parse()
// 	lib.Init(*config)
// 	defer lib.Destroy()
//
// 	fmt.Println("http://localhost:8848/swagger/index.html")
//
// 	// 启动服务
// 	router.HttpServerRun()
//
// 	// 监听退出
// 	quit := make(chan os.Signal)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
//
// 	// 关闭服务
// 	router.HttpServerStop()
// }
