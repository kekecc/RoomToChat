package reorlo

import (
	"fmt"
	"room/database"
	"room/model"
)

func Register(name string, pwd string) bool{
	DB := database.GetDB()
	var user model.User
    //查询该用户名是否注册过
	DB.Where("username = ?", name).First(&user)
	if user.ID != 0 {
		fmt.Printf("该用户名已注册,请更换!\n")
		return false
	} else {
		var newUser = model.User {
			UserName: name,
			PassWord: pwd,
		}
		DB.Create(&newUser)
		fmt.Printf("注册成功!")
	}
	return true
}

func Login(name, pwd string) bool {
	DB := database.GetDB()
	var user model.User
	//查询该用户是否存在
	DB.Where("username = ?", name).First(&user)
	if user.ID == 0 {
		fmt.Printf("用户名不存在,请重新输入!\n")
	}
	
}