package main

import "fmt"

func main() {
	var a int
	fmt.Println("a = ",a)
	fmt.Printf("type of a = %T\n",a)

	var b int = 100
	fmt.Println("b = ",b)

	var c = 100
	fmt.Println("c = ",c)
	fmt.Printf("type of c = %T\n",c)

	//常用方法，省去var关键字，不支持全局变量
	e := 100
	fmt.Println("e = ",e)
	//多变量
	var f,g int = 100,200
	fmt.Println("f = ",f,"g = ",g)
	var h,i = 100,"aaa"
	fmt.Println("h = ",h,"i = ",i)

	var (
		j int = 100
		k string = "aaa"
	)
	fmt.Println("j = ",j,"k = ",k)
	
}