package main

import "fmt"

//切片截取
func main() {
	s:= []int{1, 2, 3,}

	//[0,2)
	s1 := s[0:2]  //[1,2]

	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s)

	//copy函数可以拷贝切片
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)
}