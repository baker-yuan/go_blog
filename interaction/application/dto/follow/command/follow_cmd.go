// Package command 命令模式入参
package command

// AddFollowCMD 新增关注关系
type AddFollowCMD struct {
	UID       uint32 // 关注的人
	FollowUID uint32 // 被关注的人
}
