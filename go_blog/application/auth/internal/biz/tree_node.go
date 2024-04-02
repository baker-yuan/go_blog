package biz

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// TreeNode 树形结构
type TreeNode struct {
	ID       uint32             // 节点ID
	Name     string             // 节点名称
	ParentID uint32             // 父节点ID
	BizType  pb.TreeNodeBizType // 业务类型
	Children []*TreeNode
	// 公共字段
	CreateUserID uint32
	UpdateUserID uint32
	CreateTime   uint32
	UpdateTime   uint32
}

type TreeNodes []*TreeNode

// TreeNodeRepo 树形结构表repo
type TreeNodeRepo interface {
	// GetTreeNodeByID 根据树形结构表id集合查询树形结构
	GetTreeNodeByID(ctx context.Context, id uint32) (*TreeNode, error)
	// GetTreeNodeByIDs 根据树形结构表id集合查询树形结构
	GetTreeNodeByIDs(ctx context.Context, ids []uint32) (TreeNodes, error)
	// Save 保存树形结构
	Save(ctx context.Context, treeNode *TreeNode) (uint32, error)
	// UpdateByID 根据ID修改树形结构
	UpdateByID(ctx context.Context, treeNode *TreeNode) error
	// DeleteByID 根据ID删除树形结构
	DeleteByID(ctx context.Context, id uint32) error
	// SearchTreeNode 树形结构表搜索
	SearchTreeNode(ctx context.Context, req *pb.SearchTreeNodeReq) (TreeNodes, uint32, error)

	// GetTreeByBizType 根据业务类型查询树形结构
	GetTreeByBizType(ctx context.Context, bizType pb.TreeNodeBizType) (TreeNodes, error)
}

// TreeNodeUseCase 树形结构表管理
type TreeNodeUseCase struct {
	*CommonUseCase
	repo TreeNodeRepo
}

// NewTreeNodeUseCase 创建树形结构表管理业务逻辑实现
func NewTreeNodeUseCase(
	commonUseCase *CommonUseCase,
	repo TreeNodeRepo,
) *TreeNodeUseCase {
	return &TreeNodeUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// TreeNodeDetail 树形结构表详情
func (c *TreeNodeUseCase) TreeNodeDetail(ctx context.Context, req *pb.TreeNodeDetailReq) (*pb.TreeNode, error) {
	treeNode, err := c.repo.GetTreeNodeByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	pbTreeNode := TreeNodeEntityToModel(treeNode)
	return pbTreeNode, nil
}

// SearchTreeNode 树形结构表搜索
func (c *TreeNodeUseCase) SearchTreeNode(ctx context.Context, req *pb.SearchTreeNodeReq) ([]*pb.TreeNode, uint32, error) {
	treeNodes, total, err := c.repo.SearchTreeNode(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.TreeNode, 0)
	for _, treeNode := range treeNodes {
		data = append(data, TreeNodeEntityToModel(treeNode))
	}
	return data, total, nil
}

// AddOrUpdateTreeNode 添加修改树形结构表
func (c *TreeNodeUseCase) AddOrUpdateTreeNode(ctx context.Context, req *pb.AddOrUpdateTreeNodeReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addTreeNode(ctx, userID, req)
	} else {
		dbTreeNode, err := c.repo.GetTreeNodeByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateTreeNode(ctx, dbTreeNode, userID, req)
	}
}

func (c *TreeNodeUseCase) addTreeNode(ctx context.Context, userID uint32, req *pb.AddOrUpdateTreeNodeReq) (uint32, error) {
	treeNode := AddOrUpdateTreeNodeReqToEntity(req)
	treeNode.CreateUserID = userID
	treeNode.CreateTime = uint32(time.Now().Unix())
	treeNode.UpdateUserID = userID
	treeNode.UpdateTime = uint32(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, treeNode)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_TREE_NODE,
		"{}", treeNode,
		"新增树形结构表",
	)

	return lastInsertID, nil
}

func (c *TreeNodeUseCase) updateTreeNode(ctx context.Context, dbTreeNode *TreeNode, userID uint32, req *pb.AddOrUpdateTreeNodeReq) (uint32, error) {
	saveTreeNode := AddOrUpdateTreeNodeReqToEntity(req)
	saveTreeNode.CreateUserID = dbTreeNode.CreateUserID
	saveTreeNode.CreateTime = dbTreeNode.CreateTime
	saveTreeNode.UpdateUserID = userID
	saveTreeNode.UpdateTime = uint32(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveTreeNode); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_TREE_NODE,
		dbTreeNode, saveTreeNode,
		"全字段修改树形结构表",
	)

	return req.GetId(), nil
}

// DeleteTreeNode 删除树形结构表
func (c *TreeNodeUseCase) DeleteTreeNode(ctx context.Context, req *pb.DeleteTreeNodeReq) error {
	treeNode, err := c.repo.GetTreeNodeByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_TREE_NODE,
		treeNode, "{}",
		"删除树形结构表",
	)

	return nil
}

// GetTreeByBizType 获取指定业务类型的树形结构
func (c *TreeNodeUseCase) GetTreeByBizType(ctx context.Context, req *pb.GetTreeByBizTypeReq) ([]*pb.TreeNode, error) {
	treeNodes, err := c.repo.GetTreeByBizType(ctx, req.GetBizType())
	if err != nil {
		return nil, nil
	}
	res := make([]*pb.TreeNode, 0)
	for _, treeNode := range treeNodes {
		res = append(res, TreeNodeEntityToModel(treeNode))
	}
	return c.buildTree(res), nil
}

// buildTree 根据业务类型从平面列表构建树形结构
func (c *TreeNodeUseCase) buildTree(nodes []*pb.TreeNode) []*pb.TreeNode {
	idMap := make(map[uint32]*pb.TreeNode)
	for _, node := range nodes {
		idMap[node.Id] = node
	}
	var rootNodes []*pb.TreeNode
	for _, node := range nodes {
		if node.ParentId == 0 {
			rootNodes = append(rootNodes, node)
			continue
		}
		if parent, ok := idMap[node.ParentId]; ok {
			parent.Children = append(parent.Children, node)
		}
	}
	return rootNodes
}
