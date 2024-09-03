package main

import "fmt"

type Car struct {
	Name string
	Age int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Driving() {
	fmt.Println("driving...")
}

func (c Car) GetName() string{
	return c.Name
}

func main () {
	c := Car{
		Name: "chevy",
		Age: 2,
		ModelNo: 4,
	}

	c.Print()
	c.Driving()
	fmt.Println(c.GetName())
}