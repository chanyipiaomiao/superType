# superType
golang类型扩充 , 超级类型, 面向对象用法. 这里只起到一个抛砖引玉的作用, 欢迎广大面向对象爱好者一起扩充

## 用法示例
```go
func main()  {
	var ss superType.String = " asdf,sdfsa"

	a:= ss.Split(",")
	b:= ss.Length()
	c:= ss.String()
	d:= ss.Index(",")
	e:= ss.Trim()
	f:= ss.Trim("a")
	g:= ss.Replace("a", "m", -1)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
```
结果
```sh
[ asdf sdfsa]
11
 asdf,sdfsa
5
asdf,sdfsa
 asdf,sdfs
 msdf,sdfsm
```

