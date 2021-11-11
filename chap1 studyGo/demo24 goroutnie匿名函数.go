package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {	//创建匿名函数并且执行
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()	//退出goroutine,return只是退出函数
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {
		time.Sleep(1*time.Second)
	}
}
