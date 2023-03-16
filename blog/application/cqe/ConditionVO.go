package cqe

import (
	"github.com/baker-yuan/go-blog/public"
	"github.com/gin-gonic/gin"
)

// ConditionVO 查询条件
type ConditionVO struct {
	Current    *uint64 `json:"current" form:"current"`       // 页码
	Size       *uint64 `json:"size" form:"size"`             // 条数
	Keywords   *string `json:"keywords" form:"keywords"`     // 搜索内容
	CategoryId *uint32 `json:"categoryId" form:"categoryId"` // 分类id
	TagId      *uint32 `json:"tagId" form:"tagId"`           // 标签id
	AlbumId    *uint32 `json:"albumId" form:"albumId"`       // 相册id
	LoginType  *uint32 `json:"loginType" form:"loginType"`   // 登录类型
	Type       *uint32 `json:"type" form:"type"`             // 类型
	Status     *uint32 `json:"status" form:"status"`         // 状态
	StartTime  *uint32 `json:"startTime" form:"startTime"`   // 开始时间
	EndTime    *uint32 `json:"endTime" form:"endTime"`       // 结束时间
	IsDelete   *uint32 `json:"isDelete" form:"isDelete"`     // 是否删除
	IsReview   *uint32 `json:"isReview" form:"isReview"`     // 是否审核
}

// BindValidParam 参数绑定
func (param *ConditionVO) BindValidParam(c *gin.Context) error {
	// 参数绑定 && 参数校验
	return public.DefaultGetValidParams(c, param)
}
