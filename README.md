# superType
golang类型扩充 , 超级类型, 面向对象用法. 这里只起到一个抛砖引玉的作用, 欢迎广大面向对象爱好者一起扩充

## 用法示例
```go
package main

import (
	"github.com/gohouse/superType"
)

func main()  {
	var str superType.String = "Hel lo, Fizzday.net  "
	// or
	//var str = superType.String(" Hel lo, Fizzday.net  ")

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
	println(str.Split(","))  // result : [Hel lo  Fizzday.net  ]
	println(str.Trim())  // result : Hel lo, Fizzday.net
	println(str.Trim("H"))  // result : el lo, Fizzday.net
	println(str.STrim().Trim())  // result : Hel lo, Fizzday.net
	println(str.TrimLeft("H"))  // result : el lo, Fizzday.net
	println(str.STrimSpace().TrimRight("et"))  // result : Hel lo, Fizzday.n
	println(str.SReplace("e","a").SReplace("a","e").Replace("net", "com"))  // result : Hel lo, Fizzdey.com
	println(str.STrimSpace().SReplace(" ","").SReplace("net","net.cn").
		STrimRight(".cn").String()) // result : Hello,Fizzday.net
}
```

