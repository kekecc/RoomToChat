package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `column:"username"`
	PassWord string `column:"password"`
}