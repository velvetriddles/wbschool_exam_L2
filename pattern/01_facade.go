package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Class1 struct{}

func (c *Class1) print() {
	fmt.Println("Class1")
}

type Class2 struct{}

func (c *Class2) print() {
	fmt.Println("Class2")
}

type Class3 struct{}

func (c *Class3) print() {
	fmt.Println("Class3")
}

type Class struct {
	C1 Class1
	C2 Class2
	C3 Class3
}

func (c *Class) print1() {
	c.C1.print()
}

func (c *Class) print2() {
	c.C2.print()
}

func (c *Class) print3() {
	c.C3.print()
}
