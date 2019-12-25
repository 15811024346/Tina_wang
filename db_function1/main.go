package db_function1

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

type User22 struct {
	id   int
	name string
	age  int
}

func InitDb() (err error) {

	dsn := "root:wangshuai1@tcp(127.0.0.1:3306)/goMysql1"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}
func QueryRow() {
	var u1 User22
	sqlstr := `select id,name,age from user where id=?;`
	err := db.QueryRow(sqlstr, 2).Scan(&u1.id, &u1.name, &u1.age) ///取出来的是个指针类型的值。必须带上&符号
	if err != nil {
		return
	}
	fmt.Printf("sqlobj%#v\n", u1)

}
