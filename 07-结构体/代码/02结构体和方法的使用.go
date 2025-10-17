package main

import "fmt"

type Hero struct { 
		name string
		Ad int
		level int
	}

func (this Hero) Show(){
	fmt.Println("Hero:",this.name)
	fmt.Println("Ad:",this.Ad)	
	fmt.Println("level:",this.level)
}

func (this Hero) GetHero() string{
	return this.name

}

func (this Hero) SetHero(newName string){
	this.name = newName
	
}

func main() {
	hero := Hero{name:"hero1",Ad:100,level:1}
	hero.Show()
}