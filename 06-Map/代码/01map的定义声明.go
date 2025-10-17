package main

import "fmt"

func main() {
	//声明map变量，key是string，value是string
	var myMap map[string]string

	//使用map前需要分配空间
	myMap = make(map[string]string, 10)

	myMap["one"] = "java"
	myMap["two"] = "python"
	myMap["three"] = "golang"

	fmt.Println(myMap)

	//第二种方式
	myMap2 := map[string]string{
		"one": "java",
		"two": "python",
		"three": "golang",
	}
	fmt.Println(myMap2)	

}