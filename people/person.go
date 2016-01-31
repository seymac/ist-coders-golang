package people

import "fmt"

type People interface {
	SayHello()
}

type Person struct {
	Name     string
	Lastname string
}

func (p Person) SayHello() {
	greetingText := fmt.Sprintf("Hello %s %s!", p.Name, p.Lastname)
	fmt.Println(greetingText)
}

type Student struct {
	Person
	University string
}

func (s Student) SayHello() {
	greetingText := fmt.Sprintf("Hello %s %s from %s!", s.Name, s.Lastname, s.University)
	fmt.Println(greetingText)
}

type Developer struct {
	Person
	Company string
}

func (d Developer) SayHello() {
	greetingText := fmt.Sprintf("Hello %s %s from company %s!", d.Name, d.Lastname, d.Company)
	fmt.Println(greetingText)
}
