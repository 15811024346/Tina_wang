package main

import (
	"fmt"
	"io"
	"os"
)

//写一个从文件中间插入内容的方法

func f1(){
	//打开要操作的文件
	fileObj,err := os.OpenFile("./test1.txt",os.O_RDWR,0644)
	if err != nil{
		fmt.Printf("open file failed err :%v",err)
		return
	}
	//因为没有办法直接插入写入源文件，所以需要借助一个tem临时文件来操作
	tmpFile,err := os.OpenFile("./tmp.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0644)
	if err !=nil{
		fmt.Printf("creart file faild err:%v",err)
		return
	}
	defer tmpFile.Close()//用defer语句，在方法用完的时候关闭文件
	defer fileObj.Close()
	//读取源文件，写入到临时文件
	var ret [1]byte	//定义一个数组
	n,err := fileObj.Read(ret[:])
	if err!=nil{
		fmt.Printf("read file failed err:%v",err)
		return
	}
	tmpFile.Write(ret[:n])
	//再写入要插入的内容
	var s []byte
	s=[]byte{'a'}
	tmpFile.Write(s)
	//紧接着把源文件内容写入到临时文件中
	var x[1024]byte
	for{
		n,err:= fileObj.Read(x[:])	//把源文件读入到一个切片中
		if err ==io.EOF{
			tmpFile.Read(x[:n])	//，如果读完了。把读取到的内容，写入到临时文件中去，break跳出循环
			break
		}
		if err!= nil{
			fmt.Printf("read from file faild err :%v\n",err)
			return
		}
		tmpFile.Write(x[:n])
	}
	//全部写完以后，重命名文件名
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./tmp.txt","./test1.txt")


}
func main(){
	f1()
}
