package main

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

type phone struct{}

func (p *phone) Touch(a string) string {
	c := a
	return c
}

func main() {
	for i := 0; i < 10; i++ {
		go SayHello()
		fmt.Println(i)
	}
	p := phone{}
	go fmt.Println(p.Touch("World"))
}
