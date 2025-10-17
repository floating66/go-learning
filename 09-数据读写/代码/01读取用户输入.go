package main

import "fmt"

var (
	firstName,LastName,s string
	i int
	f float64
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &LastName)
	fmt.Printf("Hi %s %s\n", firstName, LastName)
	fmt.Sscanf(format, &f, &i, &s)  //Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，与 Scanf 相同
	fmt.Println("f:", f, "i:", i, "s:", s)
}