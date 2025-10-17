package main

import "fmt"

// 匿名字段 内嵌
type inner struct {
	in1 int
	in2 int
}

type outer struct {
	out1 int
	out2 float64
	int 
	inner
}

func main() {
	o := new(outer)
	o.out1 = 1
	o.out2 = 2.0
	o.int = 5
	o.in1 = 3
	o.in2 = 4
	fmt.Println(o.out1)
	fmt.Println(o.out2)
	fmt.Println(o.int)
	fmt.Println(o.in1)
	fmt.Println(o.in2)

}