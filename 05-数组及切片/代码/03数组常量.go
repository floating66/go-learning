package main
import "fmt"

//若数组值已经知道，可以通过数组常量的方法来初始化数组
func main() {
    // var arrAge = [5]int{18, 20, 15, 22, 16}    [5]int 可以从左边起开始忽略，[10]int {1, 2, 3} : 这是一个有 10 个元素的数组，除了前三个元素外其他元素都为 0
    // var arrLazy = [...]int{5, 6, 7, 8, 22}
    // var arrLazy = []int{5, 6, 7, 8, 22}
    var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}    //只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串
    // var arrKeyValue = []string{3: "Chris", 4: "Ron"}

    for i:=0; i < len(arrKeyValue); i++ {
        fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
    }
}