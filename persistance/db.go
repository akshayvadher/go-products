package persistance

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"training/go-products/config"
)

var DB *gorm.DB

func CreateConnection() {
	var (
		username = config.GetEnv("DB_USERNAME", "root")
		password = config.GetEnv("DB_PASSWORD", "my-secret-pw")
		host     = config.GetEnv("DB_HOST", "127.0.0.1:3306")
		dbName   = config.GetEnv("DB_NAME", "products")
		dsn      = username + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db

}
