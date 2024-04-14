package entity

import (
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
)

// Tag 文章标签表
type Tag struct {
	ID          uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	Name        string `gorm:"uniqueIndex:uk_name;column:name;type:varchar(50);not null;default:'';comment:标签名"`
	Description string `gorm:"column:description;type:varchar(100);not null;default:'';comment:标签描述"`
	Color       string `gorm:"column:color;type:varchar(64);not null;default:'#000000';comment:标签颜色(十六进制)"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Tag 结构体对应的数据库表名
func (Tag) TableName() string {
	return TagTbName
}

var (
	TagTbName            = "blog_tag"
	TagFieldID           = "id"
	TagFieldName         = "name"
	TagFieldDescription  = "description"
	TagFieldColor        = "color"
	TagFieldCreateUserID = "create_user_id"
	TagFieldUpdateUserID = "update_user_id"
	TagFieldCreateTime   = "create_time"
	TagFieldUpdateTime   = "update_time"
)

type Tags []*Tag

// GetOperatorUserID 获取操作人ID
func (t Tags) GetOperatorUserID() []uint32 {
	res := make([]uint32, 0)
	for _, item := range t {
		res = append(res, item.CreateUserID)
		res = append(res, item.UpdateUserID)
	}
	return util.NewSliceUtils[uint32]().SliceUnique(res)
}
