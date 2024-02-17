package main

import "fmt"

//Liskov Substitution Principle primarily deals with the inheritance
// If you have some API that takes the Base Class and works correctly with base class, it should also work
// correctly with the derived class as well.

// Interface Segregation Principle states that, try to break up an interface into separate parts that people would need. 

type Shape interface {
	//Getters and Setters
	GetLength() int
	SetLength(length int)
	GetBreadth() int
	SetBreadth(breadth int)
}

type Rectangle struct {
	length, breadth int
}

//Implementing the interface by setters and Getters

func (r *Rectangle) GetLength() int {
	return r.length
}

func (r *Rectangle) SetLength(length int) {
	r.length = length
}

func (r *Rectangle) GetBreadth() int {
	return r.breadth
}

func (r *Rectangle) SetBreadth(breadth int) {
	r.breadth = breadth
}

func Use(shape Shape) {
	shape.SetLength(10)
	shape.SetBreadth(20)
	expectedArea := 10 * 20
	actualArea := shape.GetLength() * shape.GetBreadth()

	fmt.Println("Actual are is : %d, Expected area is : %d", actualArea, expectedArea)
}

func main() {
	rect := &Rectangle{2, 3}
	Use(rect)
}
