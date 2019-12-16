package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main(){


	fileobj,err :=os.OpenFile("./test_waibu/src/logXuexiTest/logTest1.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Printf("open file failed err :%v",err)
		return
	}
	log.SetOutput(fileobj)
	for{
		log.Println("这是一条测试日志///")
		time.Sleep(time.Second*3)
	}
}
