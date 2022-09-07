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
		fmt.Printf("--------------欢迎来到到聊天室------------\n")
	    fmt.Printf("---------------请选择你想干的-------------\n")
	    fmt.Printf("-------------   1.加入群聊    ------------\n")
	    fmt.Printf("-----------   2.和某人私密通话    --------\n")
		fmt.Printf("-----------   3.ping一下服务器  ----------\n")

		var temp int
		fmt.Scanf("\n%d\n", &temp)
		
		switch temp {
		case 1:
			//展示历史信息
			show.ShowFormerMes(help.GroupMes)
			go connect.InstantRead(conn)
			go connect.InstantWrite(conn, name)
			time.Sleep(100*time.Minute)
	    case 2:
			var anothername string
			fmt.Printf("请输入你要聊天的人的名字!\n")
			fmt.Printf("\n%s\n", anothername)
			go connect.InstantReadForPrivate(conn)
			go connect.InstantWriteForPrivate(conn, name, anothername)
		case 3:
			fmt.Printf("请输入ping\n")
			
		}
	}
}