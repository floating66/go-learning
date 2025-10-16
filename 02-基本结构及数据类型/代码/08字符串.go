package main

import (
    "fmt"
    "strings"
)

func main() {
	//修改字符串大小写
    var orig string = "Hey, how are you George?"
    var lower string
    var upper string

    fmt.Printf("The original string is: %s\n", orig)
    lower = strings.ToLower(orig) //ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符
    fmt.Printf("The lowercase string is: %s\n", lower)
    upper = strings.ToUpper(orig) //ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符
    fmt.Printf("The uppercase string is: %s\n", upper)
}