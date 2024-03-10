package data

import (
	"context"

	pb "github.com/baker-yuan/go-blog/application/blog/api/blog/v1"
	"github.com/baker-yuan/go-blog/application/blog/internal/biz"
	"gorm.io/gorm"
)

type articleRepo struct {
	db *gorm.DB
}

// NewArticleRepo 创建文章数据库操作结构体
func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleRepo{
		db: data.gormDB,
	}
}

// GetArticleByID 根据文章id查询文章
func (r *articleRepo) GetArticleByID(ctx context.Context, id int) (*biz.Article, error) {
	//dest := &entity.Article{}
	//statement := r.getStatement().AndEqual(entity.ArticleFieldID, id)
	//if err := r.client.FindOne(ctx, statement, dest); err != nil {
	//	return nil, err
	//}
	//if dest.ID == 0 {
	//	return nil, retcode.BuildErrorFmtMsg(retcode.RetResourceNotExist)
	//}
	//return dest, nil

	return nil, nil
}

// GetArticleByIDs 根据文章id集合查询文章
func (r *articleRepo) GetArticleByIDs(ctx context.Context, ids []int) (biz.Articles, error) {
	//dest := make([]*entity.Article, 0)
	//if len(ids) == 0 {
	//	return dest, nil
	//}
	//var where = orm.WhereCond{entity.ArticleFieldID: ids}
	//var statement = r.getStatement().Where(where)
	//if err := r.client.FindAll(ctx, statement, &dest); err != nil {
	//	return nil, err
	//}
	//return dest, nil

	return nil, nil
}

// Save 保存文章
func (r *articleRepo) Save(ctx context.Context, article *biz.Article) (uint32, error) {
	//statement := r.getStatement().InsertStruct(article)
	//lastInsertId, err := r.client.Insert(ctx, statement)
	//if err != nil {
	//	return 0, err
	//}
	//return uint32(lastInsertId), nil

	return 0, nil
}

// UpdateByID 根据ID修改文章
func (r *articleRepo) UpdateByID(ctx context.Context, article *biz.Article) error {
	//if article.ID == 0 {
	//	return retcode.BuildErrorFmtMsg(retcode.IllegalArgument)
	//}
	//statement := r.getStatement().
	//	AndEqual(entity.ArticleFieldID, article.ID).
	//	UpdateStruct(article)
	//_, err := r.client.Update(ctx, statement)
	//if err != nil {
	//	return err
	//}
	return nil
}

// DeleteByID 根据ID删除文章
func (r *articleRepo) DeleteByID(ctx context.Context, id int) error {
	//statement := r.getStatement().AndEqual(entity.ArticleFieldID, id)
	//_, err := r.client.Delete(ctx, statement)
	//if err != nil {
	//	return err
	//}
	return nil
}

// SearchArticle 文章搜索
func (r *articleRepo) SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (biz.Articles, uint32, error) {
	// 查询
	var res []*biz.Article
	var pageTotal int64

	offset := (req.GetPageNum() - 1) * req.GetPageSize()

	tx := r.db.WithContext(ctx)

	tx = tx.Offset(int(offset)).Limit(int(req.GetPageSize())).Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}

	return res, uint32(pageTotal), nil
}
