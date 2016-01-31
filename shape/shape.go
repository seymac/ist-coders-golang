package shape

import "fmt"

type Shape interface {
	Area() float64
}

type Triangle struct {
	height float64
	base   float64
}

func (t Triangle) Area() float64 {
	return (t.height * t.base) / 2
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return (r.height * r.width)
}

func PrintShapeArea(s Shape) {
	fmt.Println(fmt.Sprintf("Shape area is: %s", s.Area()))
}
