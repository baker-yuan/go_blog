package repo

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/common/db"
	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
	"gorm.io/gorm"
)

// ArticleRepo 文章Repo
type ArticleRepo struct {
	*db.GenericDao[entity.Article, uint32]
}

// NewArticleRepo 创建文章Repo
func NewArticleRepo(gormDB *gorm.DB) *ArticleRepo {
	return &ArticleRepo{
		GenericDao: db.NewGenericDao[entity.Article, uint32](gormDB),
	}
}

func init() {
	registerInitField(initArticleField)
}

var (
	// 全字段修改Article哪些字段不修改
	notUpdateArticleField = []string{
		entity.ArticleFieldCreateTime,
	}
	updateArticleField []string
)

// InitJsonDynamicConfigField 全字段修改
func initArticleField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Article{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateArticleField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateArticleField...)
	return nil
}

// GetArticleByID 根据文章id查询文章
func (r *ArticleRepo) GetArticleByID(ctx context.Context, id uint32) (*entity.Article, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetArticleByIDs 根据文章id集合查询文章
func (r *ArticleRepo) GetArticleByIDs(ctx context.Context, ids []uint32) (entity.Articles, error) {
	return r.GenericDao.GetByIDs(ctx, ids)
}

// Save 保存文章
func (r *ArticleRepo) Save(ctx context.Context, article *entity.Article) (uint32, error) {
	if article.ID > 0 {
		return 0, errors.New("illegal argument article id exist")
	}
	return article.ID, r.GenericDao.Create(ctx, article)
}

// UpdateByID 根据ID修改文章
func (r *ArticleRepo) UpdateByID(ctx context.Context, article *entity.Article) error {
	if article.ID == 0 {
		return errors.New("illegal argument article exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateArticleField).Updates(article).Error
}

// DeleteByID 根据ID删除文章
func (r *ArticleRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchArticle 文章搜索
func (r *ArticleRepo) SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (entity.Articles, uint32, error) {
	var (
		res       []*entity.Article
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.DB.WithContext(ctx),
		func(search map[string]*db.SearchValue) {

		},
	)
	if err != nil {
		return nil, 0, err
	}
	tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).
		Limit(int(req.GetPageSize())).Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}
