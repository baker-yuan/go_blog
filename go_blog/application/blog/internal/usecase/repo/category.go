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

// CategoryRepo 文章分类
type CategoryRepo struct {
	*db.GenericDao[entity.Category, uint32]
}

// NewCategoryRepo 创建
func NewCategoryRepo(gormDB *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		GenericDao: &db.GenericDao[entity.Category, uint32]{
			DB: gormDB,
		},
	}
}

func init() {
	registerInitField(initCategoryField)
}

var (
	// 全字段修改Category那些字段不修改
	notUpdateCategoryField = []string{
		"created_at",
		"create_time",
	}
	updateCategoryField []string
)

// InitCategoryField 全字段修改
func initCategoryField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(entity.Category{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCategoryField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateCategoryField...)
	return nil
}

// GetCategoryByID 根据文章分类id查询文章分类
func (r *CategoryRepo) GetCategoryByID(ctx context.Context, id uint32) (*entity.Category, error) {
	dbCategory, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbCategory, nil
}

// GetCategoryByIDs 根据文章分类id集合查询文章分类
func (r *CategoryRepo) GetCategoryByIDs(ctx context.Context, ids []uint32) (entity.CategoryList, error) {
	dbCategoryList, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return dbCategoryList, nil
}

// Save 保存文章分类
func (r *CategoryRepo) Save(ctx context.Context, category *entity.Category) (uint32, error) {
	if category.ID > 0 {
		return 0, errors.New("illegal argument category id exist")
	}
	if err := r.GenericDao.Create(ctx, category); err != nil {
		return 0, err
	}
	return category.ID, nil
}

// UpdateByID 根据ID修改文章分类
func (r *CategoryRepo) UpdateByID(ctx context.Context, category *entity.Category) error {
	if category.ID == 0 {
		return errors.New("illegal argument category exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCategoryField).Updates(category).Error
}

// DeleteByID 根据ID删除文章分类
func (r *CategoryRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchCategory 文章分类搜索
func (r *CategoryRepo) SearchCategory(ctx context.Context, req *pb.SearchCategoryReq) (entity.CategoryList, uint32, error) {
	var (
		res       []*entity.Category
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
