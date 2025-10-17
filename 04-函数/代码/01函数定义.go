package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func foo1(x int, y string)(r1,r2 int){
	println("fool1----------------")
	println("x = ",x)
	println("y = ",y)

	r1 = 100
	r2 = 200
	return
}

func foo2(x int, y string)(int,int){
	println("fool2----------------")
	println("x = ",x)
	println("y = ",y)

	return 666, 777
}

func main() {
	
	fmt.Println(add(1, 2))

	ret1 , ret2 := foo1(100,"aaa")
	fmt.Println("ret1 = ",ret1,"ret2 = ",ret2)

	ret1 , ret2 = foo2(200,"bbb")
	fmt.Println("ret1 = ",ret1,"ret2 = ",ret2)
}