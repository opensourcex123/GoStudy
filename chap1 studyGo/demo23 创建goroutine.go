package main

import (
	"fmt"
	"time"
)
//子goroutnie
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine : i=%d\n",i)
		time.Sleep(1*time.Second)
	}
}
//主goroutine,当主线程结束时，子线程就会随之消失
func main() {
	//创建一个goroutine 去执行newTask流程
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine : i=%d\n",i)
		time.Sleep(1*time.Second)
	}
}
