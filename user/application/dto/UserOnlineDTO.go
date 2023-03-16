package dto

// UserOnlineDTO 在线用户
type UserOnlineDTO struct {
	UserInfoId    uint32 `json:"userInfoId"`    // 用户信息id
	Nickname      string `json:"nickname"`      // 用户昵称
	Avatar        string `json:"avatar"`        // 用户头像
	IpAddress     string `json:"ipAddress"`     // 用户登录ip
	IpSource      string `json:"ipSource"`      // ip来源
	Browser       string `json:"browser"`       // 浏览器
	Os            string `json:"os"`            // 操作系统
	LastLoginTime uint32 `json:"lastLoginTime"` // 最近登录时间
}
