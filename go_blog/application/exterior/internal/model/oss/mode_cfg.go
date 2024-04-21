package oss

import (
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
)

// ModeCfg OSS模块配置
type ModeCfg struct {
	ID   uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	Mode string `gorm:"uniqueIndex:uk_mode;column:mode;type:varchar(100);not null;default:'';comment:模块"`
	Path string `gorm:"column:path;type:varchar(100);not null;default:'';comment:路径"`
	Desc string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 FilePath 结构体对应的数据库表名
func (ModeCfg) TableName() string {
	return ModeCfgTbName
}

var (
	ModeCfgTbName            = "oss_mode_cfg"
	ModeCfgFieldID           = "id"
	ModeCfgFieldMode         = "mode"
	ModeCfgFieldPath         = "path"
	ModeCfgFieldDesc         = "desc"
	ModeCfgFieldCreateUserID = "create_user_id"
	ModeCfgFieldUpdateUserID = "update_user_id"
	ModeCfgFieldCreateTime   = "create_time"
	ModeCfgFieldUpdateTime   = "update_time"
)

type ModeCfgList []*ModeCfg

// GetOperatorUserID 获取操作人ID
func (l ModeCfgList) GetOperatorUserID() []uint32 {
	res := make([]uint32, 0)
	for _, item := range l {
		res = append(res, item.CreateUserID)
		res = append(res, item.UpdateUserID)
	}
	return util.NewSliceUtils[uint32]().SliceUnique(res)
}
