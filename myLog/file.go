package myLog

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往其他位置写入日志
//定义一个 文件写入的结构体。。
type FileLogger struct {
	level       LogLevel //调用之前定义的loglevel类型
	filePath    string   //日志保存的路径
	filename    string   //日志报春的文件名字
	fileobj     *os.File //定义指针类型的文件。（因为可能会大。直接返回指针会比较节省资源。）
	errfileobj  *os.File
	maxFilesize int64 //定义最大log文件的大小。
}

//写一个构造函数
//定义一个newlog输入到文件的函数， 传入 log界别，文件地址，文件名称，文件最大值的参数。返回一个结构体指针。
func NewFileLogger(levelstr, fp, fn string, mf int64) *FileLogger {
	loglevel, err := parseLoglevel(levelstr)
	if err != nil {
		panic(err)
	}
	//获取到输入的log文件的信息，传入到一个变量中。然后打开log文件。
	fl := &FileLogger{
		level:       loglevel,
		filePath:    fp,
		filename:    fn,
		maxFilesize: mf,
	}
	//调用initfile方法，
	err = fl.initFile() //按照文件路径和名字将文件打开
	//             1111111111111111111111111111111111111111---------------------对应下边的哪个nil。。---------------------
	//正常打开文件以后，error返回的就是个nil。
	if err != nil {
		panic(err)
	}
	return fl

}

//定一个初始化文件的方法，
func (f *FileLogger) initFile() error {
	//使用os.join拼接log文件的路径，和文件名称。
	fullfilename := path.Join(f.filePath, f.filename)
	//打开和创建一个文件按照上边拼接的路径来。第二个是标志位。第三个是文件权限。
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
	//日志文件都已经正常打开
	f.fileobj = fileobj
	f.errfileobj = errfileobj

	//因为这个函数返回的是个error错误，所以打开好文件以后，就返回一个nil    1111111111111111111和上边的1对应----------------------
	return nil
}

////Newlog 是初始化log级别的方法，传入一个字符串类型，返回 一个自定义的loglevel类型
func (f *FileLogger) unable(loglevel LogLevel) bool {
	return loglevel >= f.level
}

//该方法传入一个 文件类型的指针，返回布尔值  用于检测文件大小的方法。。   传入一个文件，返回布尔值。
func (f *FileLogger) checkSize(file *os.File) bool {
	//获取文件的信息， file。stat 方法。
	fileInfo, err := file.Stat()
	//如果有报错，则返回false
	if err != nil {
		fmt.Printf("get file info failed err:%v", err)
		return false
	}
	//用取到的文件信息，里的size对比 前边设置好的文件最大值。如果大于等于就返回true   ----这个f值，获取的是main函数中传进来的最大值。。
	return fileInfo.Size() >= f.maxFilesize
}
func (f FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	//1，关闭当前文件
	file.Close()
	//2，备份当前日志文件

	now1 := time.Now().Format("2006-01-02 15:04:05")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get fileInof 11111111111111111failed err :%v \n", err)
		return nil, err
	}
	logname := path.Join(f.filePath, fileInfo.Name())    //那到当前log文件的完整路径
	newLogname := fmt.Sprintf("%s.bak%s", logname, now1) //拼接一个日志文件备份的名字。
	//重命名新的log日志文件名
	os.Rename(logname, newLogname)
	//3，打开一个新的日志文件
	fileobj, err := os.OpenFile(logname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open new file failed err:%v \n", err)
		return nil, err
	}
	//将新打开的fileobj复制给f.fileobj
	return fileobj, nil
}

//定义一个以consloeLogger 为接收者的log函数 返回 log等级，和getinfo方法中获取到的log信息，//方法名，调用log的函数，行号等信息。并以一个接口类型返回。
//接口类型的返回值，可以是任意类型，格式的。比较方便，传回给我什么，我就返回什么。
func (f FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.unable(lv) {
		//这个msg变量保存的是log中的所有信息。
		msg := fmt.Sprintf(format, a...)
		//时间变量。当前时间，用于拼接
		now := time.Now()
		//初始化三个变量，用来调用getinfo方法。3 是调用当前这个函数的层级
		fileName, funcNmae, linnb := getInfo(3)
		if f.checkSize(f.fileobj) {
			newfile, err := f.splitFile(f.fileobj)
			if err != nil {
				return
			}
			f.fileobj = newfile
		}
		//格式化输入对应的log日志信息。。
		fmt.Fprintf(f.fileobj, "[%s] [%s] [%s--%s--%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcNmae, linnb, msg)
		if lv >= ERROR {
			if f.checkSize(f.errfileobj) {
				newerrfile, err := f.splitFile(f.errfileobj)
				if err != nil {
					return
				}
				f.errfileobj = newerrfile
			}
		}
		//如果要记录的日志级别大于error，再在errorfile中再记录一遍。
		fmt.Fprintf(f.errfileobj, "[%s] [%s] [%s--%s--%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcNmae, linnb, msg)
	}

}

//一下是各种级别的log定义方法。 。。
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

//关闭两个日志文件。
func (f *FileLogger) close() {
	f.errfileobj.Close()
	f.fileobj.Close()
}
