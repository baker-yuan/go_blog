package assembler

import (
	oss_model "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/model/oss"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
)

// ModeCfgEntityListToModel entity转pb
func ModeCfgEntityListToModel(ModeCfgList []*oss_model.ModeCfg, uidToNameMap map[uint32]string) []*oss_pb.OssModeCfg {
	res := make([]*oss_pb.OssModeCfg, 0)
	for _, ModeCfg := range ModeCfgList {
		res = append(res, ModeCfgEntityToModel(ModeCfg, uidToNameMap))
	}
	return res
}

// ModeCfgEntityToModel entity转pb
func ModeCfgEntityToModel(ModeCfg *oss_model.ModeCfg, uidToNameMap map[uint32]string) *oss_pb.OssModeCfg {
	modelRes := &oss_pb.OssModeCfg{
		Id:           ModeCfg.ID,
		Mode:         ModeCfg.Mode,
		Path:         ModeCfg.Path,
		Desc:         ModeCfg.Desc,
		CreateUserId: ModeCfg.CreateUserID,
		UpdateUserId: ModeCfg.UpdateUserID,
		CreateTime:   uint32(ModeCfg.CreateTime),
		UpdateTime:   uint32(ModeCfg.UpdateTime),
	}
	modelRes.CreateUser = uidToNameMap[ModeCfg.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[ModeCfg.UpdateUserID]
	return modelRes
}

// AddOrUpdateModeCfgReqToEntity pb转entity
func AddOrUpdateModeCfgReqToEntity(pbModeCfg *oss_pb.AddOrUpdateOssModeCfgReq) *oss_model.ModeCfg {
	entityRes := &oss_model.ModeCfg{
		ID:   pbModeCfg.Id,
		Mode: pbModeCfg.Mode,
		Path: pbModeCfg.Path,
		Desc: pbModeCfg.Desc,
	}
	return entityRes
}
