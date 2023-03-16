package entity

import "github.com/baker-yuan/go-blog/all_packaged_library/common"

// UserAuthDO 用户账号
type UserAuthDO struct {
	common.Model
	UserInfoId    uint32 `json:"userInfoId"`    // 用户信息id
	Username      string `json:"username"`      // 用户名
	Password      string `json:"password"`      // 密码
	LoginType     uint32 `json:"loginType"`     // 登录类型
	IpAddress     string `json:"ipAddress"`     // 用户登录ip
	IpSource      string `json:"ipSource"`      // ip来源
	LastLoginTime uint32 `json:"lastLoginTime"` // 最近登录时间
}

func (a UserAuthDO) TableName() string {
	return "tb_user_auth"
}
