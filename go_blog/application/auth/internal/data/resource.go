package data

import (
	"github.com/baker-yuan/go-blog/common/db"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// Resource 接口资源表
type Resource struct {
	// 基本信息
	ID       uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	ParentID uint32 `gorm:"column:parent_id;type:int unsigned;not null;default:0;comment:父接口ID"`
	Name     string `gorm:"column:name;type:varchar(150);not null;default:'';comment:接口名称"`
	Describe string `gorm:"column:describe;type:varchar(255);not null;default:'';comment:接口描述"`
	//
	ResourceType pb.ResourceResourceType `gorm:"column:resource_type;type:tinyint unsigned;not null;default:0;comment:资源类型 1-目录 2-接口"`
	Status       pb.ResourceStatus       `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:接口状态 0-未发布 1-已发布 2-以下线"`
	// 路径定位
	URL        string `gorm:"column:url;type:varchar(255);not null;default:'';comment:接口URL（包含路径变量）"`
	HTTPMethod string `gorm:"column:http_method;type:varchar(64);not null;default:'';comment:请求方式 1-get 2-post 3-put 4-patch 5-delete"`
	// 权限校验
	IsNeedLogin      uint8 `gorm:"column:is_need_login;type:tinyint;not null;default:1;comment:是否需要登录校验 0-不需要 1-需要"`
	IsNeedPermission uint8 `gorm:"column:is_need_permission;type:tinyint;not null;default:1;comment:是否需要权限校验 0-不需要 1-需要"`
	// 下游服务信息
	Service string `gorm:"column:service;type:varchar(30);not null;default:'';comment:API服务名"`
	Method  string `gorm:"column:method;type:varchar(255);not null;default:'';comment:API接口名"`
	// 公共字段
	CreateUser uint32       `gorm:"column:create_user;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUser uint32       `gorm:"column:update_user;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 AuthResource 结构体对应的数据库表名
func (Resource) TableName() string {
	return "baker_auth_resource"
}
