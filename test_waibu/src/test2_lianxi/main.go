package main

import "fmt"

func main(){
	var r = [...]int{1,0,23,22,11,43,2,5}
	fmt.Println(r)
	for _,v :=range r{
		fmt.Println(v)

	}
}
