package main

import "fmt"

func main() {
	//第一种声明方式
	//声明myMap1是一种map类型 key是int，值是string
	var myMap1 map[int] string
	if myMap1==nil{
		fmt.Println("myMap1 is null")
	}
	//开辟空间，容量为10，也是动态开辟，不够系统会自动增加10容量
	myMap1=make(map[int] string,10)
	myMap1[0]="java"
	myMap1[1]="c++"
	myMap1[2]="python"
	fmt.Println(myMap1)	//哈希排序，内部是乱序的


	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[0]="java"
	myMap2[1]="c++"
	myMap2[2]="python"
	fmt.Println(myMap2)	//哈希排序，内部是乱序的

	//第三种声明方式
	myMap3 := map[int]string{
		1:"java",
		2:"c++",
		3:"python",
		4:"golang",	//最后一行也要写逗号
	}
	fmt.Println(myMap3)
}
