package main

import "fmt"

func main() {
	//声明slice1是一个切片，并分配了空间
	silce1 := []int{1, 2, 3, 4, 5}

	//没有分配空间
	silce2 := []int{}
	silce2 = make([]int, 3)

	var silce3 []int = make([]int, 3)

	//通过:=声明
	silce4 := make([]int, 3)

	fmt.Println("len = %d, silce1 = %v",len(silce1), silce1)

	//判断是否为空
	if silce1 == nil {
		fmt.Println("silce1 is null")
	} else {
		fmt.Println("silce1 is not null")
	}
}