package entity

import (
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// FriendLink 友链表
type FriendLink struct {
	ID          uint32              `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	LinkName    string              `gorm:"column:link_name;type:varchar(20);not null;default:'';comment:链接名"`
	LinkAvatar  string              `gorm:"column:link_avatar;type:varchar(255);not null;default:'';comment:链接头像"`
	LinkAddress string              `gorm:"column:link_address;type:varchar(50);not null;default:'';comment:链接地址"`
	LinkIntro   string              `gorm:"column:link_intro;type:varchar(100);not null;default:'';comment:链接介绍"`
	Status      pb.FriendLinkStatus `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:友链状态 1-已发布 2-以下线"`
	Sort        uint32              `gorm:"column:sort;type:int unsigned;not null;default:0;comment:友链排序"`
	// 逻辑删除
	IsDeleted BoolBit `gorm:"column:is_deleted;type:bit(1);not null;default:b'0';comment:是否删除"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 FriendLink 结构体对应的数据库表名
func (FriendLink) TableName() string {
	return FriendLinkTbName
}

var (
	FriendLinkTbName            = "blog_friend_link"
	FriendLinkFieldID           = "id"
	FriendLinkFieldLinkName     = "link_name"
	FriendLinkFieldLinkAvatar   = "link_avatar"
	FriendLinkFieldLinkAddress  = "link_address"
	FriendLinkFieldLinkIntro    = "link_intro"
	FriendLinkFieldStatus       = "status"
	FriendLinkFieldSort         = "sort"
	FriendLinkFieldIsDeleted    = "is_deleted"
	FriendLinkFieldCreateUserId = "create_user_id"
	FriendLinkFieldUpdateUserId = "update_user_id"
	FriendLinkFieldCreateTime   = "create_time"
	FriendLinkFieldUpdateTime   = "update_time"
)

type FriendLinks []*FriendLink

// GetOperatorUserID 获取操作人ID
func (l FriendLinks) GetOperatorUserID() []uint32 {
	res := make([]uint32, 0)
	for _, item := range l {
		res = append(res, item.CreateUserID)
		res = append(res, item.UpdateUserID)
	}
	return util.NewSliceUtils[uint32]().SliceUnique(res)
}
