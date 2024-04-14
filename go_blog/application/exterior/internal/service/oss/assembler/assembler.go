package assembler

import (
	oss_model "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/model/oss"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
)

// OssUploadRecordEntityListToModel entity转pb
func OssUploadRecordEntityListToModel(ossUploadRecordList []*oss_model.UploadRecord, uidToNameMap map[uint32]string) []*pb.OssUploadRecord {
	res := make([]*pb.OssUploadRecord, 0)
	for _, ossUploadRecord := range ossUploadRecordList {
		res = append(res, OssUploadRecordEntityToModel(ossUploadRecord, uidToNameMap))
	}
	return res
}

// OssUploadRecordEntityToModel entity转pb
func OssUploadRecordEntityToModel(ossUploadRecord *oss_model.UploadRecord, uidToNameMap map[uint32]string) *pb.OssUploadRecord {
	modelRes := &pb.OssUploadRecord{
		Id:           ossUploadRecord.ID,
		Uid:          ossUploadRecord.UID,
		Patch:        ossUploadRecord.Patch,
		Mode:         ossUploadRecord.Mode,
		OsType:       ossUploadRecord.OSType,
		Status:       ossUploadRecord.Status,
		CreateUserId: ossUploadRecord.CreateUserID,
		UpdateUserId: ossUploadRecord.UpdateUserID,
		CreateTime:   uint32(ossUploadRecord.CreateTime),
		UpdateTime:   uint32(ossUploadRecord.UpdateTime),
	}
	modelRes.CreateUser = uidToNameMap[ossUploadRecord.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[ossUploadRecord.UpdateUserID]
	return modelRes
}
