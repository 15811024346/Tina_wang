package myLog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

const(
	UNKONWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)
func parseLoglevel(s string)(LogLevel,error){
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		err := errors.New("无效的日志级别")
		return UNKONWN,err
	}
}
func getLogString(lv LogLevel)string{
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
func getInfo(n int)(funcName,fileName string ,linnb int){
	pc,file,linnb,ok:=runtime.Caller(n)
	if !ok{
		fmt.Printf("runtime.caller() failed \n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)

	return
}