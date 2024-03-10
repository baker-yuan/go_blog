package entity

import pb "github.com/baker-yuan/go-blog/protocol/blog"

// FriendLink 友链表
type FriendLink struct {
	ID          uint32              `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	LinkName    string              `gorm:"column:link_name;type:varchar(20);not null;default:'';comment:链接名"`
	LinkAvatar  string              `gorm:"column:link_avatar;type:varchar(255);not null;default:'';comment:链接头像"`
	LinkAddress string              `gorm:"column:link_address;type:varchar(50);not null;default:'';comment:链接地址"`
	LinkIntro   string              `gorm:"column:link_intro;type:varchar(100);not null;default:'';comment:链接介绍"`
	Status      pb.FriendLinkStatus `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:友链状态 1-已发布 2-以下线"`
	Sort        uint32              `gorm:"column:sort;type:int unsigned;not null;default:0;comment:友链排序"`
	IsDeleted   BoolBit             `gorm:"column:is_deleted;type:bit(1);not null;default:b'0';comment:是否删除"`
	CreateTime  Timestamp           `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime  Timestamp           `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 FriendLink 结构体对应的数据库表名
func (FriendLink) TableName() string {
	return "blog_friend_link"
}
