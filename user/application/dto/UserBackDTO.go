package dto

// UserBackDTO 后台用户
type UserBackDTO struct {
	Id            uint32        `json:"id"`            // id
	UserInfoId    uint32        `json:"userInfoId"`    // 用户信息id
	Avatar        string        `json:"avatar"`        // 头像
	Nickname      string        `json:"nickname"`      // 昵称
	RoleList      []UserRoleDTO `json:"roleList"`      // 用户角色
	LoginType     uint32        `json:"loginType"`     // 登录类型
	IpAddress     string        `json:"ipAddress"`     // 用户登录ip
	IpSource      string        `json:"ipSource"`      // ip来源
	CreateTime    uint32        `json:"createTime"`    // 创建时间
	LastLoginTime uint32        `json:"lastLoginTime"` // 最近登录时间
	IsDisable     uint32        `json:"isDisable"`     // 用户评论状态
	Status        uint32        `json:"status"`        // 状态
}
