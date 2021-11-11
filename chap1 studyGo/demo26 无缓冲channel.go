package main

import "fmt"

func main() {
	c := make(chan int,3)	//创建一个缓冲为3的channel

	go func() {
		defer fmt.Println("sub goroutine over")

		for i :=0; i<4; i++ {
			c <- i
			fmt.Println("sub routine is running,the element is",i)
		}
	}()

	for i := 0; i<3; i++ {
		num := <-c
		fmt.Println("num is",num)
	}

	fmt.Println("main over")
}
