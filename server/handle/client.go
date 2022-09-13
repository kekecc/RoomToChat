package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"room/help"
	"time"
)

type Client struct {
	UserName string
	Socket net.Conn         //保存连接
	Send chan []byte         //保存要发送的信息
	ReceiveTime time.Time    //创建时间
	ExpireTime time.Duration //过期时间
}

func (c *Client) ReadMes() {
	defer func() {  //如果出现意外则关闭并删除该连接
		err := c.Socket.Close()
		if help.ErrorHandle(err) {
			return
		}
		Manager.UnRegister <- c
	}()
	
	for {
		var data = make([]byte, 1024)
		length, err := c.Socket.Read(data)
		if help.ErrorHandle(err) {
			return
		}

		var mes help.Message
		err = json.Unmarshal(data[:length], &mes)
		if help.ErrorHandle(err) {
			log.Println("服务器端解析json出错!")
			return
		}

		switch mes.Type {
		case 6:
			resp, _ := json.Marshal(&help.Message{Type: 6, Data: "pong", Username: "服务器"})
			c.ReceiveTime = time.Now()
			c.Send <- resp

		case 1:
			//获取在线人数
			length := len(Manager.Clients)
			resp, _ := json.Marshal(&help.Message{Type: 1, Data: fmt.Sprintf("当前在线人数:%d", length), Username: "服务器"})
			c.ReceiveTime = time.Now()
			c.Send <- resp

		case help.PrivateMes:
			resp,_ := json.Marshal(&help.MessageForPrivate{Type:help.PrivateMes, Data: mes.Data, Username: c.UserName, Toname: mes.Toname})
			log.Println("测试 发送1")
			Manager.PrivateSend <- resp

		case help.GroupMes:
			// 广播消息
			resp, _ := json.Marshal(&help.Message{Type: help.GroupMes, Data: mes.Data, Username: c.UserName})
			Manager.BroadCast <- resp
		}
	}
}

func (c *Client) WriteMes() {
	defer func() {  //如果出现意外则关闭并删除该连接
		err := c.Socket.Close()
		if help.ErrorHandle(err) {
			return
		}
		Manager.UnRegister <- c
	}()
	
	for {
		select {
		case mes := <- c.Send : //从管道读取信息
		    log.Println("管道读消息")
			_, err := c.Socket.Write(mes)
			if help.ErrorHandle(err) {
				return
			}
		}
	}
}


func (c *Client) CheckTime() {
	//查看用户是否过期
	for {
		nowTime := time.Now()
		durationTime := nowTime.Sub(c.ReceiveTime)

		if durationTime >= c.ExpireTime {
			Manager.UnRegister <- c
			break
		}
	}
}