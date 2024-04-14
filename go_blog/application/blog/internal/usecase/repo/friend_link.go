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

// FriendLinkRepo 友链表
type FriendLinkRepo struct {
	*db.GenericDao[entity.FriendLink, uint32]
}

// NewFriendLinkRepo 创建
func NewFriendLinkRepo(gormDB *gorm.DB) *FriendLinkRepo {
	return &FriendLinkRepo{
		GenericDao: db.NewGenericDao[entity.FriendLink, uint32](gormDB),
	}
}

func init() {
	registerInitField(initFriendLinkField)
}

var (
	// 全字段修改FriendLink那些字段不修改
	notUpdateFriendLinkField = []string{
		"created_at",
		"create_time",
	}
	updateFriendLinkField []string
)

// InitFriendLinkField 全字段修改
func initFriendLinkField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(entity.FriendLink{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateFriendLinkField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateFriendLinkField...)
	return nil
}

// GetFriendLinkByID 根据友链表id查询友链表
func (r *FriendLinkRepo) GetFriendLinkByID(ctx context.Context, id uint32) (*entity.FriendLink, error) {
	dbFriendLink, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbFriendLink, nil
}

// GetFriendLinkByIDs 根据友链表id集合查询友链表
func (r *FriendLinkRepo) GetFriendLinkByIDs(ctx context.Context, ids []uint32) (entity.FriendLinks, error) {
	dbFriendLinks, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return dbFriendLinks, nil
}

// Save 保存友链表
func (r *FriendLinkRepo) Save(ctx context.Context, friendLink *entity.FriendLink) (uint32, error) {
	if friendLink.ID > 0 {
		return 0, errors.New("illegal argument friendLink id exist")
	}
	if err := r.GenericDao.Create(ctx, friendLink); err != nil {
		return 0, err
	}
	return friendLink.ID, nil
}

// UpdateByID 根据ID修改友链表
func (r *FriendLinkRepo) UpdateByID(ctx context.Context, friendLink *entity.FriendLink) error {
	if friendLink.ID == 0 {
		return errors.New("illegal argument friendLink exist")
	}
	return r.GenericDao.GetDB(ctx).Select(updateFriendLinkField).Updates(friendLink).Error
}

// DeleteByID 根据ID删除友链表
func (r *FriendLinkRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchFriendLink 友链表搜索
func (r *FriendLinkRepo) SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) (entity.FriendLinks, uint32, error) {
	var (
		res       []*entity.FriendLink
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
