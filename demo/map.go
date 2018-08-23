package main

import (
	"github.com/gohouse/superType"
	"fmt"
)

func main() {
	var arr = []map[string]interface{}{
		{
			"a": "a1",
			"b": 2,
		},
		{
			"a": "a2",
			"b": 3,
		},
	}
	brr := superType.MapSlice(arr)

	fmt.Println(brr.Column("a"))

	var crr = superType.Map{
		"a":2,
		"b":3,
	}
	fmt.Println(crr.Keys())
	fmt.Println(crr.Values())
}
