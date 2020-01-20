package main

import (
	"fmt"
	"math/rand"
	"time"
)

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		fmt.Println("rand :", rand.Intn(100))
	}
}
