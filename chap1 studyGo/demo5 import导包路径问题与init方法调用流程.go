package main

import (
	"fmt"
	_ "goProject/lib1"	//匿名导包，无法使用包里的接口，但能使用init方法
	mylib2 "goProject/lib2"	//别名导包
)

func main() {
	fmt.Println("111")
	//lib1.Lib1Test()
	mylib2.Lib2Test()
}