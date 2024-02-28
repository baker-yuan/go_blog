package service

import (
	"github.com/baker-yuan/go-blog/blog/application/cqe"
	"github.com/baker-yuan/go-blog/blog/domain/entity"
	"github.com/baker-yuan/go-blog/blog/domain/port"

	// "github.com/baker-yuan/go-blog/all_packaged_library/util"
	"github.com/baker-yuan/go-blog/blog/application/dto"
)

// ArticleApp 文章服务
type ArticleApp struct {
	ArticlePort port.ArticlePort
}

// func NewArticleApp(repo port.ArticleRepo) ArticleApp {
// 	return ArticleApp{
// 		articleRepo: repo,
// 	}
// }

// ListArchives 分页查询文章归档
//
//	@return	文章归档
func (a *ArticleApp) ListArchives(currentPage uint32, pageSize uint32) (archives []*dto.ArchiveDTO, total uint32, err error) {
	return a.ArticlePort.ListArchives(currentPage, pageSize)
}

// // ListArticles 查询首页文章
// //
// //	@return	文章列表
// func (a *ArticleApp) ListArticles(currentPage uint32, pageSize uint32) ([]*ArticleHomeDTO, error) {
// 	var (
// 		current      uint32
// 		articles     []*ArticleHomeDTO
// 		dbArticles   []entity.Article
// 		search       map[string]interface{}
// 		articleCount int64
// 	)
// 	// 分页查询文章信息
// 	current = (currentPage - 1) * pageSize
// 	search = map[string]interface{}{
// 		"status":    1,
// 		"is_delete": 0,
// 	}
// 	lib.DB.Table("tb_article").
// 		Offset(int(current)).Limit(int(pageSize)).
// 		Where(search).
// 		Find(&dbArticles).
// 		Offset(-1).Limit(-1).Count(&articleCount)
//
// 	// 补充标签、分类
//
// 	// 组装数据
// 	util.DeepCopyByJson(dbArticles, &articles)
// 	return articles, nil
// }
//
// // ListArticlesByCondition 根据条件查询文章列表
// //
// // param condition 条件
// // return 文章列表
// func (a *ArticleApp) ListArticlesByCondition(condition ConditionVO) *ArticlePreviewListDTO {
// 	return nil
// }
//
// // ListArticlesBySearch 搜索文章
// //
// // param condition 条件
// // return 文章列表
// func (a *ArticleApp) ListArticlesBySearch(condition ConditionVO) []*ArticleSearchDTO {
// 	return nil
// }
//
// // ArticleVO 根据id查看后台文章
// //
// // param articleId 文章id
// // return 文章列表
// func (a *ArticleApp) getArticleBackById(articleId uint32) *ArticleVO {
// 	return nil
// }
//
// // GetArticleById 根据id查看文章
// //
// // param articleId 文章id
// // return 文章信息
// func (a *ArticleApp) GetArticleById(articleId uint32) (*ArticleDTO, error) {
// 	var (
// 		article    entity.Article
// 		articleDTO ArticleDTO
// 	)
// 	// 查询文章详情
// 	lib.DB.Raw("select * from tb_article where id = ? ", articleId).Scan(&article)
//
// 	// 类型转换
// 	util.DeepCopyByJson(article, &articleDTO)
//
// 	articleDTO.TagDTOList = []TagDTO{}
// 	articleDTO.RecommendArticleList = []ArticleRecommendDTO{}
// 	articleDTO.NewestArticleList = []ArticleRecommendDTO{}
// 	return &articleDTO, nil
// }
//
// // SaveArticleLike 点赞文章
// //
// // param articleId 文章id
// func (a *ArticleApp) SaveArticleLike(articleId uint32) {}
//
// // SaveOrUpdateArticle 添加或修改文章
// //
// // param articleVO 文章信息
// func (a *ArticleApp) SaveOrUpdateArticle(articleVO ArticleVO) {
//
// }

// UpdateArticleTop 修改文章置顶
//
// param articleTopVO 文章置顶信息
func (a *ArticleApp) UpdateArticleTop(articleTopVO cqe.ArticleTopVO) error {
	var (
		article *entity.Article
		err     error
	)
	if article, err = a.ArticlePort.FindById(articleTopVO.ID); err != nil {
		return err
	}
	article.Top()
	if err = a.ArticlePort.Save(article); err != nil {
		return err
	}
	return nil
}

//
// // UpdateArticleDelete
// //
// // param deleteVO 逻辑删除对象
// func (a *ArticleApp) UpdateArticleDelete(deleteVO DeleteVO) {
//
// }
//
// // DeleteArticles 物理删除文章
// //
// // param  articleIdList 文章id集合
// func (a *ArticleApp) DeleteArticles(articleIdList []uint32) {
//
// }
