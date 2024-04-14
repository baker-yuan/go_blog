package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
	"gorm.io/gorm"
)

// TagRepo 文章标签
type TagRepo struct {
	*db.GenericDao[entity.Tag, uint32]
}

// NewTagRepo 创建
func NewTagRepo(gormDB *gorm.DB) *TagRepo {
	return &TagRepo{
		GenericDao: db.NewGenericDao[entity.Tag, uint32](gormDB),
	}
}

func init() {
	registerInitField(initTagField)
}

var (
	// 全字段修改Tag那些字段不修改
	notUpdateTagField = []string{
		entity.TagFieldCreateUserID,
		entity.TagFieldCreateTime,
	}
	updateTagField []string
)

// InitTagField 全字段修改
func initTagField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(entity.Tag{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateTagField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateTagField...)
	return nil
}

// GetTagByID 根据文章标签id查询文章标签
func (r *TagRepo) GetTagByID(ctx context.Context, id uint32) (*entity.Tag, error) {
	dbTag, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbTag, nil
}

// GetTagByIDs 根据文章标签id集合查询文章标签
func (r *TagRepo) GetTagByIDs(ctx context.Context, ids []uint32) (entity.Tags, error) {
	dbTags, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return dbTags, nil
}

// Save 保存文章标签
func (r *TagRepo) Save(ctx context.Context, tag *entity.Tag) (uint32, error) {
	if tag.ID > 0 {
		return 0, errors.New("illegal argument tag id exist")
	}
	if err := r.GenericDao.Create(ctx, tag); err != nil {
		return 0, err
	}
	return tag.ID, nil
}

// UpdateByID 根据ID修改文章标签
func (r *TagRepo) UpdateByID(ctx context.Context, tag *entity.Tag) error {
	if tag.ID == 0 {
		return errors.New("illegal argument tag exist")
	}
	return r.GenericDao.GetDB(ctx).Select(updateTagField).Updates(tag).Error
}

// DeleteByID 根据ID删除文章标签
func (r *TagRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchTag 文章标签搜索
func (r *TagRepo) SearchTag(ctx context.Context, req *pb.SearchTagReq) (entity.Tags, uint32, error) {
	var (
		res       []*entity.Tag
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.GetDB(ctx),
		func(search map[string]*db.SearchValue) {

		},
	)
	tx = tx.Order(fmt.Sprintf("%s desc", entity.TagFieldID))
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
