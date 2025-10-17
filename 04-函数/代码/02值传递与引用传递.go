package main

import "fmt"

// func swap(a int, b int) { 
// 	var temp int
// 	temp = a
// 	a = b
// 	b = temp
// 	fmt.Println("函数内部","a = ",a,", b = ",b)
// }


func swap(a *int, b *int) { 
	var temp int
	temp = *a
	*a = *b
	*b = temp
	fmt.Println("函数内部","a = ",*a,", b = ",*b)
}

func main() {
	var a int = 10
	var b int = 20

	// swap(a,b)
	swap(&a,&b)
	//swap
	fmt.Println("a = ",a,", b = ",b)

	//二级指针
	var p *int
	p = &a
	fmt.Println("*p = ",p)
	fmt.Println("&a = ",&a)
}