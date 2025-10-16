package main

import "fmt"


func main001() {
	//布尔类型 bool
	var zz bool = true
	fmt.Println(zz)

	//整型类型
	var a int
    var b int32
    a = 15
    // b = a + a   编译错误,Go 中不允许不同类型之间的混合使用
    b = b + 5    // 因为 5 是常量，所以可以通过编译
	fmt.Println("a = ,b = ",a,b)

	//浮点类型
	var c float32
	c=3.14//3.140000000124
	c=c*c
	fmt.Println(c)
}

