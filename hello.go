package main

import (
	"fmt"

	"github.com/seymac/ist-coders-golang/composition"
	"github.com/seymac/ist-coders-golang/shape"
)

func main() {
	fmt.Println("Hello world")
	//variable definitions
	var result int
	var x = 1
	y := 2

	//function call
	result = Sum(x, y)

	fmt.Println(fmt.Sprintf("result is: %d", result))

	//interface sample
	triangle := shape.Triangle{10, 8}
	rectangle := shape.Rectangle{6, 5}

	shape.PrintShapeArea(triangle)
	shape.PrintShapeArea(rectangle)

	//private function call
	// shape.privateMultiplyFunction(3,4)

	//struct initialization
	person := composition.Person{Name: "Jane", Lastname: "Doe"}
	person.SayHello()

	//function call after composition
	student := composition.Student{composition.Person{"John", "Smith"}, "Harvard University"}
	student.SayHello()

}

//sample function
func Sum(x int, y int) int {
	return x + y
}
