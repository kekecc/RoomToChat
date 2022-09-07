package help

import "gorm.io/gorm"


const (  //消息类型
	GroupMes = 3
	PrivateMes = 2
)


type Message struct {
	Type     uint
	Data     string
	Username string
	Toname string
}

type FormerMessage struct {
	gorm.Model
	Data string
	Username string
	Type uint
}

type MessageForPrivate struct {
	gorm.Model
	Data string
	Username string
	Type uint
	Toname string
	
}