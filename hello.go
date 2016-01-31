package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	//variable definitions
	var result int
	var x = 1
	y := 2

	//function call
	result = Sum(x, y)

	fmt.Println(fmt.Sprintf("result is: %d", result))
}

//sample function
func Sum(x int, y int) int {
	return x + y
}
