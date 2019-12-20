package main

import "fmt"

func maoPao(s []int) {

	var temp int
	for i := 0; i < len(s); i++ {
		//fmt.Printf("第%v遍循环", i)
		//println("a[i]是：", a[i], "a[i]是")
		for j := 0; j < len(s)-1-i; j++ {
			//fmt.Printf("=====内部第%v遍循环=====", j)
			//fmt.Println("-----------------")
			if s[j] > s[j+1] {
				//fmt.Println("------a[j]----", a[j], a[j+1], "-----a[j+1]-----")
				temp = s[j]
				//fmt.Println("------temp----", temp, a[j], "-----a[j]-----")
				s[j] = s[j+1]
				//fmt.Println("-----a[j]-----", a[j], a[j+1], "-----a[j+1]-----")
				s[j+1] = temp
				//fmt.Println("-----temp-----", temp, a[j+1], "-----a[j]-----")
			}
			//println(a[j])
		}
	}
	fmt.Println(s)
}
func main() {
	a := []int{2, 142, 51, 65, 7, 22, 33}
	maoPao(a)
}
