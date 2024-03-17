package data

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/common/db"
	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
	"gorm.io/gorm"
)

// Resource 接口资源表
type Resource struct {
	// 基本信息
	ID       uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	ParentID uint32 `gorm:"column:parent_id;type:int unsigned;not null;default:0;comment:父接口ID"`
	Name     string `gorm:"column:name;type:varchar(150);not null;default:'';comment:接口名称"`
	Describe string `gorm:"column:describe;type:varchar(255);not null;default:'';comment:接口描述"`
	//
	ResourceType pb.ResourceResourceType `gorm:"column:resource_type;type:tinyint unsigned;not null;default:0;comment:资源类型 1-目录 2-接口"`
	Status       pb.ResourceStatus       `gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:接口状态 0-未发布 1-已发布 2-以下线"`
	// 路径定位
	URL        string `gorm:"column:url;type:varchar(255);not null;default:'';comment:接口URL（包含路径变量）"`
	HTTPMethod string `gorm:"column:http_method;type:varchar(64);not null;default:'';comment:请求方式 1-get 2-post 3-put 4-patch 5-delete"`
	// 权限校验
	IsNeedLogin      pb.NeedLogin      `gorm:"column:is_need_login;type:tinyint;not null;default:1;comment:是否需要登录校验 0-不需要 1-需要"`
	IsNeedPermission pb.NeedPermission `gorm:"column:is_need_permission;type:tinyint;not null;default:1;comment:是否需要权限校验 0-不需要 1-需要"`
	// 下游服务信息
	Service string `gorm:"column:service;type:varchar(30);not null;default:'';comment:API服务名"`
	Method  string `gorm:"column:method;type:varchar(255);not null;default:'';comment:API接口名"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 AuthResource 结构体对应的数据库表名
func (Resource) TableName() string {
	return ResourceTbName
}

var (
	ResourceTbName = "baker_auth_resource"
	// 基本信息
	ResourceFieldId       = "id"
	ResourceFieldParentId = "parent_id"
	ResourceFieldName     = "name"
	ResourceFieldDescribe = "describe"
	//
	ResourceFieldResourceType = "resource_type"
	ResourceFieldStatus       = "status"
	// 路径定位
	ResourceFieldUrl        = "url"
	ResourceFieldHttpMethod = "http_method"
	// 权限校验
	ResourceFieldIsNeedLogin      = "is_need_login"
	ResourceFieldIsNeedPermission = "is_need_permission"
	// 下游服务信息
	ResourceFieldService = "service"
	ResourceFieldMethod  = "method"
	// 公共字段
	ResourceFieldCreateUser = "create_user_id"
	ResourceFieldUpdateUser = "update_user_id"
	ResourceFieldCreateTime = "create_time"
	ResourceFieldUpdateTime = "update_time"
)

// resourceRepo 接口
type resourceRepo struct {
	*db.GenericDao[Resource, uint32]
}

// 强制resourceRepo实现biz.ResourceRepo
var _ biz.ResourceRepo = &resourceRepo{}

// NewResourceRepo 创建
func NewResourceRepo(data *Data) biz.ResourceRepo {
	return &resourceRepo{
		GenericDao: &db.GenericDao[Resource, uint32]{
			DB: data.GetDB(),
		},
	}
}

func init() {
	registerInitField(initResourceField)
}

var (
	// 全字段修改Resource那些字段不修改
	notUpdateResourceField = []string{
		"created_at",
	}
	updateResourceField []string
)

// InitResourceField 全字段修改
func initResourceField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&Resource{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateResourceField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateResourceField...)
	return nil
}

// GetResourceByID 根据接口id集合查询接口
func (r resourceRepo) GetResourceByID(ctx context.Context, id uint32) (*biz.Resource, error) {
	dbResource, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ResourcePOToEntity(dbResource), err
}

// GetResourceByIDs 根据接口id集合查询接口
func (r resourceRepo) GetResourceByIDs(ctx context.Context, ids []uint32) (biz.Resources, error) {
	dbResources, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return ResourcePOsToEntity(dbResources), nil
}

// Save 保存接口
func (r resourceRepo) Save(ctx context.Context, resource *biz.Resource) (uint32, error) {
	if resource.ID > 0 {
		return 0, errors.New("illegal argument menu id exist")
	}
	dbResource := ResourceEntityToPO(resource)
	if err := r.GenericDao.Create(ctx, dbResource); err != nil {
		return 0, err
	}
	return dbResource.ID, nil
}

// UpdateByID 根据ID修改接口
func (r resourceRepo) UpdateByID(ctx context.Context, resource *biz.Resource) error {
	if resource.ID == 0 {
		return errors.New("illegal argument menu exist")
	}
	dbResource := ResourceEntityToPO(resource)
	return r.GenericDao.DB.WithContext(ctx).Select(updateResourceField).Updates(dbResource).Error
}

// DeleteByID 根据ID删除接口
func (r resourceRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchResource 接口搜索
func (r resourceRepo) SearchResource(ctx context.Context, req *pb.SearchResourceReq) (biz.Resources, uint32, error) {
	var (
		res       []*Resource
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.DB.WithContext(ctx),
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
	return ResourcePOsToEntity(res), uint32(pageTotal), nil
}

// GetEffectiveResource 获取有效状态下的接口
func (r resourceRepo) GetEffectiveResource(ctx context.Context, req *pb.GetEffectiveResourceReq) (biz.Resources, error) {
	res := make([]*Resource, 0)
	tx := r.GenericDao.DB.WithContext(ctx)
	tx = tx.Where(ResourceFieldStatus, pb.ResourceStatus_PUBLISHED)
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return ResourcePOsToEntity(res), nil
}
