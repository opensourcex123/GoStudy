package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var _ = fmt.Printf // 用于调试，结束时删除。
var _ io.Reader    // 用于调试，结束时删除。

func main() {
	fd, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: use fd.
	_ = fd
}
