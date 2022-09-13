package help

import (
	"log"
	"math/rand"
	"time"
)

func ErrorHandle(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func RandomName() string {
	//提供一个随机的名称
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	rname := make([]byte, 5)
	rand.Seed(time.Now().Unix())
	for i := range rname {
		rname[i]  = letters[rand.Intn(len(letters))]
	}
	return string(rname)
}