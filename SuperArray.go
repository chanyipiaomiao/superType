package superType

import (
	"fmt"
	"strings"
)

type Arrayer interface {
	Merge(arr []interface{}) []interface{}
	Search(arr []interface{}) int

	Unique(arr []interface{}) []interface{}
	Append(arr interface{}) []interface{}
	Pop() interface{}                     // 弹出最后一个值,原数组值少一个
	Shift() interface{}                   // 弹出第一个值,原数组值少一个
	Sum() interface{}                     // 求和
	Diff(arr []interface{}) []interface{} // 差集
	Flip() []interface{}                  // 键值互换

	Map(arr func())
	Reuce(arr func())

	Join(sep string) string
}

type Array []interface{}

func (a Array) Append(arr interface{}) []interface{} {
	return append(([]interface{})(a), arr)
}

func (a *Array) SAppend(arr interface{}) *Array {
	*a = append(([]interface{})(*a), arr)
	return a
}

func (a Array) Join(sep string) string {
	t := make([]string, len(a))
	for i, v := range a {
		t[i] = fmt.Sprint(v)
	}
	return strings.Join(t, sep)
}
