package connect

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"room/help"
	"strings"
)

func InstantRead(conn net.Conn) {
	for {
		var data = make([]byte, 1024)
		length, err := conn.Read(data)
		if help.ErrorHandle(err) {
			break
		}

		var mes help.Message
		err = json.Unmarshal(data[:length], &mes)
		if help.ErrorHandle(err) {
			break
		}

		fmt.Printf("say:%s", mes.Data)
	}
}

func InstantWrite(conn net.Conn) {
	for {
		//一行发数据
		reader := bufio.NewReader(os.Stdin)
		
		line, err := reader.ReadString('\n')
		temp := strings.Trim(line, " \r\n") 
		if temp == "exit" {
			break
		}
		help.ErrorHandle(err) 
		
		data := help.Message{Type: 3, Data: line}
        mes ,_:= json.Marshal(data)
		_, err = conn.Write(mes)
		if help.ErrorHandle(err) {
			break
		}
	}
}