package main

import "fmt"

func printArr(myArr []int) {
	//切片是引用传递，数组是值传递
	//_表示匿名
	for _,value := range myArr { 
		fmt.Println("value:",value)
	} 

	myArr[0] = 100
}

func main() {
	//动态数组
	myArr := []int{1,2,3,4}

	printArr(myArr)
	
	fmt.Println("====================")
	for _,value := range myArr { 
		fmt.Println("value:",value)
	} 

}