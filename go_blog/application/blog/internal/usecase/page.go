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

// PageUseCase 页面管理
type PageUseCase struct {
	*CommonUseCase
	repo IPageRepo
}

// NewPageUseCase 创建页面管理业务逻辑实现
func NewPageUseCase(
	commonUseCase *CommonUseCase,
	repo IPageRepo,
) *PageUseCase {
	return &PageUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// PageDetail 页面详情
func (c *PageUseCase) PageDetail(ctx context.Context, req *pb.PageDetailReq) (*pb.Page, error) {
	page, err := c.repo.GetPageByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, page.CreateUserID, page.UpdateUserID)
	if err != nil {
		return nil, err
	}
	pbPage := assembler.PageEntityToModel(page, uidToNameMap)
	return pbPage, nil
}

// SearchPage 页面搜索
func (c *PageUseCase) SearchPage(ctx context.Context, req *pb.SearchPageReq) ([]*pb.Page, uint32, error) {
	pages, total, err := c.repo.SearchPage(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, pages.GetOperatorUserID()...)
	if err != nil {
		return nil, 0, err
	}
	return assembler.PageEntityListToModel(pages, uidToNameMap), total, nil
}

// AddOrUpdatePage 添加修改页面
func (c *PageUseCase) AddOrUpdatePage(ctx context.Context, req *pb.AddOrUpdatePageReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addPage(ctx, userID, req)
	} else {
		dbPage, err := c.repo.GetPageByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updatePage(ctx, dbPage, userID, req)
	}
}

func (c *PageUseCase) addPage(ctx context.Context, userID uint32, req *pb.AddOrUpdatePageReq) (uint32, error) {
	page := assembler.AddOrUpdatePageReqToEntity(req)
	page.CreateUserID = userID
	page.CreateTime = db.Timestamp(time.Now().Unix())
	page.UpdateUserID = userID
	page.UpdateTime = db.Timestamp(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, page)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_PAGE,
		"{}", page,
		"新增页面",
	)

	return lastInsertID, nil
}

func (c *PageUseCase) updatePage(ctx context.Context, dbPage *entity.Page, userID uint32, req *pb.AddOrUpdatePageReq) (uint32, error) {
	savePage := assembler.AddOrUpdatePageReqToEntity(req)
	savePage.CreateUserID = dbPage.CreateUserID
	savePage.CreateTime = dbPage.CreateTime
	savePage.UpdateUserID = userID
	savePage.UpdateTime = db.Timestamp(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, savePage); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_PAGE,
		dbPage, savePage,
		"全字段修改页面",
	)

	return req.GetId(), nil
}

// DeletePage 删除页面
func (c *PageUseCase) DeletePage(ctx context.Context, req *pb.DeletePageReq) error {
	page, err := c.repo.GetPageByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_PAGE,
		page, "{}",
		"删除页面",
	)

	return nil
}
