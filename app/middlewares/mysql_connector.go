package middlewares

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	dsn := "root:rootpass@tcp(db:3306)/godb?charset=utf8mb4&parseTime=True"

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = gormDB
}

func DB() *gorm.DB {
	return db
}
