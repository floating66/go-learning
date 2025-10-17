package main

import "fmt"

type Person struct {
	name string
	age int
}

//传递的是副本
func changePeople(p Person) { 
	p.name = "Jerry"
}

func changePeople2(p *Person) { 
	p.name = "Jerry"
}

func main() {
	var people Person
	people.name = "Tom"
	people.age = 18
	fmt.Println(people)

	fmt.Println("------------")

	changePeople(people)
	fmt.Println(people)

	fmt.Println("------------")
	changePeople2(&people)
	fmt.Println(people)
}