//一个新的接口 TopologicalGenus，
// 用来给 shape 排序（这里简单地实现为返回 int）。
// 需要做的是给想要满足接口的类型实现 Rank() 方法
package main

import "fmt"

type Shaper interface {
    Area() float32
}

type TopologicalGenus interface {
    Rank() int
}

type Square struct {
    side float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (sq *Square) Rank() int {
    return 1
}

type Rectangle struct {
    length, width float32
}

func (r Rectangle) Area() float32 {
    return r.length * r.width
}

func (r Rectangle) Rank() int {
    return 2
}

func main() {
    r := Rectangle{5, 3} // Area() of Rectangle needs a value
    q := &Square{5}      // Area() of Square needs a pointer
    shapes := []Shaper{r, q}
    fmt.Println("Looping through shapes for area ...")
    for n, _ := range shapes {
        fmt.Println("Shape details: ", shapes[n])
        fmt.Println("Area of this shape is: ", shapes[n].Area())
    }
    topgen := []TopologicalGenus{r, q}
    fmt.Println("Looping through topgen for rank ...")
    for n, _ := range topgen {
        fmt.Println("Shape details: ", topgen[n])
        fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
    }
}