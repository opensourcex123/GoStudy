package main

import "fmt"
//所有的类型都实现于interface{}类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	//判断arg的底层数据类型，类型断言机制
	value ,ok := arg.(string)
	if !ok{
		fmt.Println("arg is not string")
		fmt.Println("------------")
	} else {
		fmt.Println("arg is string,value is",value)
		fmt.Println("------------")
	}
}

type Book1 struct {
	name string
}

func main() {
	book := Book1{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
}
