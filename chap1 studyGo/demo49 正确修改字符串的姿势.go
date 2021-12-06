package main

import "fmt"

func main() {
	s := "hello"
	sRune := []rune(s)
	sRune[0] = '我'
	s = string(sRune)
	fmt.Println(s)
}
