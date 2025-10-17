package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

//Square implements Shaper interface
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	squ := new(Square)
	squ.side = 5
	//可以直接在 Square 的实例上调用此方法
	fmt.Printf("Square area: %f\n", squ.Area())

	//通过接口实例访问
	//将一个 Square 类型的变量赋值给一个接口类型的变量shp = squ，此时shp指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()
	var shp Shaper
	shp = squ
	fmt.Printf("Square area: %f\n", shp.Area())

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type square is %T ", t)
	case *Circle:
		fmt.Printf("Type circle is %T ", t)
	case nil:
		fmt.Println("nil")
	case default:
		fmt.Printf("Unknown type %T", t)
	}
}