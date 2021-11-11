package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case <-ch1:
		fmt.Println("case1可执行。。。。")
		case <-ch2:
			fmt.Println("case2可执行...")
			case <-time.After(3 * time.Second):	// 其他两个case因为没有数据在chan中，所以被阻塞
				fmt.Println("case3执行，timeout。。")
	}
}
