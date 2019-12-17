package myLog

import (
	"fmt"
	"time"
)
//定义一个unit16类型的logleve级别

//创建一个log类
type Logger struct {
	Level LogLevel
}

//Newlog 是Logger的方法
func Newlog(levelstr string) Logger {
	level,err := parseLoglevel(levelstr)
	if err!= nil{
		panic(err)
	}
	return Logger{
		Level:level,
	}
}
func (l Logger) unable(loglevel LogLevel) bool{
	return  loglevel >= l.Level
}
func log(lv LogLevel,msg string){
	now := time.Now()
	fileName , funcNmae,linnb := getInfo(3)
	fmt.Printf("[%s] [%s] [%s--%s--%d] %s\n",now.Format("2006-01-02 15:04:05"),getLogString(lv),fileName,funcNmae,linnb,msg)
}
func (l Logger) Debug(msg string) {
	if l.unable(DEBUG){
		log(DEBUG,msg)
	}

}
func (l Logger) Info(msg string) {
	if l.unable(INFO){
		log(INFO,msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.unable(WARNING){
		log(WARNING,msg)
	}
}
func (l Logger) Error(msg string) {
	if l.unable(ERROR){
		log(ERROR,msg)
	}
}
func (l Logger) fatal(msg string) {
	if l.unable(FATAL){
		log(FATAL,msg)
	}
}
