package main

import "fmt"
//如果类名首字母大写，表示其他包也能够访问
type Hero struct{
	//如果说类的属性首字母大写，表示该属性是对外能够访问的，否则只能够类的内部访问
	name string
	level int
}

func (this *Hero) show() {
	fmt.Println("name=",this.name)
	fmt.Println("level=",this.level)
}

func (this *Hero) getName() string{
	return this.name
}

func (this *Hero) setName(newName string) {
	//this是调用该方法的对象的一个副本，即拷贝
	this.name=newName
}

func main() {
	//对象的定义方式
	hero := Hero{name:"zhangsan",level: 78}

	hero.show()
	fmt.Println("---------")
	hero.setName("666")
	hero.show()
}
