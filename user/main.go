package main

import (
	"fmt"
	"room/user/reorlo"
)

func main() {
	for {
		fmt.Printf("--------------欢迎登录到聊天室------------\n")
	    fmt.Printf("---------------请选择你想干的-------------\n")
	    fmt.Printf("---------------   1.注册    --------------\n")
	    fmt.Printf("---------------   2.登录    --------------\n")
	    fmt.Printf("--------------3.进入多人聊天室-------------\n")
	    fmt.Printf("--------------4.和某人加密通话-------------\n")
		
		var temp int
	    fmt.Scanf("%d", &temp)

		switch temp {
		case 1:
			var name, pwd string
			fmt.Printf("请输入您的用户名!\n")
			fmt.Scanf("%s",&name)
			fmt.Printf("请输入您的密码!\n")
			fmt.Scanf("%s", &pwd)
			while(!reorlo.Register(name, pwd)) {
				fmt.Printf("请输入您的用户名!\n")
			    fmt.Scanf("%s",&name)
			    fmt.Printf("请输入您的密码!\n")
			    fmt.Scanf("%s", &pwd)
			}
		}
	    case 2:
			
	}
}  