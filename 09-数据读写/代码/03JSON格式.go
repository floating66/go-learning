package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct { 
	Title string `json:"title"`
	Year int `json:"year"`
	Price int `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"张学友", "周杰伦", "李连杰"}}

	//将结构体变量movie编码为json字符串(编码)
	jsonStr,err := json.Marshal(movie)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("json string is: %s", jsonStr)

	//将json字符串解码为结构体变量movie(解码)
	mymovie := Movie{}
	mymovie := json.Unmarshal(jsonStr, &mymovie)  //&mymovie表示接收的结构体变量
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("movie is: %v", mymovie)
}