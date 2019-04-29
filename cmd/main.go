package main

import (
	"fmt"
)

type Person struct {
	name string
}

type Human interface {
	speak() string
}

func (p *Person) speak() string {
	return "I'm " + p.name
}

func saySomething(h Human) {
	fmt.Println("He's speaking ->", h.speak())
}

func main() {
	p := Person{"Bob"}
	//works
	saySomething(&p)

	//doesn't work
	saySomething(p)

	//works
	p.speak()
}
