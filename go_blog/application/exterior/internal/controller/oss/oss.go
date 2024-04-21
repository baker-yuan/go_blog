package oss

import (
	"context"

	oss_service "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/service/oss"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
)

// SrvImpl oss实现
type SrvImpl struct {
	oss     oss_service.IUploadRecordUseCase
	ModeCfg oss_service.IModeCfgUseCase
}

// 强制OssSrvImpl实现oss_pb.OssApiService
var _ oss_pb.OssApiService = &SrvImpl{}

// NewOssApiService 构造函数
func NewOssApiService(
	oss oss_service.IUploadRecordUseCase,
	ModeCfg oss_service.IModeCfgUseCase,
) oss_pb.OssApiService {
	return &SrvImpl{
		oss:     oss,
		ModeCfg: ModeCfg,
	}
}

// SearchOssUploadRecord OSS文件上传记录搜索
func (o *SrvImpl) SearchOssUploadRecord(ctx context.Context, req *oss_pb.SearchOssUploadRecordReq) (*oss_pb.SearchOssUploadRecordRsp, error) {
	res := &oss_pb.SearchOssUploadRecordRsp{}
	records, total, err := o.oss.SearchOssUploadRecord(ctx, req)
	if err != nil {
		return nil, err
	}
	res.PageTotal = total
	res.Data = records
	return res, nil
}

// GetUploadSign 获取oss上传签名
func (o *SrvImpl) GetUploadSign(ctx context.Context, req *oss_pb.GetUploadSignReq) (*oss_pb.GetUploadSignRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := o.oss.GetUploadSign(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SearchOssModeCfg 查询模块配置
func (o *SrvImpl) SearchOssModeCfg(ctx context.Context, req *oss_pb.SearchOssModeCfgReq) (*oss_pb.SearchOssModeCfgRsp, error) {
	rsp := &oss_pb.SearchOssModeCfgRsp{}
	modeCfgList, pageTotal, err := o.ModeCfg.SearchModeCfg(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = modeCfgList
	rsp.PageTotal = pageTotal
	return rsp, nil
}

// OssModeCfgDetail 模块配置详情
func (o *SrvImpl) OssModeCfgDetail(ctx context.Context, req *oss_pb.OssModeCfgDetailReq) (*oss_pb.OssModeCfg, error) {
	return o.ModeCfg.ModeCfgDetail(ctx, req)
}

// AddOrUpdateOssModeCfg 添加修改模块配置
func (o *SrvImpl) AddOrUpdateOssModeCfg(ctx context.Context, req *oss_pb.AddOrUpdateOssModeCfgReq) (*oss_pb.AddOrUpdateRsp, error) {
	rsp := &oss_pb.AddOrUpdateRsp{}
	id, err := o.ModeCfg.AddOrUpdateModeCfg(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteOssModeCfg 删除模块配置
func (o *SrvImpl) DeleteOssModeCfg(ctx context.Context, req *oss_pb.DeleteOssModeCfgReq) (*oss_pb.EmptyRsp, error) {
	rsp := &oss_pb.EmptyRsp{}
	if err := o.ModeCfg.DeleteModeCfg(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}
