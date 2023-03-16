package dto

// UserDetailDTO 用户信息
type UserDetailDTO struct {
	// UserDetailDTO implements UserDetails.go
	// UserDetails
	ID             uint32        `json:"id"`             // 用户账号id
	UserInfoId     uint32        `json:"userInfoId"`     // 用户信息id
	Email          string        `json:"email"`          // 邮箱号
	LoginType      uint32        `json:"loginType"`      // 登录方式
	Username       string        `json:"username"`       // 用户名
	Password       string        `json:"password"`       // 密码
	RoleList       []string      `json:"roleList"`       // 用户角色
	Nickname       string        `json:"nickname"`       // 用户昵称
	Avatar         string        `json:"avatar"`         // 用户头像
	Intro          string        `json:"intro"`          // 用户简介
	WebSite        string        `json:"webSite"`        // 个人网站
	ArticleLikeSet []interface{} `json:"articleLikeSet"` // 点赞文章集合
	CommentLikeSet []interface{} `json:"commentLikeSet"` // 点赞评论集合
	TalkLikeSet    []interface{} `json:"talkLikeSet"`    // 点赞说说集合
	IpAddress      string        `json:"ipAddress"`      // 用户登录ip
	IpSource       string        `json:"ipSource"`       // ip来源
	IsDisable      uint32        `json:"isDisable"`      // 是否禁用
	Browser        string        `json:"browser"`        // 浏览器
	Os             string        `json:"os"`             // 操作系统
	LastLoginTime  uint32        `json:"lastLoginTime"`  // 最近登录时间
}
