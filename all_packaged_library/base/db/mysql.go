package db

import (
	"github.com/baker-yuan/go-blog/all_packaged_library/base/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initMySQL() {
	mySQLConf := config.GetMySQLConf()
	for k, v := range mySQLConf {
		dbSession, err := gorm.Open(mysql.Open(v.DataSourceName), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
		mysqlEngineMap[k] = dbSession
		if k == "default" {
			DB = dbSession
		}
	}
}
