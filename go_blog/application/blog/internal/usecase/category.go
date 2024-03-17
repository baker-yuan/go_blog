package usecase

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/application/blog/internal/usecase/assembler"
	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
)

// CategoryUseCase 文章分类管理
type CategoryUseCase struct {
	*CommonUseCase
	repo ICategoryRepo
}

// NewCategoryUseCase 创建文章分类管理业务逻辑实现
func NewCategoryUseCase(
	commonUseCase *CommonUseCase,
	repo ICategoryRepo,
) *CategoryUseCase {
	return &CategoryUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// CategoryDetail 文章分类详情
func (c *CategoryUseCase) CategoryDetail(ctx context.Context, req *pb.CategoryDetailReq) (*pb.Category, error) {
	category, err := c.repo.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	pbCategory := assembler.CategoryEntityToModel(category)
	return pbCategory, nil
}

// SearchCategory 文章分类搜索
func (c *CategoryUseCase) SearchCategory(ctx context.Context, req *pb.SearchCategoryReq) ([]*pb.Category, uint32, error) {
	categoryList, total, err := c.repo.SearchCategory(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.Category, 0)
	for _, category := range categoryList {
		data = append(data, assembler.CategoryEntityToModel(category))
	}
	return data, total, nil
}

// AddOrUpdateCategory 添加修改文章分类
func (c *CategoryUseCase) AddOrUpdateCategory(ctx context.Context, req *pb.AddOrUpdateCategoryReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addCategory(ctx, userID, req)
	} else {
		dbCategory, err := c.repo.GetCategoryByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateCategory(ctx, dbCategory, userID, req)
	}
}

func (c *CategoryUseCase) addCategory(ctx context.Context, userID uint32, req *pb.AddOrUpdateCategoryReq) (uint32, error) {
	category := assembler.AddOrUpdateCategoryReqToEntity(req)
	category.CreateUser = userID
	category.CreateTime = entity.Timestamp(uint32(time.Now().Unix()))
	category.UpdateUser = userID
	category.UpdateTime = entity.Timestamp(uint32(time.Now().Unix()))

	lastInsertID, err := c.repo.Save(ctx, category)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_CATEGORY,
		"{}", category,
		"新增文章分类",
	)

	return lastInsertID, nil
}

func (c *CategoryUseCase) updateCategory(ctx context.Context, dbCategory *entity.Category, userID uint32, req *pb.AddOrUpdateCategoryReq) (uint32, error) {
	saveCategory := assembler.AddOrUpdateCategoryReqToEntity(req)
	saveCategory.CreateUser = dbCategory.CreateUser
	saveCategory.CreateTime = dbCategory.CreateTime
	saveCategory.UpdateUser = userID
	saveCategory.UpdateTime = entity.Timestamp(uint32(time.Now().Unix()))

	if err := c.repo.UpdateByID(ctx, saveCategory); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_CATEGORY,
		dbCategory, saveCategory,
		"全字段修改文章分类",
	)

	return req.GetId(), nil
}

// DeleteCategory 删除文章分类
func (c *CategoryUseCase) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryReq) error {
	category, err := c.repo.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_CATEGORY,
		category, "{}",
		"删除文章分类",
	)

	return nil
}
