package main

import "fmt"

//同一个结构体可以实现多个接口
//接口嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}
type eater interface {
	eat(string)
}
//cat实现了mover和eater这两个接口
type cat struct {
	name string
	feet int8
}
func (c cat) eat(food string){
	fmt.Printf("猫吃%s～～～",food)
}
func (c cat)move(){
	fmt.Println("走猫步～～")
}
func main(){
	fmt.Println("hello world")

}