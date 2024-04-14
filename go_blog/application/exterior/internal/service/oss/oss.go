package oss

import (
	"context"

	oss_repo "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/repository/oss"
	"github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/service/oss/assembler"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
)

// OssUploadRecord OSS文件上传记录
type (
	// IOssUploadRecordUseCase 业务逻辑
	IOssUploadRecordUseCase interface {
		// SearchOssUploadRecord 查询OSS文件上传记录
		SearchOssUploadRecord(ctx context.Context, req *oss_pb.SearchOssUploadRecordReq) ([]*oss_pb.OssUploadRecord, uint32, error)
	}
)

// OssUploadRecordUseCase OSS文件上传记录管理
type OssUploadRecordUseCase struct {
	repo oss_repo.IUploadRecordRepo
}

// 强制OssUploadRecordUseCase实现IOssUploadRecordUseCase
var _ IOssUploadRecordUseCase = &OssUploadRecordUseCase{}

// NewOssUploadRecordUseCase 创建OSS文件上传记录管理业务逻辑实现
func NewOssUploadRecordUseCase(
	repo oss_repo.IUploadRecordRepo,
) *OssUploadRecordUseCase {
	return &OssUploadRecordUseCase{
		repo: repo,
	}
}

// SearchOssUploadRecord 查询OSS文件上传记录
func (c OssUploadRecordUseCase) SearchOssUploadRecord(ctx context.Context, req *oss_pb.SearchOssUploadRecordReq) ([]*oss_pb.OssUploadRecord, uint32, error) {
	ossUploadRecords, total, err := c.repo.SearchOssUploadRecord(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	// 补充操作人信息
	//uidToNameMap, err := c.UidToName(ctx, ossUploadRecords.GetOperatorUserID()...)
	//if err != nil {
	//	return nil, 0, err
	//}
	//return assembler.OssUploadRecordEntityListToModel(ossUploadRecords, uidToNameMap), total, nil

	return assembler.OssUploadRecordEntityListToModel(ossUploadRecords, nil), total, nil
}
