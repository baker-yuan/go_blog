package cqe

import (
	"github.com/baker-yuan/go-blog/public"
	"github.com/gin-gonic/gin"
)

// ArticleSearch 文章搜索
type ArticleSearch struct {
	ID             *uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime     *uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime     *uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
	UserId         *uint32 `json:"userId"`                           // 作者
	CategoryId     *uint32 `json:"categoryId"`                       // 文章分类
	ArticleCover   *string `json:"articleCover"`                     // 文章缩略图
	ArticleTitle   *string `json:"articleTitle"`                     // 标题
	ArticleContent *string `json:"articleContent"`                   // 内容
	Type           *uint32 `json:"type"`                             // 文章类型
	OriginalUrl    *string `json:"originalUrl"`                      // 原文链接
	IsTop          *uint32 `json:"isTop"`                            // 是否置顶
	IsDelete       *uint32 `json:"isDelete"`                         // 是否删除
	Status         *uint32 `json:"status"`                           // 文章状态 1-公开 2-私密 3-评论可见
}

// ArticleTopVO 文章置顶信息
type ArticleTopVO struct {
	ID    uint32 `json:"id"`    // id
	IsTop uint32 `json:"isTop"` // 置顶状态
}

// BindValidParam 参数绑定
func (param *ArticleTopVO) BindValidParam(c *gin.Context) error {
	// 参数绑定 && 参数校验
	return public.DefaultGetValidParams(c, param)
}
