package main

import "fmt"

type Human struct{
	name string
	sex string
}

func (this *Human) Eat() {
	fmt.Println("Human is eating...")
}

func (this *Human) Walk() {
	fmt.Println("Human is walking...")
}

type Teacher struct {
	Human	//代表该类是Human类的子类
	age int
}
//重写了父类的方法
func (this *Teacher) Eat() {
	fmt.Println("Teacher is eating...")
}
//子类的新方法
func (this *Teacher) Teach() {
	fmt.Println("Teacher is teaching...")
}
func main() {
	h := Human{"zhangsan","man"}
	h.Eat()
	h.Walk()

	//t := Teacher{Human{"lisi","woman"},10}	//声明子类的方式
	var t Teacher
	t.name="lisi"
	t.sex="woman"
	t.age=10
	t.Eat()	//调用子类重写的方法
	t.Teach()	//调用子类新添加的方法
}
