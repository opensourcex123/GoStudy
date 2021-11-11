package main

import "fmt"

func sum(c,quit chan int) {
	x,y := 1,1

	for {
		select {
			case c <- x:	//如果c可写，执行
				x=y
				y=x+y
			case <-quit:	//如果c可读，执行
				fmt.Println("quit")
				return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i :=0; i<6; i++ {
			fmt.Println(<-c)
		}
		quit <-0
	}()
	sum(c,quit)
	fmt.Println("main finish")
}