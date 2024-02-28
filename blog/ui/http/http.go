// Package http http接口
package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/config"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/blog/ui/http/blog_handles"
	"github.com/baker-yuan/go-blog/blog/ui/http/routers"
	"github.com/gin-gonic/gin"
)

func NewHttp(blogSrv blog_handles.ArticleUI) {
	var (
		httpConf = config.GetHttpConf()
	)
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	routers.SetRouters(g, blogSrv)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", httpConf.Addr),
		Handler:        g,
		ReadTimeout:    time.Duration(httpConf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(httpConf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info(context.Background(), "server start success port: %d, pid: %d", httpConf.Addr, os.Getpid())
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
