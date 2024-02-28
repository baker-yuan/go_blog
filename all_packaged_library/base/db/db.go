package db

import "gorm.io/gorm"

var (
	DB             *gorm.DB
	mysqlEngineMap = make(map[string]*gorm.DB, 0)
)

func Init() {
	initMySQL()
}

func GetMysqlDb(name ...string) *gorm.DB {
	if len(name) == 0 {
		return DB
	}
	return mysqlEngineMap[name[0]]
}
