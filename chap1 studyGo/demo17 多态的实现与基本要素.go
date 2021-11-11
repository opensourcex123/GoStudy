package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {	//当这个类实现了接口的所有方法后，就默认该类继承了接口
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping...")
}

func (this *Cat) GetColor() string{
	return this.color
}

func (this *Cat) GetType() string{
	return "Cat"
}

func main() {
	var animal AnimalIF	//接口相当于指针，指向某个具体的类
	animal=&Cat{"yellow"}
	animal.Sleep()	//调用了cat类的方法，触发了多态的现象
}