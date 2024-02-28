package routers

import (
	middleware "github.com/baker-yuan/go-blog/all_packaged_library/middleware/gin"
	"github.com/baker-yuan/go-blog/blog/ui/http/blog_handles"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine, blogSrv blog_handles.ArticleUI) {
	middleware.SetCorsRouters(r)
	middleware.SetRecovery(r)
}
