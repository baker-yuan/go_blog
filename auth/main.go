package main

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	"github.com/baker-yuan/go-blog/all_packaged_library/util"
	"github.com/baker-yuan/go-blog/auth/ui/http"
	"github.com/baker-yuan/go-blog/auth/ui/http/menu_handles"
)

func main() {
	base.Init()
	http.NewHttp(initMenuService())
	util.QuitSignal(func() {
		log.Info(context.Background(), "server exit")
	})
}

func initMenuService() menu_handles.MenuUI {
	return nil
}
