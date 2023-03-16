package db

import "gorm.io/gorm"

var (
	mysqlEngine    *gorm.DB
	mysqlEngineMap map[string]*gorm.DB
)

func Init() {
	initMySQL()
}

func GetMysqlDb(name ...string) *gorm.DB {
	if len(name) == 0 {
		return mysqlEngine
	}
	return mysqlEngineMap[name[0]]
}
