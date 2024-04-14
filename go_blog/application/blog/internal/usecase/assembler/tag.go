package assembler

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// TagEntityListToModel entity转pb
func TagEntityListToModel(tags []*entity.Tag, uidToNameMap map[uint32]string) []*pb.Tag {
	res := make([]*pb.Tag, 0)
	for _, tag := range tags {
		res = append(res, TagEntityToModel(tag, uidToNameMap))
	}
	return res
}

// TagEntityToModel entity转pb
func TagEntityToModel(tag *entity.Tag, uidToNameMap map[uint32]string) *pb.Tag {
	modelRes := &pb.Tag{
		Id:           tag.ID,
		Name:         tag.Name,
		Description:  tag.Description,
		Color:        tag.Color,
		CreateUserId: tag.CreateUserID,
		UpdateUserId: tag.UpdateUserID,
		CreateTime:   uint32(tag.CreateTime),
		UpdateTime:   uint32(tag.UpdateTime),
	}
	modelRes.CreateUser = uidToNameMap[tag.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[tag.UpdateUserID]
	return modelRes
}

// AddOrUpdateTagReqToEntity pb转entity
func AddOrUpdateTagReqToEntity(pbTag *pb.AddOrUpdateTagReq) *entity.Tag {
	entityRes := &entity.Tag{
		ID:          pbTag.Id,
		Name:        pbTag.Name,
		Description: pbTag.Description,
		Color:       pbTag.Color,
	}
	return entityRes
}
