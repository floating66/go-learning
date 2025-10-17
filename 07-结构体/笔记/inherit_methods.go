/*
定义一个结构体类型 Base，它包含一个字段 id，方法 Id() 返回 id，方法 SetId() 修改 id。

结构体类型 Person 包含 Base，及 FirstName 和 LastName 字段。结构体类型 Employee 包含一个 Person 和 salary 字段。

创建一个 employee 实例，然后显示它的 id。

*/
package main

import "fmt"

type Base struct {
	id int
}

func (this *Base) Id() int {
	return this.id
}

func (this *Base) SetId(id int) {
	this.id = id
}

type Person struct {
	Base
	FirstName string
	LastName string
}

type Employee struct {
	Person
	Salary int
}

func main() {
	emp := new(Employee)
	emp.FirstName = "li"
	emp.LastName = "si"
	emp.SetId(1)
	emp.Salary = 1000
	fmt.Println("Id:",emp.Id())
	fmt.Println("Name:",emp.FirstName,emp.LastName)
}