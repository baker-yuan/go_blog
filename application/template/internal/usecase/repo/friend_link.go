package repo

import (
	"context"

	pb "github.com/baker-yuan/go-blog/protocol/template"
	"github.com/baker-yuan/go-blog/template/internal/entity"
	"gorm.io/gorm"
)

// FriendLinkRepo 友链Repo
type FriendLinkRepo struct {
	gormDB *gorm.DB
}

// NewFriendLinkRepo 创建友链Repo
func NewFriendLinkRepo(gormDB *gorm.DB) *FriendLinkRepo {
	return &FriendLinkRepo{
		gormDB: gormDB,
	}
}

//func (r *FriendLinkRepo) getStatement() *orm.Statement {
//	return orm.NewDbStatement().SetTableName(entity.FriendLinkTbName)
//}

// GetFriendLinkByID 根据友链id查询友链
func (r *FriendLinkRepo) GetFriendLinkByID(ctx context.Context, id int) (*entity.FriendLink, error) {
	dest := &entity.FriendLink{}
	//statement := r.getStatement().AndEqual(entity.FriendLinkFieldID, id)
	//if err := r.client.FindOne(ctx, statement, dest); err != nil {
	//	return nil, err
	//}
	//if dest.ID == 0 {
	//	return nil, retcode.BuildErrorFmtMsg(retcode.RetResourceNotExist)
	//}
	return dest, nil
}

// GetFriendLinkByIDs 根据友链id集合查询友链
func (r *FriendLinkRepo) GetFriendLinkByIDs(ctx context.Context, ids []int) (entity.FriendLinks, error) {
	dest := make([]*entity.FriendLink, 0)
	//if len(ids) == 0 {
	//	return dest, nil
	//}
	//var where = orm.WhereCond{entity.FriendLinkFieldID: ids}
	//var statement = r.getStatement().Where(where)
	//if err := r.client.FindAll(ctx, statement, &dest); err != nil {
	//	return nil, err
	//}
	return dest, nil
}

// SearchFriendLink 友链搜索
func (r *FriendLinkRepo) SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) (entity.FriendLinks, uint32, error) {

	//var where = orm.WhereCond{}
	//
	//if len(req.Id) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldId)] = "%" + req.GetId() + "%"
	//}
	//if len(req.LinkName) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldLinkName)] = "%" + req.GetLinkName() + "%"
	//}
	//if len(req.LinkAvatar) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldLinkAvatar)] = "%" + req.GetLinkAvatar() + "%"
	//}
	//if len(req.LinkAddress) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldLinkAddress)] = "%" + req.GetLinkAddress() + "%"
	//}
	//if len(req.LinkIntro) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldLinkIntro)] = "%" + req.GetLinkIntro() + "%"
	//}
	//if len(req.Status) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldStatus)] = "%" + req.GetStatus() + "%"
	//}
	//if len(req.Sort) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldSort)] = "%" + req.GetSort() + "%"
	//}
	//if len(req.IsDeleted) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldIsDeleted)] = "%" + req.GetIsDeleted() + "%"
	//}
	//if len(req.CreateTime) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldCreateTime)] = "%" + req.GetCreateTime() + "%"
	//}
	//if len(req.UpdateTime) != nil {
	//	where[util.StrUtils.OrmLike(entity.FriendLinkFieldUpdateTime)] = "%" + req.GetUpdateTime() + "%"
	//}
	//
	//commonStatement := func() *orm.Statement {
	//	return r.getStatement().
	//		Where(where).
	//		Order(entity.FriendLinkFieldID, true)
	//}

	// 查询
	res := make([]*entity.FriendLink, 0)
	//searchStatement := commonStatement().LimitOffset(orm.QuickPaginate(req.GetPageSize(), req.GetPageNum()))
	//if err := r.client.FindAll(ctx, searchStatement, &res); err != nil {
	//	return nil, 0, err
	//}
	//// 分页
	//pageStatement := commonStatement()
	//total, err := r.client.Count(ctx, pageStatement)
	//if err != nil {
	//	return nil, 0, err
	//}

	if err := r.gormDB.WithContext(ctx).Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, uint32(0), nil
}

// Save 保存友链
func (r *FriendLinkRepo) Save(ctx context.Context, friendLink *entity.FriendLink) (uint32, error) {
	//statement := r.getStatement().InsertStruct(friendLink)
	//lastInsertId, err := r.client.Insert(ctx, statement)
	//if err != nil {
	//	return 0, err
	//}
	return uint32(0), nil
}

// UpdateByID 根据ID修改友链
func (r *FriendLinkRepo) UpdateByID(ctx context.Context, friendLink *entity.FriendLink) error {
	//if friendLink.ID == 0 {
	//	return retcode.BuildErrorFmtMsg(retcode.IllegalArgument)
	//}
	//statement := r.getStatement().
	//	AndEqual(entity.FriendLinkFieldID, friendLink.ID).
	//	UpdateStruct(friendLink)
	//_, err := r.client.Update(ctx, statement)
	//if err != nil {
	//	return err
	//}
	return nil
}

// DeleteByID 根据ID删除友链
func (r *FriendLinkRepo) DeleteByID(ctx context.Context, id int) error {
	//statement := r.getStatement().AndEqual(entity.FriendLinkFieldID, id)
	//_, err := r.client.Delete(ctx, statement)
	//if err != nil {
	//	return err
	//}
	return nil
}
