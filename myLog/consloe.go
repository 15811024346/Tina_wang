package myLog

import "fmt"

//创建一个log类
type Logger struct {
}

//Newlog 是Logger的方法
func Newlog() Logger {
	return Logger{}
}
func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}
func (l Logger) Info(msg string) {
	fmt.Println(msg)
}
func (l Logger) Warning(msg string) {
	fmt.Println(msg)
}
func (l Logger) Error(msg string) {
	fmt.Println(msg)
}
func (l Logger) fatal(msg string) {
	fmt.Println(msg)
}
