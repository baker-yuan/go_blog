package biz

import pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"

// TreeNodeEntityToModel entity转pb
func TreeNodeEntityToModel(treeNode *TreeNode) *pb.TreeNode {
	modelRes := &pb.TreeNode{
		Id:           treeNode.ID,
		Name:         treeNode.Name,
		ParentId:     treeNode.ParentID,
		BizType:      treeNode.BizType,
		CreateUserId: treeNode.CreateUserID,
		UpdateUserId: treeNode.UpdateUserID,
		CreateTime:   treeNode.CreateTime,
		UpdateTime:   treeNode.UpdateTime,
	}
	return modelRes
}

// AddOrUpdateTreeNodeReqToEntity pb转entity
func AddOrUpdateTreeNodeReqToEntity(pbTreeNode *pb.AddOrUpdateTreeNodeReq) *TreeNode {
	entityRes := &TreeNode{
		ID:       pbTreeNode.Id,
		Name:     pbTreeNode.Name,
		ParentID: pbTreeNode.ParentId,
		BizType:  pbTreeNode.BizType,
	}
	return entityRes
}
