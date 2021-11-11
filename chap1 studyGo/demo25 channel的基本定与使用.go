package main

import "fmt"

func main() {
	//定义一个channel，用于两个线程之间传递数据
	c := make(chan int)	//无缓冲的channel

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine is running...")

		c <- 666	//将666 发送给c
	}()

	num := <-c	//从c中接收数据，并赋值给num
	fmt.Println("nums is",num)
	fmt.Println("main over")
}
