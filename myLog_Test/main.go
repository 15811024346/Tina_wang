package main

import (
	"../myLog"
	"fmt"
	"time"
)

func main() {
	fmt.Println("111111111")
	log := myLog.Newlog()
	for {
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		time.Sleep(time.Second * 2)
	}

}
