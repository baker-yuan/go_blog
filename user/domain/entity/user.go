package entity

type User struct {
	Info UserInfo
	Auth UserAuth
}

// UserAuth 用户账号
type UserAuth struct {
	ID            uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime    uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime    uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
	UserInfoId    uint32 `json:"userInfoId"`                       // 用户信息id
	Username      string `json:"username"`                         // 用户名
	Password      string `json:"password"`                         // 密码
	LoginType     uint32 `json:"loginType"`                        // 登录类型
	IpAddress     string `json:"ipAddress"`                        // 用户登录ip
	IpSource      string `json:"ipSource"`                         // ip来源
	LastLoginTime uint32 `json:"lastLoginTime"`                    // 最近登录时间
}

// UserInfo 用户信息
type UserInfo struct {
	ID         uint32 `gorm:"primaryKey"`                       // 主键
	CreateTime uint32 `json:"createTime" gorm:"autoCreateTime"` // 创建时间，使用时间戳秒数填充创建时间
	UpdateTime uint32 `json:"updateTime" gorm:"autoUpdateTime"` // 修改时间，使用时间戳秒数填充更新时间
	Email      string `json:"email"`                            // 邮箱号
	Nickname   string `json:"nickname"`                         // 用户昵称
	Avatar     string `json:"avatar"`                           // 用户头像
	Intro      string `json:"intro"`                            // 用户简介
	WebSite    string `json:"webSite"`                          // 个人网站
	IsDisable  uint32 `json:"isDisable"`                        // 是否禁言
}
