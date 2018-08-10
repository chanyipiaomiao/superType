package superType

import "strings"

type Stringer interface {
	String() string
	Length() int
	Index(sep string) int
	Split(sep string) []string
	Replace(old, new string, n ...int) string
	Contains(substr string) bool
	Repeat(count int) string
	ToUpper() string
	ToLower() string
	ToTitle() string
	Trim(flag ...string) string
	TrimLeft(cutset string) string
	TrimRight(cutset string) string
	StartWith(substr string) bool
	EndWith(substr string) bool
}
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
func (ss String) Replace(old, new string, n ...int) string {
	tmp := -1
	if len(n) > 0 {
		tmp = n[0]
	}
	return strings.Replace(ss.String(), old, new, tmp)
}
func (ss String) Contains(substr string) bool {
	return strings.Contains(ss.String(), substr)
}
func (ss String) Repeat(count int) string {
	return strings.Repeat(ss.String(), count)
}
func (ss String) ToUpper() string {
	return strings.ToUpper(ss.String())
}
func (ss String) ToLower() string {
	return strings.ToLower(ss.String())
}
func (ss String) ToTitle() string {
	return strings.ToTitle(ss.String())
}
func (ss String) Trim(flag ...string) string {
	var sub = " "
	if len(flag) > 0 {
		sub = flag[0]
	}
	return strings.Trim(ss.String(), sub)
}
func (ss String) TrimLeft(cutset string) string {
	return strings.TrimLeft(ss.String(), cutset)
}
func (ss String) TrimRight(cutset string) string {
	return strings.TrimRight(ss.String(), cutset)
}

func (ss String) StartWith(substr string) bool {
	strlen := len(substr)
	if strlen<=ss.Length() && ss.String()[:strlen]==substr {
		return true
	}
	return false
}
func (ss String) EndWith(substr string) bool {
	strlen := len(substr)
	if strlen<=ss.Length() && ss.String()[ss.Length()-strlen:strlen]==substr {
		return true
	}
	return false
}
