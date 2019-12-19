package myLog

import (
	"fmt"
	"time"
)

//定义一个unit16类型的logleve级别

//创建一个log类，用来定义log的级别
type ConsloeLogger struct {
	Level LogLevel
}

//Newlog 是初始化log级别的方法，传入一个字符串类型，返回 一个自定义的loglevel类型
func Newlog(levelstr string) ConsloeLogger {
	level, err := parseLoglevel(levelstr)
	if err != nil {
		panic(err)
	}
	return ConsloeLogger{
		Level: level,
	}
}

//用来对比传过来的参数的大小比较。  传过来一个loglevel 对比 之前定义好的方法，只显示大于等于该级别的log
func (c ConsloeLogger) unable(loglevel LogLevel) bool {
	return loglevel >= c.Level
}

//定义一个以consloeLogger 为接收者的log函数 返回 log等级，和getinfo方法中获取到的log信息，//方法名，调用log的函数，行号等信息。并以一个接口类型返回。
//接口类型的返回值，可以是任意类型，格式的。比较方便，传回给我什么，我就返回什么。
func (c ConsloeLogger) log(lv LogLevel, format string, a ...interface{}) {
	//如果传入的 log等级
	if c.unable(lv) {
		msg := fmt.Sprintf(format, a...)
		//获取当前时间。
		now := time.Now()
		//通过getInfo方法获取到会触发log的层数。。。。这个3，需要后期再检查怎么换成参数。。。。
		fileName, funcNmae, linnb := getInfo(3)
		fmt.Printf("[%s] [%s] [%s--%s--%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcNmae, linnb, msg)
	}
}

//一下是各种级别的log定义方法。 。。
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
