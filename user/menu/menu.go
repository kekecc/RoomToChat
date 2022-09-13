package menu

import (
	"fmt"
	"net"
	"room/help"
	"room/user/connect"
	"room/user/show"
	"time"
)

func ShowMenu(conn net.Conn, name string) {
	for {
		fmt.Println("--------------欢迎来到到聊天室------------")
	    fmt.Println("---------------请选择你想干的-------------")
	    fmt.Println("-------------   1.加入群聊    ------------")
	    fmt.Println("-----------   2.和某人私密通话    --------")
		fmt.Println("-----------   3.ping一下服务器  ----------")

		var temp int
		fmt.Scanf("%d\n", &temp)
		
		switch temp {
		case 1:
			//展示历史信息
			show.ShowFormerMes(help.GroupMes)
			go connect.InstantRead(conn)
			go connect.InstantWrite(conn, name)
			time.Sleep(100*time.Minute)
	    case 2:
			var anothername string 
			fmt.Println("请输入你要聊天的人的名字!")
			fmt.Scanf("%s\n", &anothername)
			//log.Println(anothername + "在线吗?")
			fmt.Println("聊天开始了!")
			go connect.InstantReadForPrivate(conn)
			go connect.InstantWriteForPrivate(conn, name, anothername)
			time.Sleep(100*time.Minute)
		case 3:
			fmt.Println("请输入ping")
		}
	}
}