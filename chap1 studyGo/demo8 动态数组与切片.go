package main

import "fmt"

func printArray(myArray[] int) {
	//引用传递
	//下划线表示匿名
	for _,value := range myArray{
		fmt.Println("value=",value)
	}

	myArray[1]=100	//修改了数组中的值
}

func main() {
	myArray := [] int {1,2,3,4}	//动态数组，slice切片
	fmt.Printf("type of myArray is %T\n",myArray)
	printArray(myArray)

	for _,value:=range myArray{
		fmt.Println("value=",value)
	}
}