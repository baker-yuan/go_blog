package biz

import (
	"strings"

	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// TreeNodeEntityListToModel entity转pb
func TreeNodeEntityListToModel(treeNodeList []*TreeNode, uidToNameMap map[uint32]string, treeNodeMap map[uint32][]*pb.TreeNode) []*pb.TreeNode {
	res := make([]*pb.TreeNode, 0)
	for _, treeNode := range treeNodeList {
		res = append(res, TreeNodeEntityToModel(treeNode, uidToNameMap, treeNodeMap))
	}
	return res
}

// TreeNodeEntityToModel entity转pb
func TreeNodeEntityToModel(treeNode *TreeNode, uidToNameMap map[uint32]string, treeNodeMap map[uint32][]*pb.TreeNode) *pb.TreeNode {
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
	modelRes.CreateUser = uidToNameMap[treeNode.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[treeNode.UpdateUserID]

	if node, exist := treeNodeMap[treeNode.ID]; exist {
		treeNodePath := make([]string, 0)
		for i := len(node) - 1; i >= 0; i-- {
			treeNodePath = append(treeNodePath, node[i].Name)
		}
		modelRes.TreeNodePath = strings.Join(treeNodePath, "/")
	}

	switch treeNode.BizType {
	case pb.TreeNodeBizType_RESOURCE_NODE:
		modelRes.BizTypeDesc = "接口资源"
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
