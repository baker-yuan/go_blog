package biz

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
	user_pb "github.com/baker-yuan/go-blog/go_blog/protocol/user"
)

// CommonUseCase 通用service
type CommonUseCase struct {
	treeNodeRepo TreeNodeRepo
}

// NewCommonUseCase 创建通用service
func NewCommonUseCase(treeNodeRepo TreeNodeRepo) *CommonUseCase {
	return &CommonUseCase{
		treeNodeRepo: treeNodeRepo,
	}
}

// SaveChangeLog 保存变更日志
func (c *CommonUseCase) SaveChangeLog(ctx context.Context,
	resourceID uint32, resourceType pb.ResourceType,
	changeBefore interface{}, changeAfter interface{},
	notes string) {
	//before := c.convertToJSONString(changeBefore)
	//after := c.convertToJSONString(changeAfter)
	//c.resourceChangeLogRepo.Save(ctx, resourceID, resourceType, before, after, notes)
}

// convertToJSONString 将任意类型的数据转换为JSON字符串
func (c *CommonUseCase) convertToJSONString(data interface{}) string {
	//if data == nil {
	//	return "{}"
	//}
	//switch data.(type) {
	//case string:
	//	return data.(string)
	//default:
	//	return util.StrUtils.Obj2Json(data)
	//}

	return ""
}

// GetReversalTreeMap key=根节点 value=根结点往上找的集合
func (c *CommonUseCase) GetReversalTreeMap(ctx context.Context, bizType pb.TreeNodeBizType) (map[uint32][]*pb.TreeNode, error) {
	treeNodes, err := c.treeNodeRepo.GetTreeByBizType(ctx, bizType)
	if err != nil {
		return nil, nil
	}
	res := make([]*pb.TreeNode, 0)
	for _, treeNode := range treeNodes {
		res = append(res, TreeNodeEntityToModel(treeNode))
	}
	return c.buildLeafToRootMaps(c.buildTree(res)), nil
}

func (c *CommonUseCase) buildLeafToRootMaps(rootNodes []*pb.TreeNode) map[uint32][]*pb.TreeNode {
	leafMap := make(map[uint32][]*pb.TreeNode)
	for _, rootNode := range rootNodes {
		c.buildLeafToRootMap(rootNode, []*pb.TreeNode{}, leafMap)
	}
	return leafMap
}
func (c *CommonUseCase) buildLeafToRootMap(node *pb.TreeNode, path []*pb.TreeNode, leafMap map[uint32][]*pb.TreeNode) {
	// 将当前节点添加到路径中
	currentPath := append([]*pb.TreeNode{node}, path...)
	// 如果当前节点是叶子节点，将路径添加到映射中
	if len(node.Children) == 0 {
		leafMap[node.Id] = currentPath
		return
	}
	// 否则，递归遍历子节点
	for _, child := range node.Children {
		c.buildLeafToRootMap(child, currentPath, leafMap)
	}
}

// buildTree 根据业务类型从平面列表构建树形结构
func (c *CommonUseCase) buildTree(nodes []*pb.TreeNode) []*pb.TreeNode {
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

// UidToName 根据用户ID获取用户名称
func (c *CommonUseCase) UidToName(ctx context.Context, userIDs ...uint32) (map[uint32]string, error) {
	userIDs = util.NewSliceUtils[uint32]().SliceFilter(userIDs, func(u uint32) bool {
		return u == 0
	})
	proxy := user_pb.NewUserApiClientProxy()
	rsp, err := proxy.UidToName(ctx, &user_pb.UidToNameReq{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}
	return rsp.GetUidToNameMap(), nil
}
