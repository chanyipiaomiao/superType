package superType

import (
	"strings"
)

type Stringer interface {
	String() string
	Length() int
	Index(sep string) int
	Split(sep string) []string
	Contains(substr string) bool
	ToUpper() string
	ToLower() string
	ToTitle() string

	StartWith(substr string) bool
	EndWith(substr string) bool

	// 以下方法, 在方法前加 S 的, 都是自返回, 可链式多次调用的
	Replace(old, new string, n ...int) string
	SReplace(old, new string, n ...int) String

	Repeat(count int) string
	SRepeat(count int) String

	Trim(flag ...string) string
	STrim(flag ...string) String

	TrimLeft(cutset string) string
	STrimLeft(cutset string) String

	TrimRight(cutset string) string
	STrimRight(cutset string) String
}
type String string

func (ss String) String() string {
	return string(ss)
}
func (ss String) SString() String {
	nameRune := []rune(ss.String())
	return String(nameRune)
}
func (ss String) Length() int {
	nameRune := []rune(ss.String())
	return len(nameRune)
}
func (ss String) Index(sep string) int {
	return strings.Index(ss.String(), sep)
}
func (ss String) Split(sep string) []string {
	return strings.Split(ss.String(), sep)
}

func (ss String) Contains(substr string) bool {
	return strings.Contains(ss.String(), substr)
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

func (ss String) StartWith(substr string) bool {
	strlen := String(substr).Length()
	if strlen<=ss.Length() && ss.String()[:strlen]==substr {
		return true
	}
	return false
}
func (ss String) EndWith(substr string) bool {
	strlen := String(substr).Length()
	sslen := ss.Length()
	if strlen<=ss.Length() && ss.String()[sslen-strlen:sslen]==substr {
		return true
	}
	return false
}

// 拥有自返回方法的函数

func (ss String) Replace(old, new string, n ...int) string {
	tmp := -1
	if len(n) > 0 {
		tmp = n[0]
	}
	return strings.Replace(ss.String(), old, new, tmp)
}
func (ss String) SReplace(old, new string, n ...int) String {
	return String(ss.Replace(old,new,n...))
}

func (ss String) Repeat(count int) string {
	return strings.Repeat(ss.String(), count)
}
func (ss String) SRepeat(count int) String {
	return String(ss.Repeat(count))
}

func (ss String) Trim(cutset ...string) string {
	var sub = " "
	if len(cutset) > 0 {
		sub = cutset[0]
	}
	return strings.Trim(ss.String(), sub)
}
func (ss String) STrim(cutset ...string) String {
	return String(ss.Trim(cutset...))
}

func (ss String) TrimLeft(cutset ...string) string {
	var sub = " "
	if len(cutset) > 0 {
		sub = cutset[0]
	}
	return strings.TrimLeft(ss.String(), sub)
}
func (ss String) STrimLeft(cutset string) String {
	return String(ss.TrimLeft(cutset))
}

func (ss String) TrimRight(cutset ...string) string {
	var sub = " "
	if len(cutset) > 0 {
		sub = cutset[0]
	}
	return strings.TrimRight(ss.String(), sub)
}
func (ss String) STrimRight(cutset string) String {
	return String(ss.TrimRight(cutset))
}

func (ss String) TrimSpace() string {
	return strings.TrimSpace(ss.String())
}
func (ss String) STrimSpace() String {
	return String(ss.TrimSpace())
}
