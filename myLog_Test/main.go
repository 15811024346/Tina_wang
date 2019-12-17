package main

import (
	"../myLog"
	"fmt"
	"time"
)

func main() {
	fmt.Println("111111111")
	log := myLog.Newlog("info")
	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		log.Error("这是一条erroo日志")
		time.Sleep(time.Second * 2)
	}

}
