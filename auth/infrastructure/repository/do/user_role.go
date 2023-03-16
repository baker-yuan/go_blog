package do

import "github.com/baker-yuan/go-blog/all_packaged_library/common"

// UserRole 用户角色
type UserRole struct {
	common.Model
	UserId uint32 `json:"userId"` // 用户id
	RoleId uint32 `json:"roleId"` // 角色id
}

func (a UserRole) TableName() string {
	return "tb_user_role"
}
