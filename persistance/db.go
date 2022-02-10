package persistance

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"training/go-products/config"
)

var DB *gorm.DB

func CreateConnection() {
	var (
		username = config.AppConf.Db.Username
		password = config.AppConf.Db.Password
		host     = config.AppConf.Db.Host
		dbName   = config.AppConf.Db.Name
		dsn      = username + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db

}
