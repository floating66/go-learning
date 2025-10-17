package main

import "fmt"

type List []int

// 值接收者方法
func (l List) Len() int {
    return len(l)
}

// 指针接收者方法
func (l *List) Append(val int) {
    *l = append(*l, val)
}

func CountInto(a Appender, start, end int) {
    for i := start; i <= end; i++ {
        a.Append(i)
    }
}

//	接口Appende   *List 实现了 Appender 接口（因为它有 Append 方法）
type Appender interface {
    Append(int)
}

// 接口Lener      List 实现了 Lener 接口（因为它有 Len 方法）
type Lener interface {
    Len() int
}

func LongEnough(l Lener) bool {
    return l.Len()*10 > 42
}

func main() {
	var lst List // List类型
	//CountInto(lst, 1, 5)  编译错误,lst么有实现Appender接口

	if LongEnough(lst) {
		fmt.Println("lst is long enough")
	}

	plst := new(List)  // *List类型
	CountInto(plst, 1, 5)
	if LongEnough(plst) {
		fmt.Println("plst is long enough")
	}
}