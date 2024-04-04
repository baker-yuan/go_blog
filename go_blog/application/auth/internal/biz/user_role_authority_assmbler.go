package biz

import pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"

// UserRoleEntityListToModel entity转pb
func UserRoleEntityListToModel(treeNodeList []*UserRole) []*pb.UserRole {
	res := make([]*pb.UserRole, 0)
	for _, treeNode := range treeNodeList {
		res = append(res, UserRoleEntityToModel(treeNode))
	}
	return res
}

// UserRoleEntityToModel entity转pb
func UserRoleEntityToModel(treeNode *UserRole) *pb.UserRole {
	modelRes := &pb.UserRole{
		Id:           treeNode.ID,
		RoleId:       treeNode.RoleID,
		UserId:       treeNode.UserID,
		CreateUserId: treeNode.CreateUserID,
		CreateTime:   uint32(treeNode.CreateTime),
	}
	return modelRes
}
