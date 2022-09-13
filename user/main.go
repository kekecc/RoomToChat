package main

import (
	"fmt"
	"net"
	"room/help"
	"room/user/menu"
	"room/user/reorlo"

	"encoding/json"
)

func main() {
	for {
		fmt.Println("--------------欢迎登录到聊天室------------")
	    fmt.Println("---------------请选择你想干的-------------")
	    fmt.Println("---------------   1.注册    --------------")
	    fmt.Println("---------------   2.登录    --------------")
		
		var temp int
	    fmt.Scanf("%d\n", &temp)

		switch temp {
		case 1:
			var name, pwd string
			fmt.Println("请输入您的用户名!")
			fmt.Scanf("%s\n",&name)
			fmt.Println("请输入您的密码!")
			fmt.Scanf("%s\n", &pwd)
			for ;!reorlo.Register(name, pwd); {
				fmt.Println("请输入您的用户名!")
			    fmt.Scanf("%s\n",&name)
			    fmt.Println("请输入您的密码!")
			    fmt.Scanf("%s\n", &pwd)
			}

	    case 2:
			var name, pwd string
			fmt.Println("请输入您的用户名!")
			fmt.Scanf("%s\n",&name)
			fmt.Println("请输入您的密码!")
			fmt.Scanf("%s\n", &pwd)
			for ; !reorlo.Login(name, pwd); {
				fmt.Println("请输入您的用户名!")
			   fmt.Scanf("%s\n",&name)
		       fmt.Println("请输入您的密码!")
			   fmt.Scanf("%s\n", &pwd)
			}
            //登录成功了,获取连接
			conn, err := net.Dial("tcp", ":8080")
			if help.ErrorHandle(err) {
				return
			}
			defer conn.Close()
			data, _ := json.Marshal(name)
			conn.Write(data)
            menu.ShowMenu(conn, name)
	    }
    }
}