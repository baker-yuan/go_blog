package oss

import (
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
)

// UploadRecord OSS文件上传记录表
type UploadRecord struct {
	ID     uint32       `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	UID    uint32       `gorm:"column:uid;type:int unsigned;not null;default:0;comment:用户ID"`
	Patch  string       `gorm:"column:patch;type:varchar(250);not null;default:'';comment:图片路径"`
	Mode   string       `gorm:"column:mode;type:varchar(64);not null;default:'';comment:模块"`
	OSType pb.OssType   `gorm:"column:os_type;type:tinyint;not null;default:0;comment:OSS厂商 1-阿里云"`
	Status pb.OssStatus `gorm:"column:status;type:tinyint;not null;default:0;comment:状态 0-未处理 1-以获取签名 2-上传成功 3-上传失败"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 UploadRecord 结构体对应的数据库表名
func (UploadRecord) TableName() string {
	return UploadRecordTbName
}

var (
	UploadRecordTbName            = "oss_upload_record"
	UploadRecordFieldId           = "id"
	UploadRecordFieldUid          = "uid"
	UploadRecordFieldPatch        = "patch"
	UploadRecordFieldMode         = "mode"
	UploadRecordFieldOsType       = "os_type"
	UploadRecordFieldStatus       = "status"
	UploadRecordFieldCreateTime   = "create_time"
	UploadRecordFieldUpdateTime   = "update_time"
	UploadRecordFieldCreateUserID = "create_user_id"
	UploadRecordFieldUpdateUserID = "update_user_id"
)

type UploadRecords []*UploadRecord

// GetOperatorUserID 获取操作人ID
func (l UploadRecords) GetOperatorUserID() []uint32 {
	res := make([]uint32, 0)
	for _, item := range l {
		res = append(res, item.CreateUserID)
		res = append(res, item.UpdateUserID)
	}
	return util.NewSliceUtils[uint32]().SliceUnique(res)
}
