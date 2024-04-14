package assembler

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// CategoryEntityListToModel entity转pb
func CategoryEntityListToModel(categoryList []*entity.Category, uidToNameMap map[uint32]string) []*pb.Category {
	res := make([]*pb.Category, 0)
	for _, category := range categoryList {
		res = append(res, CategoryEntityToModel(category, uidToNameMap))
	}
	return res

}

// CategoryEntityToModel entity转pb
func CategoryEntityToModel(category *entity.Category, uidToNameMap map[uint32]string) *pb.Category {
	modelRes := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Sort:        category.Sort,
		Available:   category.Available,
		//
		CreateUserId: category.CreateUserID,
		UpdateUserId: category.UpdateUserID,
		CreateTime:   uint32(category.CreateTime),
		UpdateTime:   uint32(category.UpdateTime),
	}
	modelRes.CreateUser = uidToNameMap[category.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[category.UpdateUserID]
	return modelRes
}

// AddOrUpdateCategoryReqToEntity pb转entity
func AddOrUpdateCategoryReqToEntity(pbCategory *pb.AddOrUpdateCategoryReq) *entity.Category {
	entityRes := &entity.Category{
		ID:          pbCategory.Id,
		Name:        pbCategory.Name,
		Description: pbCategory.Description,
		Sort:        pbCategory.Sort,
		Available:   pbCategory.Available,
	}
	return entityRes
}
