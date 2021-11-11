package main

import (
	"fmt"
	"reflect"
)
//resume：简历
type resume struct {	//添加标签
	Name string `info:"name" doc:"我的名字"`
	Sex string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str) //得到所有的元素

	for i := 0; i< t.NumField(); i++{
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info is",taginfo)
		fmt.Println("doc is",tagdoc)
	}
}

func main() {
	var re resume

	findTag(re)
}


