// Package db grom初始化
package db

import (
	"gorm.io/gorm"
)

type DBOption func(*gorm.DB) *gorm.DB
type DBOrderOption func(*gorm.DB) *gorm.DB
