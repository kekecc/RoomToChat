package handle

import (
	"encoding/json"
	"fmt"
	"room/help"
)

type ClientManager struct {
	Clients    map[string]*Client //维护所有在线用户
	BroadCast  chan []byte        //广播的消息
	Register   chan *Client
	UnRegister chan *Client
}

func InitManager() *ClientManager {
	manager := &ClientManager{
		Clients:    make(map[string]*Client),
		BroadCast:  make(chan []byte, 1024),
		Register:   make(chan *Client, 1024),
		UnRegister: make(chan *Client, 1024),
	}
	return manager
}

var Manager = InitManager() // 全局的clientManager

func (mger *ClientManager) Start() {
	for {
		select {
		case client := <-mger.Register:
			mger.Clients[client.UserName] = client
			//向服务器推送
			length := len(mger.Clients)
			resp, err := json.Marshal(&help.Message{Type: 1, Data: fmt.Sprintf("%s 已经上线了,当前在线人数:%d\n", client.UserName, length)})
			if help.ErrorHandle(err) {
				return
			}
			mger.BroadCast <- resp
		}
	}
}

func (mger *ClientManager) BroadCastMes() {
	for {
		select {
		case message := <- mger.BroadCast :
			//向所有在线用户推送
			for _, client := range mger.Clients {
				client.Send <- message
			}
		}
	}
}

func (mger *ClientManager) ClientQuit() {
	for {
		select {
		case client := <- mger.UnRegister :
			//删除用户
			delete(mger.Clients, client.UserName)
			//推送新的在线人数
			length := len(mger.Clients)
			resp, _ := json.Marshal(&help.Message{Type: 1, Data: fmt.Sprintf("%s已下线,当前在线人数：%d\n", client.UserName, length)})
			mger.BroadCast <- resp
		}
	}
}