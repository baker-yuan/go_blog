package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Timestamp 用于处理数据库中 timestamp 类型和 Go 中 uint32 类型之间的转换
type Timestamp uint32

// Scan 实现 sql.Scanner 接口，用于从数据库读取值时的自定义扫描逻辑
func (ts *Timestamp) Scan(value interface{}) error {
	if value == nil {
		*ts = 0
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		return errors.New("timestamp scan: type assertion to time.Time failed")
	}

	*ts = Timestamp(t.Unix())
	return nil
}

// Value 实现 driver.Valuer 接口，用于写入数据库时的自定义值逻辑
func (ts Timestamp) Value() (driver.Value, error) {
	// 将 uint32 转换为 time.Time
	t := time.Unix(int64(ts), 0)
	return t, nil
}

// BoolBit 用于处理 MySQL 中 bit(1) 类型和 Go 中 bool 类型之间的转换
type BoolBit bool

// Scan 实现 sql.Scanner 接口，用于从数据库读取值时的自定义扫描逻辑
func (bb *BoolBit) Scan(value interface{}) error {
	if value == nil {
		*bb = false
		return nil
	}

	bv, ok := value.([]byte)
	if !ok {
		return errors.New("boolBit scan: type assertion to []byte failed")
	}

	*bb = bv[0] == 1
	return nil
}

// Value 实现 driver.Valuer 接口，用于写入数据库时的自定义值逻辑
func (bb BoolBit) Value() (driver.Value, error) {
	if bb {
		return []byte{1}, nil
	}
	return []byte{0}, nil
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
