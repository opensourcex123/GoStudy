package main

import "fmt"

//声明一种新的数据类型，实际上是int的别名
type myint int
//定义一个结构体
type Book struct{
	title string
	auth string
}

func changeBook(book *Book){
	book.auth="666"
}
func main() {
	//var a myint
	//fmt.Println("a= ",a)
	//fmt.Printf("type of a is %T",a)

	var book1 Book
	book1.title="Golang"
	book1.auth="zhangsan"
	fmt.Printf("%v\n",book1)

	changeBook(&book1)
	fmt.Printf("%v\n",book1)
}
