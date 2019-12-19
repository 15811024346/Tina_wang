package main

import (
	"../myLog"
	"fmt"
	"time"
)

func main() {
	fmt.Println("111111111")
	//log := myLog.Newlog("info")

	log := myLog.NewFileLogger("info", "./", "wangshuai", 10*1024)
	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		id := 100000
		name := "理想"
		log.Warning("这是一条warning日志 ，id:%d name:%s\n", id, name)
		log.Error("这是一条erroo日志")
		time.Sleep(time.Second * 1)
	}
}
