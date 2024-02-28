// Package command 命令模式入参
package command

// NamePassLoginCMD 用户名密码登陆
type NamePassLoginCMD struct {
	Username string // 用户名
	Password string // 密码
}

// ChangePwdCMD 修改密码
type ChangePwdCMD struct {
	Username    string // 用户名
	OldPassword string // 旧密码
	NewPassword string // 新密码
}
