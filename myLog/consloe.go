package myLog

import (
	"fmt"
	"time"
)

//定义一个unit16类型的logleve级别

//创建一个log类
type ConsloeLogger struct {
	Level LogLevel
}

//Newlog 是Logger的方法
func Newlog(levelstr string) ConsloeLogger {
	level, err := parseLoglevel(levelstr)
	if err != nil {
		panic(err)
	}
	return ConsloeLogger{
		Level: level,
	}
}
func (c ConsloeLogger) unable(loglevel LogLevel) bool {
	return loglevel >= c.Level
}
func (c ConsloeLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.unable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		fileName, funcNmae, linnb := getInfo(3)
		fmt.Printf("[%s] [%s] [%s--%s--%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcNmae, linnb, msg)
	}
}
func (c ConsloeLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}
func (c ConsloeLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}
func (c ConsloeLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}
func (c ConsloeLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
func (c ConsloeLogger) fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
