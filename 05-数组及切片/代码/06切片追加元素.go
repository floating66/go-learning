package main

import "fmt"

func main() {
	var num = make([]int, 3, 5)
	//len是数组长度，cap是数组容量，slice是数组切片
	fmt.Println("len = ", len(num), "cap = ", cap(num),"slice = ", num)

	//追加元素
	num = append(num, 1)
}