package usecase

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/usecase/assembler"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// TagUseCase 文章标签管理
type TagUseCase struct {
	*CommonUseCase
	repo ITagRepo
}

// NewTagUseCase 创建文章标签管理业务逻辑实现
func NewTagUseCase(
	commonUseCase *CommonUseCase,
	repo ITagRepo,
) *TagUseCase {
	return &TagUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// TagDetail 文章标签详情
func (c *TagUseCase) TagDetail(ctx context.Context, req *pb.TagDetailReq) (*pb.Tag, error) {
	tag, err := c.repo.GetTagByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, tag.CreateUserID, tag.UpdateUserID)
	if err != nil {
		return nil, err
	}
	pbTag := assembler.TagEntityToModel(tag, uidToNameMap)
	return pbTag, nil
}

// SearchTag 文章标签搜索
func (c *TagUseCase) SearchTag(ctx context.Context, req *pb.SearchTagReq) ([]*pb.Tag, uint32, error) {
	tags, total, err := c.repo.SearchTag(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, tags.GetOperatorUserID()...)
	if err != nil {
		return nil, 0, err
	}
	return assembler.TagEntityListToModel(tags, uidToNameMap), total, nil
}

// AddOrUpdateTag 添加修改文章标签
func (c *TagUseCase) AddOrUpdateTag(ctx context.Context, req *pb.AddOrUpdateTagReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addTag(ctx, userID, req)
	} else {
		dbTag, err := c.repo.GetTagByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateTag(ctx, dbTag, userID, req)
	}
}

func (c *TagUseCase) addTag(ctx context.Context, userID uint32, req *pb.AddOrUpdateTagReq) (uint32, error) {
	tag := assembler.AddOrUpdateTagReqToEntity(req)
	tag.CreateUserID = userID
	tag.CreateTime = db.Timestamp(time.Now().Unix())
	tag.UpdateUserID = userID
	tag.UpdateTime = db.Timestamp(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, tag)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_Tag,
		"{}", tag,
		"新增文章标签",
	)

	return lastInsertID, nil
}

func (c *TagUseCase) updateTag(ctx context.Context, dbTag *entity.Tag, userID uint32, req *pb.AddOrUpdateTagReq) (uint32, error) {
	saveTag := assembler.AddOrUpdateTagReqToEntity(req)
	saveTag.CreateUserID = dbTag.CreateUserID
	saveTag.CreateTime = dbTag.CreateTime
	saveTag.UpdateUserID = userID
	saveTag.UpdateTime = db.Timestamp(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveTag); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_Tag,
		dbTag, saveTag,
		"全字段修改文章标签",
	)

	return req.GetId(), nil
}

// DeleteTag 删除文章标签
func (c *TagUseCase) DeleteTag(ctx context.Context, req *pb.DeleteTagReq) error {
	tag, err := c.repo.GetTagByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_Tag,
		tag, "{}",
		"删除文章标签",
	)

	return nil
}
