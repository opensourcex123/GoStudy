package main

import "fmt"

func main() {
	defer fmt.Println("main end1")	//defer关键字类似于c++中的析构函数，在函数最后执行
	defer fmt.Println("main end2")	//defer的执行顺序是栈的形式，先进后出
	//defer和return共存时，先执行return，再执行defer
	fmt.Println("main:hello go 1")
	fmt.Println("main:hello go 2")
}
