package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called...")
	fmt.Println(this)
}

func (this User) Show() {
	fmt.Println("aaa")
}

func main() {
	user := User{1001, "zhangsan", 18}
	user.Call()
	Reflect(user)
}

func Reflect(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is", inputType)

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is", inputValue)
	fmt.Println("---------------")
	//通过type获取input里面的字段
	//1.获取interface的reflect.type，通过type得到numfield，进行遍历
	//2.得到每个field，数据类型
	//3.通过field的Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println("field is", field.Name)
		fmt.Println("type is", field.Type)
		fmt.Println("value is", value)
		fmt.Println("---------------")
	}

	//通过type获取里面方法，只能获取非指针方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println("func name is",m.Name)
		fmt.Println("func type is",m.Type)
		fmt.Println("------------")
	}
}
