package biz

import (
	"context"

	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// CommonUseCase 通用service
type CommonUseCase struct {
}

// NewCommonUseCase 创建通用service
func NewCommonUseCase() *CommonUseCase {
	return &CommonUseCase{}
}

// SaveChangeLog 保存变更日志
func (c *CommonUseCase) SaveChangeLog(ctx context.Context,
	resourceID uint32, resourceType pb.ResourceType,
	changeBefore interface{}, changeAfter interface{},
	notes string) {
	//before := c.convertToJSONString(changeBefore)
	//after := c.convertToJSONString(changeAfter)
	//c.resourceChangeLogRepo.Save(ctx, resourceID, resourceType, before, after, notes)
}

// convertToJSONString 将任意类型的数据转换为JSON字符串
func (c *CommonUseCase) convertToJSONString(data interface{}) string {
	//if data == nil {
	//	return "{}"
	//}
	//switch data.(type) {
	//case string:
	//	return data.(string)
	//default:
	//	return util.StrUtils.Obj2Json(data)
	//}

	return ""
}
