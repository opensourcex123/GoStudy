package main

import "fmt"

func foo1(a string , b int ) int{
	fmt.Println("a=",a,", b=",b)
	c:=100
	return c
}
//返回多个返回值，匿名的
func foo2(a string , b int ) (int ,int){
	fmt.Println("a=",a,", b=",b)
	return 666,777
}
//返回多个返回值，有形参名称的
func foo3(a string ,b int)(r1 int ,r2 int){ //也可以写成(r1,r2 int)
	fmt.Println("a=",a,", b=",b)

	r1=1000
	r2=100
	return
}
func main(){
	a:="abcd"
	b:=10
	c:=foo1(a,b)
	fmt.Println("c=",c)
	fmt.Println("--------------")
	d,e:=foo2(a,b)
	fmt.Println("d=",d,", e=",e)
	fmt.Println("--------------")
	f,g:=foo3(a,b)
	fmt.Println("f=",f,", g=",g)
}
