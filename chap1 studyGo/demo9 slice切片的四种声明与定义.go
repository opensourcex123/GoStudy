package main

import "fmt"

func main() {
	//第一种方法
	//slice1 := []int {1,2,3}

	//第二种方法
	//var slice1 []int
	//slice1 = make([]int ,3)

	//第三种方法
	//var slice1 []int =make([]int ,3)

	//第四种方法,常用
	slice1 := make([]int, 3)
	fmt.Printf("len=%d,slice1=%v\n", len(slice1), slice1)

	//判断一个切片是否为空
	if slice1 == nil {
		fmt.Println("slice1为空")
	} else {	//else必须和大括号在同一行
		fmt.Println("slice1不为空")
	}
}
