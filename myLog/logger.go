package myLog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

///定义一个 自定义的uint16	类型的方法。用于自己记录
type LogLevel uint16

//定义常量，定义log级别，用于进行大小比较，来达到开关的目的。
const (
	UNKONWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//转换log的标志符，传入一个string类型的s  返回一个loglevel类型，和一个错误
//
func parseLoglevel(s string) (LogLevel, error) {
	s = strings.ToLower(s) //
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKONWN, err
	}
}

//另外一个转换log字符串的方法，反转为一个字符串类型，然后打印出来。
func getLogString(lv LogLevel) string {
	switch lv {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "UNKONWN"
}

//该方法用于获取出发log日志的文件名称，方法名，行号。
//传进来一个n参数， 返回文件包名，方法名，和行号
func getInfo(n int) (funcName, fileName string, linnb int) {
	//runtime这个内置包中的caleer方法。获取对应的信息。
	pc, file, linnb, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.caller() failed \n")
		return
	}
	//获取到包名和方法名，return出去。
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
