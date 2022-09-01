package main

import (
	"fmt"
	"net"
	"room/help"
	"room/user/menu"
	"room/user/reorlo"
)

func main() {
	for {
		fmt.Printf("--------------欢迎登录到聊天室------------\n")
	    fmt.Printf("---------------请选择你想干的-------------\n")
	    fmt.Printf("---------------   1.注册    --------------\n")
	    fmt.Printf("---------------   2.登录    --------------\n")
		
		var temp int
	    fmt.Scanf("%d", &temp)

		switch temp {
		case 1:
			var name, pwd string
			fmt.Printf("请输入您的用户名!\n")
			fmt.Scanf("%s",&name)
			fmt.Printf("请输入您的密码!\n")
			fmt.Scanf("%s", &pwd)
			for ;!reorlo.Register(name, pwd); {
				fmt.Printf("请输入您的用户名!\n")
			    fmt.Scanf("%s",&name)
			    fmt.Printf("请输入您的密码!\n")
			    fmt.Scanf("%s", &pwd)
			}

	    case 2:
			var name, pwd string
			fmt.Printf("请输入您的用户名!\n")
			fmt.Scanf("%s",&name)
			fmt.Printf("请输入您的密码!\n")
			fmt.Scanf("%s", &pwd)
			for ; !reorlo.Login(name, pwd); {
				fmt.Printf("请输入您的用户名!\n")
			    fmt.Scanf("%s",&name)
			    fmt.Printf("请输入您的密码!\n")
			    fmt.Scanf("%s", &pwd)
			}
            //登录成功了,获取连接
			conn, err := net.Dial("tcp", ":8080")
			if help.ErrorHandle(err) {
				return
			}
            menu.ShowMenu(conn)
	    }
    }
}