package main

import "fmt"

func main() {

	/*
		练习题1：打印一个金字塔，可以接收一个整数表示层数。打印金字塔
		示例：
			 *
			***
		   *****
		  *******
		先打印一个矩形

			***
			***
			***
	*/
	a1(5)
}
func a1(a int) {
	for i := 1; i <= a; i++ {
		for k := i; k <= a; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == a {
				fmt.Print("*")
			} else {
				fmt.Print(" ")

			}

		}
		fmt.Println()
	}
}
