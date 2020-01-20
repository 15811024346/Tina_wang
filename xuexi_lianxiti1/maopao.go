package main

import "fmt"

func aa(n []int) {
	var temp = 0
	for i := 0; i <= len(n)-1; i++ {
		for j := 0; j < len(n)-1-i; j++ {
			if n[j] > n[j+1] {
				temp = n[j]
				n[j] = n[j+1]
				n[j+1] = temp
			}
		}
	}
	fmt.Println(n)
}
func main() {
	var a2 = []int{3, 5, 22, 11, 33, 55}
	aa(a2)
}
