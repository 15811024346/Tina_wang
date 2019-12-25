package split_String

import (
	"reflect"
	"testing"
)

//单元测试的例子。。 文件名必须是以test结尾。 func 方法名开头必须是Test 开头。参数是（t *testing)
func TestSplit_Str1(t *testing.T) {
	got := Split_Str1("absdfawa", "b")
	want := []string{"a", "sdfawa"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v,got:%v\n", got, want)
	}
}

func TestSplit_Str12(t *testing.T) {
	got := Split_Str1("a:b:s:d:f:a:wa", ":")
	want := []string{"a", "b", "s", "d", "f", "a", "wa"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v,got:%v\n", got, want)
	}
}
