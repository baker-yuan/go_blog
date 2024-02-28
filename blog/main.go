package main

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/all_packaged_library/util"
	"github.com/baker-yuan/go-blog/blog/ui/http"
	"github.com/baker-yuan/go-blog/blog/ui/http/blog_handles"
)

func initMenuService() blog_handles.ArticleUI {
	return nil
}

func main() {
	base.Init()
	http.NewHttp(initMenuService())
	util.QuitSignal(func() {
		log.Info(context.Background(), "server exit")
	})
}
