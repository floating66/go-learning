package main
import "fmt"

//将一个大数组传递给函数会消耗很多内存。作为替代，我们可以只传递数组的指针或使用数组的切片
func main() {
    array := [3]float64{7.0, 8.5, 9.1}
    x := Sum(&array) // Note the explicit address-of operator
    // to pass a pointer to the array
    fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a { // derefencing *a to get back to the array is not necessary!
        sum += v
    }
    return
}