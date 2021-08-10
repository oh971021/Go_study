package main

import (
	"fmt"
)

type phone struct {
}

func (p *phone) Click(a, b string) string {
	c := a + b
	return c
}

func main() {
	p := phone{}
	s := p.Click("Touch", " Phone")
	fmt.Println(s)
}
