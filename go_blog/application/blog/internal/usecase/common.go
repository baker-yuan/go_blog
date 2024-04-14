package usecase

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
	user_pb "github.com/baker-yuan/go-blog/go_blog/protocol/user"
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

// UidToName 根据用户ID获取用户名称
func (c *CommonUseCase) UidToName(ctx context.Context, userIDs ...uint32) (map[uint32]string, error) {
	userIDs = util.NewSliceUtils[uint32]().SliceFilter(userIDs, func(u uint32) bool {
		return u == 0
	})
	proxy := user_pb.NewUserApiClientProxy()
	rsp, err := proxy.UidToName(ctx, &user_pb.UidToNameReq{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}
	return rsp.GetUidToNameMap(), nil
}
