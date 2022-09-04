package main

import (
	"encoding/json"
	"fmt"
	"net"
	"room/help"
	"room/server/handle"

	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080") 
	fmt.Println("服务器启动")
	if(help.ErrorHandle(err)) {
		return
	}
	
	go handle.Manager.BroadCastMes()
	go handle.Manager.ClientQuit()
	go handle.Manager.Start()

	for {
		conn, err := listener.Accept()
		if help.ErrorHandle(err) {
			return
		}
		
		var data = make([]byte, 1024)
		length, err := conn.Read(data)
		if help.ErrorHandle(err) {
			break
		}

		var mes string
		err = json.Unmarshal(data[:length], &mes)
		if help.ErrorHandle(err) {
			break
		}

		client := &handle.Client{
			UserName: mes,
			Socket: conn,
			Send: make(chan []byte, 1024),
			ReceiveTime: time.Now(),
			ExpireTime: 10 * time.Minute,
		}
		
		fmt.Printf("%s已连接\n", client.UserName)
		
		handle.Manager.Register <- client
		go client.ReadMes()
		go client.CheckTime()
		go client.WriteMes()
	}
}