package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) { 
	for key, value := range cityMap {
		fmt.Println(key, value)
	}
}

func main() {
	cityMap := make(map[string]string)
	
	//添加
	cityMap["China"] = "北京"
	cityMap["Japan"] = "东京"
	cityMap["USA"] = "纽约"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap,"China")

	fmt.Println("-----------------")
	//拷贝
	cityMap2 := make(map[string]string, len(cityMap))
	for key, value := range cityMap {
		cityMap2[key] = value
	}
	printMap(cityMap2)
}