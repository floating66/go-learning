package main
import "fmt"

type Engine interface {
    Start()
    Stop()
}

type Car struct {
    Engine
	wheelCount int
}

type Mercedes struct {
    Car
}

func (c *Car) numberOfWheels() int {
    return c.wheelCount
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("hi merkel")
}


func (c *Car) Start() {
    fmt.Println("car start")
}

func (c *Car) Stop() {
    fmt.Println("car stop")
}

func (c *Car) GoToWorkIn() {
	fmt.Println("get in car")
    // get in car
    c.Start()
	fmt.Println("drive to work")
    // drive to work
    c.Stop()
    // get out of car
	fmt.Println("get out of car")
}

func main() {
	// var car Car
	// car.GoToWorkIn()
	fmt.Println("----------------")
	mer := new(Mercedes)
	mer.Car.GoToWorkIn()
	mer.Car.wheelCount = 4
	mer.sayHiToMerkel()
}