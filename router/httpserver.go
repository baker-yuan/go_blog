package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/baker-yuan/go-blog/all_packaged_library/lib"
	"github.com/baker-yuan/go-blog/middleware"
	"github.com/gin-gonic/gin"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	baseConfig := lib.App.Base
	httpConfig := lib.App.Http
	gin.SetMode(baseConfig.DebugMode)
	r := InitRouter(middleware.Cors())
	HttpSrvHandler = &http.Server{
		Addr:           httpConfig.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(httpConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(httpConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(httpConfig.MaxHeaderBytes),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", httpConfig.Addr)
		log.Printf(" [INFO] swagger http://127.0.0.1%s/swagger/index.html\n", httpConfig.Addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", httpConfig.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
