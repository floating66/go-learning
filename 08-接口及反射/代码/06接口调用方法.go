/*
在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：

指针方法可以通过指针调用
值方法可以通过值调用
接收者是值的方法可以通过指针调用，因为指针会首先被解引用
接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

*/
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