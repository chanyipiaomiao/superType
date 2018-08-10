package superType

import (
	"testing"
	"fmt"
)

func TestSuperString_All(test *testing.T) {
	var ss String = " asdf,sdfsa"

	a:= ss.Split(",")
	b:= ss.Length()
	//c:= ss.String()
	d:= ss.Index(",")
	//e:= ss.Trim()
	f:= ss.Trim("a")
	g:= ss.Replace("a", "m", -1)

	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//fmt.Println(e)
	//fmt.Println(f)
	//fmt.Println(g)

	if a[0] != " asdf" || a[1]!="sdfsa" {
		test.Error("FAIL: %s failed", "Split")
		return
	}
	if b!=11 {
		test.Error("FAIL: %s failed", "Length")
		return
	}
	if d!=5 {
		test.Error("FAIL: %s failed", "Index")
		return
	}
	if f!=" asdf,sdfs" {
		test.Error("FAIL: %s failed", "Trim")
		return
	}
	if g!=" msdf,sdfsm" {
		test.Error("FAIL: %s failed", "Replace")
		return
	}
	test.Log(fmt.Sprintf("PASS"))
}
