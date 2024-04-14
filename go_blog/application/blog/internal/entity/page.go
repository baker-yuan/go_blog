package entity

import (
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
)

// Page 页面表
type Page struct {
	ID        uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:页面id"`
	PageName  string `gorm:"column:page_name;type:varchar(10);not null;default:'';comment:页面名"`
	PageLabel string `gorm:"column:page_label;type:varchar(20);not null;default:'';comment:页面标签"`
	PageCover string `gorm:"column:page_cover;type:varchar(255);not null;default:'';comment:页面封面"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Page 结构体对应的数据库表名
func (Page) TableName() string {
	return PageTbName
}

var (
	PageTbName          = "blog_page"
	PageFieldID         = "id"
	PageFieldPageName   = "page_name"
	PageFieldPageLabel  = "page_label"
	PageFieldPageCover  = "page_cover"
	PageFieldCreateTime = "create_time"
	PageFieldUpdateTime = "update_time"
)

type Pages []*Page

// GetOperatorUserID 获取操作人ID
func (l Pages) GetOperatorUserID() []uint32 {
	res := make([]uint32, 0)
	for _, item := range l {
		res = append(res, item.CreateUserID)
		res = append(res, item.UpdateUserID)
	}
	return util.NewSliceUtils[uint32]().SliceUnique(res)
}
