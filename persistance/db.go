package persistance

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"training/go-products/config"
)

var DB *gorm.DB

func CreateConnection() {
	db, err := gorm.Open(mysql.Open(GetConnectionString()), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db
}

func GetConnectionString() string {
	var (
		username = config.AppConf.Db.Username
		password = config.AppConf.Db.Password
		host     = config.AppConf.Db.Host
		dbName   = config.AppConf.Db.Name
		dsn      = username + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	)
	return dsn
}
