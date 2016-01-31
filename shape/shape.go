package shape

import "fmt"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

func PrintShapeArea(s Shape) {
	fmt.Println(fmt.Sprintf("Shape area is: %f", s.Area()))
}

func privateMultiplyFunction(x int, y int) int {
	return x * y
}
