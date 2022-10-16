package main

import (
	"fiberscurd/app"
	"fiberscurd/domain"
	"fmt"
)

func main() {
	chn := make(chan string, 3)
	go app.StartApp(chn)
	go domain.ConnDB()
	go domain.RedisConn()
	Servermsg := <-chn
	fmt.Println(Servermsg)
}
