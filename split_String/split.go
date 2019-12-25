package split_String

import "strings"

//分割字符串\

//str = "abcdefg"   rep ="c"    [ab,defg]
//Split_Str ...

func Split_Str1(str string, rep string) []string {
	var str1 []string
	index := strings.Index(str, rep)
	for index >= 0 {
		str1 = append(str1, str[:index])
		str = str[index+1:]
		index = strings.Index(str, rep)
	}
	str1 = append(str1, str)
	return str1
}
