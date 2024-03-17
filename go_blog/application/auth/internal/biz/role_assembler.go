package biz

import (
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// RoleEntityToModel entity转pb
func RoleEntityToModel(role *Role) *pb.Role {
	modelRes := &pb.Role{
		Id:         role.ID,
		Name:       role.Name,
		Code:       role.Code,
		Describe:   role.Describe,
		IsEnable:   role.IsEnable,
		CreateUser: role.CreateUser,
		UpdateUser: role.UpdateUser,
		CreateTime: role.CreateTime,
		UpdateTime: role.UpdateTime,
	}
	return modelRes
}

// AddOrUpdateRoleReqToEntity pb转entity
func AddOrUpdateRoleReqToEntity(pbRole *pb.AddOrUpdateRoleReq) *Role {
	entityRes := &Role{
		// 基本信息
		ID:       pbRole.Id,
		Name:     pbRole.Name,
		Code:     pbRole.Code,
		Describe: pbRole.Describe,
		//
		IsEnable: pbRole.IsEnable,
	}
	return entityRes
}
