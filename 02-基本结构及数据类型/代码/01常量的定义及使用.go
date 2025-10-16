package main

import "fmt"

//const定义枚举
const (
	//第一行iota默认值为0，每行iota递增1  只在const定义的常量中起作用
	BEIJING = iota
	SHANGHAI 
	CHENGDU 
)
func main() {
	const length int = 10
	
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("CHENGDU = ", CHENGDU)
}