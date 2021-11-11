package main

import "fmt"

func main() {
	s:=[]int {1,2,3}
	//左闭右开，和python一样
	s1:=s[0:2]
	fmt.Println(s1)

	s1[0]=100
	fmt.Println(s)
	fmt.Println(s1)
	//copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int,3)
	//将s中的值拷贝到s2中
	copy(s2,s)
	fmt.Println(s2)
}