package dto

// UserRoleDTO 用户角色选项
type UserRoleDTO struct {
	Id       uint32 `json:"id"`       // 角色id
	RoleName string `json:"roleName"` // 角色名
}
