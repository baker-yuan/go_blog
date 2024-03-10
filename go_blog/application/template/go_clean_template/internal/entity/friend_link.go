package entity

// FriendLink 友链表
type FriendLink struct {
	ID          uint32 `gorm:"column:id;type:int unsigned;primary_key;auto_increment;comment:链接ID"`
	LinkName    string `gorm:"column:link_name;type:varchar(20);not null;default:'';comment:链接名"`
	LinkAvatar  string `gorm:"column:link_avatar;type:varchar(255);not null;default:'';comment:链接头像"`
	LinkAddress string `gorm:"column:link_address;type:varchar(50);not null;default:'';comment:链接地址"`
	LinkIntro   string `gorm:"column:link_intro;type:varchar(100);not null;default:'';comment:链接介绍"`
	Status      uint8  `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:友链状态 1-已发布 2-以下线"`
	Sort        uint32 `gorm:"column:sort;type:int unsigned;not null;default:0;comment:友链排序"`

	//IsDeleted bool `gorm:"column:is_deleted;type:bit(1);not null;default:0;comment:是否删除"`

	// 公共字段
	SoftDelete
	Operator
	BaseTime
}

func (r FriendLink) TableName() string {
	return "blog_friend_link"
}

type FriendLinks []*FriendLink
