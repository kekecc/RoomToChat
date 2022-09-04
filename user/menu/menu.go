package menu

import (
	"fmt"
	"net"
	"room/user/connect"
)

func ShowMenu(conn net.Conn) {
	for {
		fmt.Printf("--------------欢迎来到到聊天室------------\n")
	    fmt.Printf("---------------请选择你想干的-------------\n")
	    fmt.Printf("-------------   1.加入群聊    ------------\n")
	    fmt.Printf("-----------   2.和某人私密通话    --------\n")
		fmt.Printf("-----------   3.ping一下服务器  ----------\n")

		var temp int
		fmt.Scanf("%d", &temp)
		
		switch temp {
		case 1:
			//展示历史信息
			go connect.InstantRead(conn)
			go connect.InstantWrite(conn)
	    case 2:
			//请输入你要对话的人的名字!
		case 3:
			fmt.Printf("请输入ping\n")
			
		}
	}
}