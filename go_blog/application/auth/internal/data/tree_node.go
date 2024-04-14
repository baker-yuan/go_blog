package data

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/go_blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
	"gorm.io/gorm"
)

// TreeNode 树形结构
type TreeNode struct {
	ID       uint32             `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:节点ID"`
	ParentID uint32             `gorm:"column:parent_id;type:int unsigned;not null;default:0;comment:父节点ID"`
	Name     string             `gorm:"column:name;type:varchar(20);not null;default:'';comment:节点名称"`
	BizType  pb.TreeNodeBizType `gorm:"column:biz_type;type:tinyint(4);unsigned;not null;default:0;comment:业务类型"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 TreeNode 结构体对应的数据库表名
func (r TreeNode) TableName() string {
	return TreeNodeTbName
}

var (
	TreeNodeTbName        = "baker_auth_tree_node"
	TreeNodeFieldID       = "id"
	TreeNodeFieldParentID = "parent_id"
	TreeNodeFieldName     = "name"
	TreeNodeFieldBizType  = "biz_type"
)

// TreeNodeRepo 树形结构表
type treeNodeRepo struct {
	*db.GenericDao[TreeNode, uint32]
}

// NewTreeNodeRepo 创建
func NewTreeNodeRepo(data *Data) biz.TreeNodeRepo {
	return &treeNodeRepo{
		GenericDao: db.NewGenericDao[TreeNode, uint32](data.GetDB()),
	}
}

func init() {
	registerInitField(initTreeNodeField)
}

var (
	// 全字段修改TreeNode那些字段不修改
	notUpdateTreeNodeField = []string{
		"created_at",
		"create_time",
	}
	updateTreeNodeField []string
)

// InitTreeNodeField 全字段修改
func initTreeNodeField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(TreeNode{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateTreeNodeField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateTreeNodeField...)
	return nil
}

// GetTreeNodeByID 根据树形结构表id查询树形结构表
func (r *treeNodeRepo) GetTreeNodeByID(ctx context.Context, id uint32) (*biz.TreeNode, error) {
	dbTreeNode, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return TreeNodePOToEntity(dbTreeNode), err
}

// GetTreeNodeByIDs 根据树形结构表id集合查询树形结构表
func (r *treeNodeRepo) GetTreeNodeByIDs(ctx context.Context, ids []uint32) (biz.TreeNodes, error) {
	dbTreeNodes, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return TreeNodePOsToEntity(dbTreeNodes), nil
}

// Save 保存树形结构表
func (r *treeNodeRepo) Save(ctx context.Context, treeNode *biz.TreeNode) (uint32, error) {
	if treeNode.ID > 0 {
		return 0, errors.New("illegal argument treeNode id exist")
	}
	dbTreeNode := TreeNodeEntityToPO(treeNode)
	if err := r.GenericDao.Create(ctx, dbTreeNode); err != nil {
		return 0, err
	}
	return treeNode.ID, nil
}

// UpdateByID 根据ID修改树形结构表
func (r *treeNodeRepo) UpdateByID(ctx context.Context, treeNode *biz.TreeNode) error {
	if treeNode.ID == 0 {
		return errors.New("illegal argument treeNode exist")
	}
	dbTreeNode := TreeNodeEntityToPO(treeNode)
	return r.GenericDao.GetDB(ctx).Select(updateTreeNodeField).Updates(dbTreeNode).Error
}

// DeleteByID 根据ID删除树形结构表
func (r *treeNodeRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchTreeNode 树形结构表搜索
func (r *treeNodeRepo) SearchTreeNode(ctx context.Context, req *pb.SearchTreeNodeReq) (biz.TreeNodes, uint32, error) {
	var (
		res       []*TreeNode
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.GetDB(ctx),
		func(search map[string]*db.SearchValue) {

		},
	)
	if err != nil {
		return nil, 0, err
	}
	tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).
		Limit(int(req.GetPageSize())).Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	return TreeNodePOsToEntity(res), uint32(pageTotal), nil
}

// GetTreeByBizType 根据业务类型查询数
func (r *treeNodeRepo) GetTreeByBizType(ctx context.Context, bizType pb.TreeNodeBizType) (biz.TreeNodes, error) {
	treeNodes := make([]*TreeNode, 0)
	tx := r.GenericDao.GetDB(ctx)
	tx = tx.Where(TreeNodeFieldBizType, bizType)
	if err := tx.Find(&treeNodes).Error; err != nil {
		return nil, err
	}
	return TreeNodePOsToEntity(treeNodes), nil
}
