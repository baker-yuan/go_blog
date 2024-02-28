package dto

// UserDetailDTO 用户信息
type UserDetailDTO struct {
	// 用户信息
	ID        uint32 `json:"id"`        // 用户账号id
	Email     string `json:"email"`     // 邮箱号
	Nickname  string `json:"nickname"`  // 用户昵称
	Avatar    string `json:"avatar"`    // 用户头像
	Intro     string `json:"intro"`     // 用户简介
	WebSite   string `json:"webSite"`   // 个人网站
	IsDisable uint32 `json:"isDisable"` // 是否禁用
	// 用户账号
	Username      string `json:"username"`      // 用户名
	LoginType     uint32 `json:"loginType"`     // 登录方式
	IpAddress     string `json:"ipAddress"`     // 用户登录ip
	IpSource      string `json:"ipSource"`      // ip来源
	LastLoginTime uint32 `json:"lastLoginTime"` // 最近登录时间
	// 权限
	RoleList []string `json:"roleList"` // 用户角色
	// 其他
	Browser string `json:"browser"` // 浏览器
	Os      string `json:"os"`      // 操作系统
}
