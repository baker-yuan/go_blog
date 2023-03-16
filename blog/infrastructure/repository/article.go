package repository

import (
	"github.com/baker-yuan/go-blog/all_packaged_library/constant"
	"github.com/baker-yuan/go-blog/all_packaged_library/lib"
	"github.com/baker-yuan/go-blog/all_packaged_library/util"
	"github.com/baker-yuan/go-blog/blog/application/cqe"
	"github.com/baker-yuan/go-blog/blog/application/dto"
	"github.com/baker-yuan/go-blog/blog/domain/entity"
	"github.com/baker-yuan/go-blog/blog/infrastructure/repository/do"
	"gorm.io/gorm"
)

type ArticleRepoImpl struct {
}

func (a *ArticleRepoImpl) articleSearch(search map[string]interface{}, currentPage uint32, pageSize uint32) ([]*do.ArticleDO, uint32, error) {
	var (
		articles = make([]*do.ArticleDO, 0)
		total    int64
	)
	var (
		current uint32
		tx      *gorm.DB
	)
	current = (currentPage - 1) * pageSize
	tx = lib.DB.Table("tb_article").Offset(int(current)).Limit(int(pageSize))
	tx = util.BuildSearch(tx, search)
	tx = tx.Find(&articles).Offset(-1).Limit(-1).Count(&total)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return articles, uint32(total), nil
}

func (a *ArticleRepoImpl) ListArchives(currentPage uint32, pageSize uint32) ([]*dto.ArchiveDTO, uint32, error) {
	var (
		articles []*dto.ArchiveDTO
		total    uint32
		err      error
	)
	var (
		dbArticles []*do.ArticleDO
		search     map[string]interface{}
	)
	search = map[string]interface{}{
		"is_delete": constant.FALSE,
		"status":    entity.ARTICLE_STATUS_PUBLIC,
	}
	dbArticles, total, err = a.articleSearch(search, currentPage, pageSize)
	if err != nil {
		return nil, 0, err
	}
	util.DeepCopyByJson(dbArticles, &articles)
	return articles, total, nil
}

func (a *ArticleRepoImpl) ListArticleBacks(condition cqe.ConditionVO, currentPage uint32, pageSize uint32) ([]*dto.ArticleBackDTO, uint32, error) {
	var (
		articles []*dto.ArticleBackDTO
		total    uint32
		err      error
	)
	var (
		dbArticles  []*do.ArticleDO
		search      map[string]interface{}
		articleTag  map[uint32][]dto.TagDTO
		category    map[uint32]string
		articleIds  = make([]uint32, 0)
		categoryIds = make([]uint32, 0)
	)
	search = map[string]interface{}{
		"is_delete":          condition.IsDelete,
		"article_title_like": condition.Keywords,
		"category_id":        condition.CategoryId,
	}
	// 标签id转博客id
	if condition.TagId != nil {
		articleIds := make([]uint32, 0)
		lib.DB.Table("tb_article_tag").
			Raw("select article_id from tb_article_tag where tag_id = ?", condition.TagId).
			Scan(&articleIds)
		if len(articleIds) != 0 {
			search["id_in"] = articleIds
		}
	}
	dbArticles, total, err = a.articleSearch(search, currentPage, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 填充分类+标签
	for _, v := range dbArticles {
		articleIds = append(articleIds, v.ID)
		categoryIds = append(categoryIds, v.CategoryId)
	}
	category = FindCategoryByCategoryIds(categoryIds)
	articleTag = FindTagByArticleIds(articleIds)
	for _, v := range dbArticles {
		item := dto.ArticleBackDTO{}
		util.DeepCopyByJson(v, &item)
		item.CategoryName = category[v.CategoryId]
		item.TagDTOList = articleTag[v.ID]
		articles = append(articles, &item)
	}
	return articles, total, nil
}

func FindTagByArticleIds(articleIds []uint32) map[uint32][]dto.TagDTO {
	type Temp struct {
		ArticleID uint32 `gorm:"column:articleId"`
		TagID     uint32 `gorm:"column:tagId"`
		TagName   string `gorm:"column:tagName"`
	}
	var (
		raws   = make([]*Temp, 0)
		result = make(map[uint32][]dto.TagDTO, 0)
	)
	lib.DB.Raw("select a.article_id as articleId, a.tag_id as tagId, b.tag_name as tagName from tb_article_tag a join tb_tag b on a.tag_id = b.id where a.article_id in ?", articleIds).Scan(&raws)

	for _, v := range raws {
		tagDto := dto.TagDTO{
			ID:      v.TagID,
			TagName: v.TagName,
		}
		if _, exist := result[v.ArticleID]; exist {
			result[v.ArticleID] = append(result[v.ArticleID], tagDto)
		} else {
			result[v.ArticleID] = []dto.TagDTO{tagDto}
		}
	}

	return result
}

func FindCategoryByCategoryIds(categoryIds []uint32) map[uint32]string {
	var (
		category = make([]do.Category, 0)
		result   = make(map[uint32]string, 0)
	)
	lib.DB.Raw("select * from tb_category where id in ?", categoryIds).
		Scan(&category)
	for _, v := range category {
		result[v.ID] = v.CategoryName
	}
	return result
}

func (a *ArticleRepoImpl) FindById(articleID uint32) (*entity.Article, error) {
	var (
		dbArticle = do.ArticleDO{}
		tx        *gorm.DB
	)
	var (
		article = entity.Article{}
	)
	tx = lib.DB.Raw("select * from tb_article where id in =", articleID).
		Scan(&dbArticle)
	if tx.Error != nil {
		return nil, tx.Error
	}
	util.DeepCopyByJson(dbArticle, &article)
	return &article, nil
}

func (a *ArticleRepoImpl) Save(article *entity.Article) error {
	return nil
}
