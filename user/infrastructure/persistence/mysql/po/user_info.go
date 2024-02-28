package po

import "github.com/baker-yuan/go-blog/all_packaged_library/do"

// UserInfoPO 用户信息
type UserInfoPO struct {
	do.Time
	UID       uint32 `gorm:"primaryKey"` // 主键
	Email     string `json:"email"`      // 邮箱号
	Nickname  string `json:"nickname"`   // 用户昵称
	Avatar    string `json:"avatar"`     // 用户头像
	Intro     string `json:"intro"`      // 用户简介
	WebSite   string `json:"webSite"`    // 个人网站
	IsDisable uint32 `json:"isDisable"`  // 是否禁言
}

func (a UserInfoPO) TableName() string {
	return "tb_user_info"
}
