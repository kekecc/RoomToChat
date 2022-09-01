package reorlo

import (
	"fmt"
	"room/database"
	"room/model"

	"golang.org/x/crypto/bcrypt"
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
		hidePwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) //密码加密
		var newUser = model.User {
			UserName: name,
			PassWord: string(hidePwd),
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
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(pwd))
	if err != nil {
		fmt.Printf("密码错误,请重新输入!\n")
		return false
	}
	fmt.Printf("登录成功!\n")
	return true
}