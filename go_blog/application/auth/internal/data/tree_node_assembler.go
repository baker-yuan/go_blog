package data

import (
	"github.com/baker-yuan/go-blog/go_blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
)

// TreeNodePOsToEntity po转entity
func TreeNodePOsToEntity(dbTreeNodes []*TreeNode) []*biz.TreeNode {
	res := make([]*biz.TreeNode, 0)
	for _, dbTreeNode := range dbTreeNodes {
		res = append(res, TreeNodePOToEntity(dbTreeNode))
	}
	return res
}

// TreeNodePOToEntity po转entity
func TreeNodePOToEntity(dbTreeNode *TreeNode) *biz.TreeNode {
	poRes := &biz.TreeNode{
		ID:           dbTreeNode.ID,
		Name:         dbTreeNode.Name,
		ParentID:     dbTreeNode.ParentID,
		BizType:      dbTreeNode.BizType,
		CreateUserID: dbTreeNode.CreateUserID,
		UpdateUserID: dbTreeNode.UpdateUserID,
		CreateTime:   uint32(dbTreeNode.CreateTime),
		UpdateTime:   uint32(dbTreeNode.UpdateTime),
	}
	return poRes
}

// TreeNodeEntityToPO entity转po
func TreeNodeEntityToPO(treeNode *biz.TreeNode) *TreeNode {
	poRes := &TreeNode{
		ID:           treeNode.ID,
		Name:         treeNode.Name,
		ParentID:     treeNode.ParentID,
		BizType:      treeNode.BizType,
		CreateUserID: treeNode.CreateUserID,
		UpdateUserID: treeNode.UpdateUserID,
		CreateTime:   db.Timestamp(treeNode.CreateTime),
		UpdateTime:   db.Timestamp(treeNode.UpdateTime),
	}
	return poRes
}
