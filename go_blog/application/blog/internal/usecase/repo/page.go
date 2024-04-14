package repo

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
	"gorm.io/gorm"
)

// PageRepo 页面
type PageRepo struct {
	*db.GenericDao[entity.Page, uint32]
}

// NewPageRepo 创建
func NewPageRepo(gormDB *gorm.DB) *PageRepo {
	return &PageRepo{
		GenericDao: db.NewGenericDao[entity.Page, uint32](gormDB),
	}
}

func init() {
	registerInitField(initPageField)
}

var (
	// 全字段修改Page那些字段不修改
	notUpdatePageField = []string{
		"created_at",
		"create_time",
	}
	updatePageField []string
)

// InitPageField 全字段修改
func initPageField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(entity.Page{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updatePageField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdatePageField...)
	return nil
}

// GetPageByID 根据页面id查询页面
func (r *PageRepo) GetPageByID(ctx context.Context, id uint32) (*entity.Page, error) {
	dbPage, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbPage, nil
}

// GetPageByIDs 根据页面id集合查询页面
func (r *PageRepo) GetPageByIDs(ctx context.Context, ids []uint32) (entity.Pages, error) {
	dbPages, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return dbPages, nil
}

// Save 保存页面
func (r *PageRepo) Save(ctx context.Context, page *entity.Page) (uint32, error) {
	if page.ID > 0 {
		return 0, errors.New("illegal argument page id exist")
	}
	if err := r.GenericDao.Create(ctx, page); err != nil {
		return 0, err
	}
	return page.ID, nil
}

// UpdateByID 根据ID修改页面
func (r *PageRepo) UpdateByID(ctx context.Context, page *entity.Page) error {
	if page.ID == 0 {
		return errors.New("illegal argument page exist")
	}
	return r.GenericDao.GetDB(ctx).Select(updatePageField).Updates(page).Error
}

// DeleteByID 根据ID删除页面
func (r *PageRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchPage 页面搜索
func (r *PageRepo) SearchPage(ctx context.Context, req *pb.SearchPageReq) (entity.Pages, uint32, error) {
	var (
		res       []*entity.Page
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.GetDB(ctx),
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
