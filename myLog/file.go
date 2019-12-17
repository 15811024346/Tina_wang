package myLog

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往其他位置写入日志

type FileLogger struct {
	level       LogLevel
	filePath    string //日志保存的路径
	filename    string //日志报春的文件名字
	fileobj     *os.File
	errfileobj  *os.File
	maxFilesize int64
}

//写一个构造函数

func NewFileLogger(levelstr, fp, fn string, mf int64) *FileLogger {
	loglevel, err := parseLoglevel(levelstr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		level:       loglevel,
		filePath:    fp,
		filename:    fn,
		maxFilesize: mf,
	}
	err = fl.initFile() //按照文件路径和名字将文件打开
	if err != nil {
		panic(err)
	}
	return fl

}
func (f *FileLogger) initFile() error {
	fullfilename := path.Join(f.filePath, f.filename)
	fileobj, err := os.OpenFile(fullfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed err:%v\n", err)
		return err
	}
	errfileobj, err := os.OpenFile(fullfilename+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed err:%v\n", err)
		return err
	}
	//日志文件已经打开
	f.fileobj = fileobj
	f.errfileobj = errfileobj
	return nil

}
func (n *FileLogger) unable(loglevel LogLevel) bool {
	return loglevel >= n.level
}
func (f FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.unable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		fileName, funcNmae, linnb := getInfo(3)
		fmt.Fprintf(f.fileobj, "[%s] [%s] [%s--%s--%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcNmae, linnb, msg)
		if lv >= ERROR {
			//如果要记录的日志级别大陆error，再在errorfile中再记录一遍。
			fmt.Fprintf(f.errfileobj, "[%s] [%s] [%s--%s--%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcNmae, linnb, msg)
		}

	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)

}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)

}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)

}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) close() {
	f.errfileobj.Close()
	f.fileobj.Close()
}
