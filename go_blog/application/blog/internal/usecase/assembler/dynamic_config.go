package assembler

import (
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// DynamicConfigEntityListToModel entity转pb
func DynamicConfigEntityListToModel(dynamicConfigList []*entity.DynamicConfig, uidToNameMap map[uint32]string) []*pb.DynamicConfig {
	res := make([]*pb.DynamicConfig, 0)
	for _, dynamicConfig := range dynamicConfigList {
		res = append(res, DynamicConfigEntityToModel(dynamicConfig, uidToNameMap))
	}
	return res
}

// DynamicConfigEntityToModel entity转pb
func DynamicConfigEntityToModel(dynamicConfig *entity.DynamicConfig, uidToNameMap map[uint32]string) *pb.DynamicConfig {
	modelRes := &pb.DynamicConfig{
		Id:           dynamicConfig.ID,
		BizType:      string(dynamicConfig.BizType),
		BizDesc:      dynamicConfig.BizDesc,
		Content:      dynamicConfig.Content,
		JsonSchema:   dynamicConfig.JsonSchema,
		CreateUserId: dynamicConfig.CreateUserID,
		UpdateUserId: dynamicConfig.UpdateUserID,
		CreateTime:   uint32(dynamicConfig.CreateTime),
		UpdateTime:   uint32(dynamicConfig.UpdateTime),
	}
	modelRes.CreateUser = uidToNameMap[dynamicConfig.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[dynamicConfig.UpdateUserID]
	return modelRes
}

// AddOrUpdateDynamicConfigReqToEntity pb转entity
func AddOrUpdateDynamicConfigReqToEntity(pbDynamicConfig *pb.AddOrUpdateDynamicConfigReq) *entity.DynamicConfig {
	entityRes := &entity.DynamicConfig{
		ID:         pbDynamicConfig.Id,
		BizType:    entity.BizType(pbDynamicConfig.BizType),
		BizDesc:    pbDynamicConfig.BizDesc,
		Content:    pbDynamicConfig.Content,
		JsonSchema: pbDynamicConfig.JsonSchema,
	}
	return entityRes
}
