package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type is",reflect.TypeOf(arg))
	fmt.Println("value is",reflect.ValueOf(arg))
}

func main() {
	var a float64=1.2345
	reflectNum(a)
}
