package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name string `json:"name"`
	Year int `json:"year"`
	Price int `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	mov := Movie{"喜剧之王",2000,10,[]string{"周星驰","张柏芝"}}

	//编码的过程 结构体->json
	jsonStr , err := json.Marshal(mov)
	if err != nil {
		fmt.Println("json marshal is error")
		return
	}

	fmt.Printf("jsonStr is %s\n",jsonStr)

	//解码的过程，json->结构体
	myMov := Movie{}
	err = json.Unmarshal(jsonStr,&myMov)
	if err != nil {
		fmt.Println("json unmarshal is error")
	}
	fmt.Printf("%v\n",myMov)
}
