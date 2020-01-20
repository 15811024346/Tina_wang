package main

import (
	"../db_function1"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := db_function1.InitDb()
	if err != nil {
		fmt.Printf("init DB failed err:%v\n", err)
		return
	}
	fmt.Println("链接成功")
	//db_function1.QueryRow()
	//多条查询。
	db_function1.QueryMore(1)

}
