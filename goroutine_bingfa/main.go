package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

////goroutin
func hello(i int) {
	fmt.Println("hello ", i)
}

//
func f1() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    //会随机出现int64位的随机数
		r2 := rand.Intn(10) //可以指定范围的显示
		fmt.Println(r1, r2)
	}
}
func f2(i int) {
	defer wg.Done()
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
func main() {
	runtime.GOMAXPROCS(6)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go f2(i)
	//}
	//wg.Wait()
	//f1()
	//for i := 0; i < 100; i++ {
	//	go hello(i)
	//}
	////开启一个单独的gotoutine 去执行hello（）这个函数
	//fmt.Println("world")
}
