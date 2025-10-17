package main

import "fmt"

type Animal interface { 
		Sleep()
		GetColor() string
		GetType() string
	}

	//具体的类
type Cat struct { 
		color string
	}
	
func (this *Cat) Sleep() { 
		fmt.Println("Cat Sleep")
	}

func (this *Cat) GetColor() string { 
		return this.color
	}

func (this *Cat) GetType() string { 
		return "Cat"
	}

func ShowAnimal(animal Animal){
	animal.Sleep()
	fmt.Println("color = ",animal.GetColor())
	fmt.Println("type = ",animal.GetType())
} 

func main() {
	
	// var animal Animal   ------>父类指针，指向子类的具体数据类型
	// animal = &Cat{"Black"}
	// animal.Sleep()

	cat := Cat{"Black"}
	ShowAnimal(&cat)

	
}