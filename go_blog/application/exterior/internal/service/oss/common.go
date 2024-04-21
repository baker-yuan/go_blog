package oss

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/go_blog/application/exterior/config"
	oss_model "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/model/oss"
	oss_repo "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/repository/oss"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
	user_pb "github.com/baker-yuan/go-blog/go_blog/protocol/user"
)

// CommonUseCase 通用service
type CommonUseCase struct {
	cfg              *config.Config
	modeCfgRepo      oss_repo.IModeCfgRepo
	uploadRecordRepo oss_repo.IUploadRecordRepo
}

// NewCommonUseCase 创建通用service
func NewCommonUseCase(
	cfg *config.Config,
	modeCfgRepo oss_repo.IModeCfgRepo,
	uploadRecordRepo oss_repo.IUploadRecordRepo,
) *CommonUseCase {
	return &CommonUseCase{
		cfg:              cfg,
		modeCfgRepo:      modeCfgRepo,
		uploadRecordRepo: uploadRecordRepo,
	}
}

// SaveChangeLog 保存变更日志
func (c *CommonUseCase) SaveChangeLog(ctx context.Context,
	resourceID uint32, resourceType oss_pb.ResourceType,
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

// GetModeCfgByMode 根据文件模块查询文件路径
func (c *CommonUseCase) GetModeCfgByMode(ctx context.Context, mode string) (*oss_model.ModeCfg, error) {
	return c.modeCfgRepo.GetModeCfgByMode(ctx, mode)
}

// CheckUploadRecordNotExist 根据模块查询是否存在记录
func (c *CommonUseCase) CheckUploadRecordNotExist(ctx context.Context, mode string) error {
	search := map[string]*db.SearchValue{
		oss_model.UploadRecordFieldMode: {
			Value:     mode,
			Condition: db.EQ,
		},
	}
	_, pageTotal, err := c.uploadRecordRepo.SearchOssUploadRecord(ctx, &oss_pb.SearchOssUploadRecordReq{
		Search:   util.StrUtils.Obj2Json(search),
		PageNum:  1,
		PageSize: 1,
	})
	if err != nil {
		return err
	}
	if pageTotal == 0 {
		return nil
	}
	return errors.New("uploadRecord exist")
}
