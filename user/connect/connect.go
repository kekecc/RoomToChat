package connect

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"room/database"
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
		fmt.Printf("%s:%s", mes.Username, mes.Data)
	}
}

func InstantWrite(conn net.Conn, name string) {
	for {
		//一行发数据
		reader := bufio.NewReader(os.Stdin)
		
		line, err := reader.ReadString('\n')
		temp := strings.Trim(line, " \r\n") 
		if temp == "exit" {
			break
		}
		help.ErrorHandle(err) 
		
		data := help.Message{Type: help.GroupMes, Data: line, Username: name}

		DB := database.GetDB()
		formerdata := help.FormerMessage{Type:help.GroupMes, Data: line, Username: name,}
		DB.Create(&formerdata)

        mes ,_:= json.Marshal(data)
		_, err = conn.Write(mes)
		if help.ErrorHandle(err) {
			break
		}
	}
}


func InstantWriteForPrivate(conn net.Conn, name string, toname string) {
	for {
		//一行发数据
		reader := bufio.NewReader(os.Stdin)
		
		line, err := reader.ReadString('\n')
		temp := strings.Trim(line, " \r\n") 
		if temp == "exit" {
			break
		}
		help.ErrorHandle(err) 
		
		data := help.Message{Type: help.PrivateMes, Data: line, Username: name, Toname: toname}

		//DB := database.GetDB()
		//formerdata := help.MessageForPrivate{Type:help.PrivateMes, Data: line, Username: name,Toname: toname}
		//DB.Create(&formerdata)

        mes ,_:= json.Marshal(data)
		_, err = conn.Write(mes)
		if help.ErrorHandle(err) {
			break
		}
	}
}


func InstantReadForPrivate(conn net.Conn) {
	for {
		var data = make([]byte, 1024)
		length, err := conn.Read(data)
		if help.ErrorHandle(err) {
			break
		}
        //log.Println("开始打印!")
		var mes help.Message
		err = json.Unmarshal(data[:length], &mes)
		if help.ErrorHandle(err) {
			log.Println("解析错误")
			break
		}
		fmt.Printf("%s:%s", mes.Username, mes.Data)
	}
}