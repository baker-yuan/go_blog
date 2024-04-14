package data

import "github.com/baker-yuan/go-blog/go_blog/application/auth/internal/biz"

// RoleAuthorityEntityListToPO entity转po
func RoleAuthorityEntityListToPO(roleAuthorityList []*biz.RoleAuthority) []*RoleAuthority {
	res := make([]*RoleAuthority, 0)
	for _, roleAuthority := range roleAuthorityList {
		res = append(res, RoleAuthorityEntityToPO(roleAuthority))
	}
	return res
}

// RoleAuthorityEntityToPO entity转po
func RoleAuthorityEntityToPO(roleAuthority *biz.RoleAuthority) *RoleAuthority {
	return &RoleAuthority{
		ID:            roleAuthority.ID,
		AuthorityID:   roleAuthority.AuthorityID,
		AuthorityType: roleAuthority.AuthorityType,
		RoleID:        roleAuthority.RoleID,
		CreateUserID:  roleAuthority.CreateUserID,
		CreateTime:    roleAuthority.CreateTime,
	}
}

// RoleAuthorityPOListToEntity po转entity
func RoleAuthorityPOListToEntity(roleAuthorityList []*RoleAuthority) []*biz.RoleAuthority {
	res := make([]*biz.RoleAuthority, 0)
	for _, roleAuthority := range roleAuthorityList {
		res = append(res, RoleAuthorityPOToEntity(roleAuthority))
	}
	return res
}

// RoleAuthorityPOToEntity po转entity
func RoleAuthorityPOToEntity(roleAuthority *RoleAuthority) *biz.RoleAuthority {
	return &biz.RoleAuthority{
		ID:            roleAuthority.ID,
		AuthorityID:   roleAuthority.AuthorityID,
		AuthorityType: roleAuthority.AuthorityType,
		RoleID:        roleAuthority.RoleID,
		CreateUserID:  roleAuthority.CreateUserID,
		CreateTime:    roleAuthority.CreateTime,
	}
}

// UserRoleEntityListToPO entity转po
func UserRoleEntityListToPO(userRoleList []*biz.UserRole) []*UserRole {
	res := make([]*UserRole, 0)
	for _, userRole := range userRoleList {
		res = append(res, UserRoleEntityToPO(userRole))
	}
	return res
}

// UserRoleEntityToPO entity转po
func UserRoleEntityToPO(userRole *biz.UserRole) *UserRole {
	return &UserRole{
		ID:           userRole.ID,
		RoleID:       userRole.RoleID,
		UserID:       userRole.UserID,
		CreateUserID: userRole.CreateUserID,
		CreateTime:   userRole.CreateTime,
	}
}

// UserRolePOListToEntity entity转po
func UserRolePOListToEntity(userRoleList []*UserRole) []*biz.UserRole {
	res := make([]*biz.UserRole, 0)
	for _, userRole := range userRoleList {
		res = append(res, UserRolePOToEntity(userRole))
	}
	return res
}

// UserRolePOToEntity entity转po
func UserRolePOToEntity(userRole *UserRole) *biz.UserRole {
	return &biz.UserRole{
		ID:           userRole.ID,
		RoleID:       userRole.RoleID,
		UserID:       userRole.UserID,
		CreateUserID: userRole.CreateUserID,
		CreateTime:   userRole.CreateTime,
	}
}
