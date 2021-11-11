package main

import (
	"fmt"
)

func main() {
	var numbers = make([]int ,3,5)	//定义一个长度为3，容量为5的切片

	fmt.Printf("len=%d, cap=%d, slice=%v\n",len(numbers),cap(numbers),numbers)

	//向numbers中追加一个元素，长度加1，容量不变
	numbers=append(numbers,1)
	fmt.Printf("len=%d, cap=%d, slice=%v\n",len(numbers),cap(numbers),numbers)

	//继续追加一个元素
	numbers=append(numbers,2)
	fmt.Printf("len=%d, cap=%d, slice=%v\n",len(numbers),cap(numbers),numbers)

	//继续追加，超过cap后，系统将会开辟一段和cap一样大的空间
	numbers=append(numbers,3)
	fmt.Printf("len=%d, cap=%d, slice=%v\n",len(numbers),cap(numbers),numbers)

}
