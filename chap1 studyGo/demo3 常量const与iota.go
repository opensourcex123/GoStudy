package main

import "fmt"

//const定义枚举类型
const(
	//可以在const()里添加iota关键字，每行的iota都会累加，第一行的默认值是0
	//iota只能在const()里使用
	BEIJING=10*iota
	SHANGHAI
	SHENZHEN
)
func main(){
	//把var改为const就是常量
	const len int = 100
	fmt.Println("len=",len)

	fmt.Println("BEIJING=",BEIJING)
	fmt.Println("SHANGHAI=",SHANGHAI)
	fmt.Println("SHENZHEN=",SHENZHEN)
}