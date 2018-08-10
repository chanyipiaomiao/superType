package superType

import "strings"

type String string

func (ss String) String() string {
	return string(ss)
}
func (ss String) Length() int {
	return len(ss.String())
}
func (ss String) Index(sep string) int {
	return strings.Index(ss.String(), sep)
}
func (ss String) Split(sep string) []string {
	return strings.Split(ss.String(), sep)
}
func (ss String) Trim(flag ...string) string {
	var sub = " "
	if len(flag) > 0 {
		sub = flag[0]
	}
	return strings.Trim(ss.String(), sub)
}
func (ss String) Replace(old, new string, n ...int) string {
	tmp := -1
	if len(n)>0 {
		tmp = n[0]
	}
	return strings.Replace(ss.String(), old, new, tmp)
}
