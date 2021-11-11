package main

import "fmt"

var gA int =100

func main(){	//前两种方法可以声明全局变量，第三种不可以
	//方法一：声明一个变量，默认初始值是0
	var a int = 10
	fmt.Println(a)

	//方法二，省略类型，根据赋值判断数据类型
	var b=100
	fmt.Println(b)
	fmt.Printf("type of b = %T\n",b)

	//方法三，（常用的方法）省略var，直接自动匹配
	c:=100
	fmt.Printf("c=%d,type of c = %T\n",c,c)

	//全局变量
	fmt.Printf("gA=%d,type of gA = %T\n",gA,gA)

	//声明多个变量
	var d,e int =100,200
	fmt.Println("d=",d,", e=",e)
	var f,g = 100,200
	fmt.Println("f=",f,", g=",g)
	//多行变量声明
	var(
		h int =100
		j string="abcd"
	)
	fmt.Println("h=",h,", j=",j)

}
