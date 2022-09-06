package show

import (
	"fmt"
	"room/database"
	"room/help"
)

func ShowFormerMes(form uint) {
	var datas []help.FormerMessage
	DB := database.GetDB()
	
	DB.Model(&help.FormerMessage{}).Where("type = ?", form).Find(&datas) // 查询到了所有的字段
	//展示在屏幕上
	for i := range datas {
		fmt.Printf("%s:%s", datas[i].Username, datas[i].Data)
	}
	
}