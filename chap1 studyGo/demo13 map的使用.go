package main

import "fmt"

func printMap(cityMap map[string]string) {
	//是一种引用传递
	//遍历
	for key,value := range cityMap {
		fmt.Println("key is", key)
		fmt.Println("value is", value)
	}
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"]="Beijing"
	cityMap["Japan"]="Tokyo"
	cityMap["USA"]="NewYork"

	printMap(cityMap)

	//删除
	delete(cityMap,"Japan")

	//修改
	cityMap["USA"]="DC"

	fmt.Println("-----------")

	printMap(cityMap)
}
