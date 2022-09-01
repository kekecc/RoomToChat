package database

import (
	"room/help"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
	help.ErrorHandle(err) 

	//自动迁移创建表
	DB.AutoMigrate()
	return DB
}

var DB = InitDB()

func GetDB() *gorm.DB{
	return DB
}