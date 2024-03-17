package assembler

import (
	"github.com/baker-yuan/go-blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
)

// CategoryEntityToModel entity转pb
func CategoryEntityToModel(category *entity.Category) *pb.Category {
	modelRes := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Sort:        uint32(category.Sort),
		Available:   category.Available,
		//
		CreateTime: uint32(category.CreateTime),
		UpdateTime: uint32(category.UpdateTime),
	}
	return modelRes
}

// AddOrUpdateCategoryReqToEntity pb转entity
func AddOrUpdateCategoryReqToEntity(pbCategory *pb.AddOrUpdateCategoryReq) *entity.Category {
	entityRes := &entity.Category{}
	return entityRes
}
