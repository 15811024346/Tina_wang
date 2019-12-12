package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func openFile1(){
	//打开文件
	fileojb,err :=os.Open("./test1.txt" )
	if err!= nil{
		fmt.Printf("open file filed err:%d",err)
		return
	}
	//记得关闭文件
	defer fileojb.Close()
	//读文件
	var tmp = make([]byte, 128)	//指定文件长度，可以动态扩容
	for {
		n,err :=fileojb.Read(tmp[:])
		if err != nil{
			fmt.Print("open file filed err:%d",err)
			return
		}
		fmt.Printf("读取了%d个字节\n",n)
		fmt.Println(string(tmp[:n]))
		if n<128{
			return
		}
	}
}

func openFile2(){
	//打开文件
	fileojb,err :=os.Open("./test1.txt" )
	if err!= nil{
		fmt.Printf("open file filed err:%d",err)
		return
	}
	//记得关闭文件
	defer fileojb.Close()
	reder1 := bufio.NewReader(fileojb)
	for{

		lines,err :=reder1.ReadString('\n')
		if err == io.EOF{
			return
		}
		if err!=nil{
			fmt.Printf("文件读取出错：%v",err)
			return
		}
		fmt.Print(lines)
	}
}

func main(){
	openFile2()
}
