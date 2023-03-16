package blog_handles

import (
	"github.com/baker-yuan/go-blog/all_packaged_library/util"
	"github.com/baker-yuan/go-blog/blog/application/cqe"
	"github.com/baker-yuan/go-blog/blog/domain/service"

	// "github.com/baker-yuan/go-blog/all_packaged_library/util"
	"github.com/baker-yuan/go-blog/blog/application/dto"
	app "github.com/baker-yuan/go-blog/blog/application/service"
	"github.com/baker-yuan/go-blog/blog/domain/repo"
	"github.com/baker-yuan/go-blog/middleware"
	"github.com/gin-gonic/gin"
)

// ArticleController 文章服务
type ArticleController struct {
	app app.ArticleApp
	srv service.ArticleService
}

func ArticleRegister(group *gin.RouterGroup, articleRepo repo.ArticleRepo) {
	article := ArticleController{
		app: app.NewArticleApp(articleRepo),
		srv: service.NewArticleService(articleRepo),
	}
	// 管理员
	group.GET("/admin/articles", article.ListArticleBacks)
	group.POST("/admin/articles", article.SaveOrUpdateArticle)
	group.PUT("/admin/articles/top", article.UpdateArticleTop)
	// group.PUT("/admin/articles", article.UpdateArticleDelete)
	// group.POST("/admin/articles/images", article.SaveArticleImages)
	// group.DELETE("/admin/articles", article.DeleteArticles)
	// group.GET("/admin/articles/{articleId}", article.GetArticleBackById)
	// // 用户
	// group.GET("/articles", article.ListArticles)
	group.GET("/articles/archives", article.ListArchives)
	// group.GET("/articles/:articleId", article.GetArticleById)
	// group.GET("/articles/condition", article.ListArticlesByCondition)
	// group.GET("/articles/search", article.ListArticlesBySearch)
	// group.POST("/articles/{articleId}/like", article.SaveArticleLike)
}

func (a ArticleController) ListArticleBacks(ginCtx *gin.Context) {
	var (
		currentPage uint32
		pageSize    uint32
		params      *cqe.ConditionVO
	)
	var (
		articles []*dto.ArticleBackDTO
		total    uint32
		err      error
	)
	// 条件
	currentPage, pageSize = util.GetPage(ginCtx)
	params = &cqe.ConditionVO{}
	if err := params.BindValidParam(ginCtx); err != nil {
		middleware.ResponseError(ginCtx, middleware.FAIL)
		return
	}
	// 搜索
	articles, total, err = a.srv.ListArticleBacks(*params, currentPage, pageSize)
	middleware.SendPageResult(ginCtx, articles, total, err)
}

func (a ArticleController) SaveOrUpdateArticle(ginCtx *gin.Context) {

}

func (a ArticleController) UpdateArticleTop(ginCtx *gin.Context) {
	var (
		params = &cqe.ArticleTopVO{}
	)
	var (
		err error
	)
	// 参数绑定
	if err := params.BindValidParam(ginCtx); err != nil {
		middleware.ResponseError(ginCtx, middleware.FAIL)
		return
	}
	// 搜索
	err = a.app.UpdateArticleTop(*params)
	middleware.SendResult(ginCtx, nil, err)
}

//
// func (a ArticleController) UpdateArticleDelete(ginCtx *gin.Context) {
//
// }
//
// func (a ArticleController) SaveArticleImages(ginCtx *gin.Context) {
//
// }
//
// func (a ArticleController) DeleteArticles(ginCtx *gin.Context) {
//
// }
//
// func (a ArticleController) GetArticleBackById(ginCtx *gin.Context) {
//
// }

func (a ArticleController) ListArchives(ginCtx *gin.Context) {
	var (
		currentPage uint32
		pageSize    uint32
	)
	var (
		err      error
		total    uint32
		archives []*dto.ArchiveDTO
	)
	currentPage, pageSize = util.GetPage(ginCtx)
	archives, total, err = a.app.ListArchives(currentPage, pageSize)
	middleware.SendPageResult(ginCtx, archives, total, err)
}

//
// func (a ArticleController) ListArticles(ginCtx *gin.Context) {
// 	var (
// 		currentPage uint32
// 		pageSize    uint32
// 		articles    []*dto.ArticleHomeDTO
// 		err         error
// 	)
// 	currentPage, pageSize = getPage(ginCtx)
// 	articles, err = a.srv.ListArticles(currentPage, pageSize)
// 	if err != nil {
// 		middleware.ResponseError(ginCtx, middleware.FAIL)
// 		return
// 	}
// 	middleware.ResponseSuccess(ginCtx, articles)
// }
//
// // GetArticleById godoc
// //
// //	@Summary		根据id查看文章
// //	@Description	根据id查看文章
// //	@Tags			文章模块
// //	@Accept			json
// //	@Produce		json
// //	@Param			articleId	path		int	true	"文章ID"
// //	@Success		200			{object}	middleware.Response{data=dto.ArticleDTO}
// //	@Router			/articles/{articleId} [get]
// func (a ArticleController) GetArticleById(ginCtx *gin.Context) {
// 	var (
// 		articleId uint32
// 		article   *dto.ArticleDTO
// 		err       error
// 	)
// 	articleId = util.StrToUInt32(ginCtx.Param("articleId"))
// 	article, err = a.srv.GetArticleById(articleId)
// 	sendResult(ginCtx, article, err)
// }
//
// func (a ArticleController) ListArticlesByCondition(ginCtx *gin.Context) {
//
// }
//
// func (a ArticleController) ListArticlesBySearch(ginCtx *gin.Context) {
//
// }
//
// func (a ArticleController) SaveArticleLike(ginCtx *gin.Context) {
//
// }
