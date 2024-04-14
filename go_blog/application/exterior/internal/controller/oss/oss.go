package oss

import (
	"context"

	oss_service "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/service/oss"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
)

// OssSrvImpl oss实现
type OssSrvImpl struct {
	oss oss_service.IOssUploadRecordUseCase
}

// 强制OssSrvImpl实现pb.OssApiService
var _ pb.OssApiService = &OssSrvImpl{}

// NewOssApiService 构造函数
func NewOssApiService(oss oss_service.IOssUploadRecordUseCase) pb.OssApiService {
	return &OssSrvImpl{
		oss: oss,
	}
}

// SearchOssUploadRecord OSS文件上传记录搜索
func (o *OssSrvImpl) SearchOssUploadRecord(ctx context.Context, req *pb.SearchOssUploadRecordReq) (*pb.SearchOssUploadRecordRsp, error) {
	res := &pb.SearchOssUploadRecordRsp{}
	records, total, err := o.oss.SearchOssUploadRecord(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Total = total
	res.Data = records
	return res, nil
}
