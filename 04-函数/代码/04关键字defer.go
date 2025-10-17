package main
import "fmt"
//defer 和 return 的执行顺序是先为返回值赋值，然后执行 defer，然后 return 到函数调用处
func main() {
    function1()

	//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("hello1")
	fmt.Println("hello2")
}

func function1() {
    fmt.Printf("In function1 at the top\n")
    defer function2()
    fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
    fmt.Printf("function2: Deferred until the end of the calling function!")
}