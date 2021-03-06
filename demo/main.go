package main

import (
	"github.com/gohouse/superType"
	"fmt"
)

func main()  {
	testArray()
}

func testArray()  {
	var arr superType.MapSlice = []map[string]interface{}{{"a":1,"b":2},{"a":11,"b":22}}
	fmt.Println(arr.Column("b"))
}

func testString()  {

	var str superType.String = "Hello, Fizzday.net  "
	//var str superType.String = "你好吗, 我很好,你 呢? "
	// or
	//var str = superType.String(" Hel lo, Fizzday.net  ")

	fmt.Println(str.Split(","))  // result : [Hel lo  Fizzday.net  ]
	println(str.Length())  // result : 21
	println(str.Index(","))  // result : 6
	println(str.Contains("Fizz"))  // result : true
	println(str.StartWith("He"))  // result : true
	println(str.TrimSpace())  // result : Hel lo, Fizzday.net
	println(str.STrimSpace().Replace(" ",""))  // result : Hello, Fizzday.net
	println(str.EndWith("net"))  // result : false
	println(str.STrimSpace().EndWith("net"))  // result : true
	println(str.Replace("net", "com"))  // result : Hel lo, Fizzday.com
	println(str.Repeat(3))  // result : Hel lo, Fizzday.net  Hel lo, Fizzday.net  Hel lo, Fizzday.net
	println(str.ToUpper())  // result : HEL LO, FIZZDAY.NET
	println(str.ToLower())  // result : hel lo, fizzday.net
	println(str.Trim())  // result : Hel lo, Fizzday.net
	println(str.Trim("H"))  // result : el lo, Fizzday.net
	println(str.STrim().Trim())  // result : Hel lo, Fizzday.net
	println(str.TrimLeft("H"))  // result : el lo, Fizzday.net
	println(str.STrimSpace().TrimRight("et"))  // result : Hel lo, Fizzday.n
	println(str.SReplace("e","a").SReplace("a","e").Replace("net", "com"))  // result : Hel lo, Fizzdey.com
	println(str.STrimSpace().SReplace(" ","").SReplace("net","net.cn").
		STrimRight(".cn").String()) // result : Hello,Fizzday.net
}