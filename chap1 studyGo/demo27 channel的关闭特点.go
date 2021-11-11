package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i<5; i++ {
			c <- i
		}
		//close关闭一个channel
		close(c)
	}()

	for {
		//ok若为true，说明channel没有关闭，false则已关闭
		if data,ok := <-c; ok{
			fmt.Println(data)
		} else{
			break
		}
	}
	fmt.Println("main finish")
}
