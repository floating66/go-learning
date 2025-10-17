package main


import "fmt"

func printArr(myArr [4]int){
	for index,value := range myArr {
		fmt.Println(index,value)
	}
}

func main() {
	//固定长度
	var myArr [10]int

	myArr2 := [10]int{1,2,3,4,5,6}

	myArr3 := [4]int{1,2,3,4}

	for i := 0; i < 10; i++ {
		fmt.Println(myArr[i])
	}

	for index,value := range myArr2 {
		fmt.Println(index,value)
	}

	//查看数组数据类型
	fmt.Println("myArr type = %T\n",myArr)
	fmt.Println("myArr2 type = %T\n",myArr2)

	printArr(myArr3)
}