package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

///mysqlconfig 的配置文件。
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadInit(filename string, data interface{}) (err error) {
	//	参数的校验， 传进来的参数，必须是指针类型。（因为要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") //创建一个错误
		return
	}

	//	传进来的参数，必须是结构体类型指针，（因为配置文件中各种键值对，需要赋值给结构提的字段——）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer") //创建一个错误
		return
	}

	//	读文件得到字节类型数据
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//string(b) //将文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	fmt.Println(lineSlice)
	var structName string
	//	一行一行的读数据。
	for idx, line := range lineSlice {
		//去掉字符串首尾的空格。
		line = strings.TrimSpace(line)
		//判断如果是空格的话，就跳过
		if len(line) == 0 {
			continue
		}
		//判断如果是注释的话就跳过。。
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//如果不是注释，就开始判断，如果是【】开头的，就是一个节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = errors.New("data param should be a struct pointer") //创建一个错误
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			//把这一行首尾的【】去掉，取到中间的内容，把中间的空格去掉，取出中间的内容。
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntex eroor", idx+1)
				return
			}
			for i := 0; i < t.Elem().NumField(); i++ {
				filed := t.Elem().Field(i)
				if sectionName == filed.Tag.Get("ini") {
					//说明找到了对应的结构体。把字段名记下来。
					structName = filed.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//如果不是【】开头，就是以 = 号分割的键值对
			//以=号分割，等号左边是key，右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error ", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//fmt.Println(value)
			//根据structname去data里取出对应的结构体信息
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				fmt.Printf("data中的%s应该是个结构体。", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			//
			for i := 0; i < sType.NumField(); i++ {
				filed := sType.Field(i) //tag是存储在类型信息中的
				fileType = filed
				if filed.Tag.Get("ini") == key {
					fieldName = filed.Name
					break
				}
			}
			if len(fieldName) == 0 {
				continue
			}
			fieldObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fieldObj.Type().Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line %d value type err ", idx+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line %d value type err ", idx+1)
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line %d value type err ", idx+1)
					return
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}
	return
}
func main() {
	var cfg Config
	//var x = new(int)		//测试 传入参数为一个int类型的指针，查看报错。
	//  ini打开文件的路径，如果和当前main函数同目录的话，需要。/后再加一个当前包名。上一层的话只输入./上上层加。。/
	err := loadInit("./ini_demo/config.ini", &cfg) //
	if err != nil {
		fmt.Printf("load ini file failed err : %v\n", err)
		return
	}

	fmt.Println(cfg)
}
