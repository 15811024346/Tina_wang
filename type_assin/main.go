package main

import "fmt"

//类型断言
//类型断言是为了了解空接口中接收的值是什么类型。。
func assign(a interface{}){
	fmt.Printf("%T\n",a)
	str,ok := a.(string)
	if !ok{
		fmt.Println("猜错了")
	}else {
		fmt.Printf("传进来的是个字符串\n",str)
	}
}
func assign2(a interface{}){
	fmt.Printf("%T\n",a)
	switch t:=a.(type) {
	case string:
		fmt.Println("是一个字符串",t)
	case int64:
		fmt.Println("是一个int64",t)
	case bool:
		fmt.Println("是一个bool",t)
	case int:
		fmt.Println("是一个int",t)
	}
}

func main(){
	assign("哈哈哈" )
	assign2("嘿嘿嘿")
	assign2(true)
	assign2(1111)
}
