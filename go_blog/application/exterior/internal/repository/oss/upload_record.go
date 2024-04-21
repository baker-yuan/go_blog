package oss

import (
	"context"
	"errors"
	"fmt"

	oss_model "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/model/oss"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
	"gorm.io/gorm"
)

// OssUploadRecord OSS文件上传记录
type (
	// IUploadRecordRepo 数据存储操作
	IUploadRecordRepo interface {
		// GetOssUploadRecordByID 根据OSS文件上传记录id集合查询OSS文件上传记录
		GetOssUploadRecordByID(ctx context.Context, id uint32) (*oss_model.UploadRecord, error)
		// GetOssUploadRecordByIDs 根据OSS文件上传记录id集合查询OSS文件上传记录
		GetOssUploadRecordByIDs(ctx context.Context, ids []uint32) (oss_model.UploadRecords, error)
		// Save 保存OSS文件上传记录
		Save(ctx context.Context, ossUploadRecord *oss_model.UploadRecord) (uint32, error)
		// UpdateByID 根据ID修改OSS文件上传记录
		UpdateByID(ctx context.Context, ossUploadRecord *oss_model.UploadRecord) error
		// DeleteByID 根据ID删除OSS文件上传记录
		DeleteByID(ctx context.Context, id uint32) error
		// SearchOssUploadRecord OSS文件上传记录搜索
		SearchOssUploadRecord(ctx context.Context, req *oss_pb.SearchOssUploadRecordReq) (oss_model.UploadRecords, uint32, error)
		// GetOssUploadRecordByPatch 根据OSS文件上传path查询OSS文件上传记录
		GetOssUploadRecordByPatch(ctx context.Context, patch string) (*oss_model.UploadRecord, error)
	}
)

// UploadRecordRepo OSS文件上传记录
type UploadRecordRepo struct {
	*db.GenericDao[oss_model.UploadRecord, uint32]
}

// NewUploadRecordRepo 创建
func NewUploadRecordRepo(gormDB *gorm.DB) *UploadRecordRepo {
	return &UploadRecordRepo{
		GenericDao: db.NewGenericDao[oss_model.UploadRecord, uint32](gormDB),
	}
}

// 强制UploadRecordRepo实现IOssUploadRecordUseCase
var _ IUploadRecordRepo = &UploadRecordRepo{}

var (
	// 全字段修改OssUploadRecord那些字段不修改
	notUpdateOssUploadRecordField = []string{
		oss_model.UploadRecordFieldCreateTime,
		oss_model.UploadRecordFieldCreateUserID,
	}
	updateOssUploadRecordField []string
)

// InitOssUploadRecordField 全字段修改
func InitOssUploadRecordField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(oss_model.UploadRecord{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateOssUploadRecordField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateOssUploadRecordField...)
	return nil
}

// GetOssUploadRecordByID 根据OSS文件上传记录id查询OSS文件上传记录
func (r *UploadRecordRepo) GetOssUploadRecordByID(ctx context.Context, id uint32) (*oss_model.UploadRecord, error) {
	dbOssUploadRecord, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbOssUploadRecord, nil
}

// GetOssUploadRecordByIDs 根据OSS文件上传记录id集合查询OSS文件上传记录
func (r *UploadRecordRepo) GetOssUploadRecordByIDs(ctx context.Context, ids []uint32) (oss_model.UploadRecords, error) {
	dbOssUploadRecords, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return dbOssUploadRecords, nil
}

// Save 保存OSS文件上传记录
func (r *UploadRecordRepo) Save(ctx context.Context, ossUploadRecord *oss_model.UploadRecord) (uint32, error) {
	if ossUploadRecord.ID > 0 {
		return 0, errors.New("illegal argument ossUploadRecord id exist")
	}
	if err := r.GenericDao.Create(ctx, ossUploadRecord); err != nil {
		return 0, err
	}
	return ossUploadRecord.ID, nil
}

// UpdateByID 根据ID修改OSS文件上传记录
func (r *UploadRecordRepo) UpdateByID(ctx context.Context, ossUploadRecord *oss_model.UploadRecord) error {
	if ossUploadRecord.ID == 0 {
		return errors.New("illegal argument ossUploadRecord exist")
	}
	return r.GenericDao.GetDB(ctx).Select(updateOssUploadRecordField).Updates(ossUploadRecord).Error
}

// DeleteByID 根据ID删除OSS文件上传记录
func (r *UploadRecordRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchOssUploadRecord OSS文件上传记录搜索
func (r *UploadRecordRepo) SearchOssUploadRecord(ctx context.Context, req *oss_pb.SearchOssUploadRecordReq) (oss_model.UploadRecords, uint32, error) {
	var (
		res       []*oss_model.UploadRecord
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
	tx = tx.Order(fmt.Sprintf("%s DESC", oss_model.UploadRecordFieldId))
	tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).
		Limit(int(req.GetPageSize())).Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}

// GetOssUploadRecordByPatch 根据OSS文件上传path查询OSS文件上传记录
func (r *UploadRecordRepo) GetOssUploadRecordByPatch(ctx context.Context, patch string) (*oss_model.UploadRecord, error) {
	tx := r.GenericDao.GetDB(ctx)
	tx = tx.Where(oss_model.UploadRecordFieldPatch, patch)
	res := &oss_model.UploadRecord{}
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	if res.ID == 0 {
		return nil, nil
	}
	return res, nil
}
