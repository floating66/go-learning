//可以给一个空接口类型的变量 var val interface {} 赋任何类型的值
package main

import (
	"fmt"
	"sort"
)

type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

func Sort(data Sorter) { 
	for pass := 1; pass<data.Len(); pass++ { 
		for i:=0; i<data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

type IntArray []int
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data)
	// a := sort.IntSlice(data)
	sort.Sort(a)
	fmt.Println("Before:", a)
	
}