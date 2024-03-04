package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// BaseTime 基础的时间字段，每个dao结构体都应该嵌套这个字段
type BaseTime struct {
	CreatedAt uint32 `gorm:"column:created_at;type:int(10);unsigned;not null;default:0;comment:创建时间"` // 使用时间戳秒数填充创建时间
	UpdatedAt uint32 `gorm:"column:updated_at;type:int(10);unsigned;not null;default:0;comment:修改时间"` // 使用时间戳秒数填充修改时间
}

// BeforeCreate gorm创建之前回调
func (bt *BaseTime) BeforeCreate(tx *gorm.DB) (err error) {
	timestamp := uint32(time.Now().Unix())
	bt.CreatedAt = timestamp
	bt.UpdatedAt = timestamp
	return
}

// BeforeUpdate gorm修改之前回调
func (bt *BaseTime) BeforeUpdate(tx *gorm.DB) (err error) {
	bt.UpdatedAt = uint32(time.Now().Unix())
	return
}

// Operator 操作人
type Operator struct {
	Creator string `gorm:"column:creator;type:varchar(20);not null;default:'';comment:创建人"`
	Reviser string `gorm:"column:reviser;type:varchar(20);not null;default:'';comment:修改人"`
}

// SoftDelete 软删除
type SoftDelete struct {
	//IsDeleted soft_delete.DeletedAt `gorm:"column:is_deleted;default:0;softDelete:flag;comment:软删除"`
}

// Init 初始化
func Init(db *gorm.DB) error {
	schemas := []tableSchema{}

	return autoMigrate(db, schemas)
}

// tableSchema 自动建表描述信息
type tableSchema struct {
	TableName string      // 表名
	StructPtr interface{} // 结构体指针
}

func autoMigrate(db *gorm.DB, schemas []tableSchema) error {
	for _, schema := range schemas {
		if err := db.
			Set("gorm:table_options", fmt.Sprintf("CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='%s'", schema.TableName)).
			AutoMigrate(schema.StructPtr); err != nil {
			return err
		}
	}
	return nil
}
