package main

import "fmt"

func main() {
    var i1 int
    var f1 float32
	//空白符用来匹配一些不需要的值，然后丢弃掉
    i1, _, f1 = ThreeValues()
    fmt.Printf("The int: %d, the float: %f \n", i1, f1)
}

func ThreeValues() (int, int, float32) {
    return 5, 6, 7.5
}