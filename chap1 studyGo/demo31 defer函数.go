package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b")) // 先执行trace函数，最后执行defer时才会进行求值trace(b)
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
