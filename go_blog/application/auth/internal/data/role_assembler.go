package data

import (
	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/common/db"
)

// RolePOsToEntity po转entity
func RolePOsToEntity(dbRoles []*Role) []*biz.Role {
	res := make([]*biz.Role, 0)
	for _, dbRole := range dbRoles {
		res = append(res, RolePOToEntity(dbRole))
	}
	return res
}

// RolePOToEntity po转entity
func RolePOToEntity(dbRole *Role) *biz.Role {
	poRes := &biz.Role{
		ID:         dbRole.ID,
		Name:       dbRole.Name,
		Code:       dbRole.Code,
		Describe:   dbRole.Describe,
		IsEnable:   bool(dbRole.IsEnable),
		CreateUser: dbRole.CreateUser,
		UpdateUser: dbRole.UpdateUser,
		CreateTime: uint32(dbRole.CreateTime),
		UpdateTime: uint32(dbRole.UpdateTime),
	}
	return poRes
}

// RoleEntityToPO entity转po
func RoleEntityToPO(role *biz.Role) *Role {
	poRes := &Role{
		ID:         role.ID,
		Name:       role.Name,
		Code:       role.Code,
		Describe:   role.Describe,
		IsEnable:   db.BoolBit(role.IsEnable),
		CreateUser: role.CreateUser,
		UpdateUser: role.UpdateUser,
		CreateTime: db.Timestamp(role.CreateTime),
		UpdateTime: db.Timestamp(role.UpdateTime),
	}
	return poRes
}
