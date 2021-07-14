package main

import "fmt"

func main() {
	a := 10
	b := 12
	c := a + b
	d := a - b
	e := a / b
	f := a * b

	fmt.Println(c, d, e, f)
	fmt.Println()

	fmt.Println("a > b", a > b)
	fmt.Println("a < b", a < b)
	fmt.Println("a == b", a == b)
	fmt.Println("a != b", a != b)
	fmt.Println("a > b && a < b", (a > b) && (a < b))
	fmt.Println("a > b || a < b", (a > b) || (a < b))
	fmt.Println()

	fmt.Println("a & b", a&b)
	fmt.Println("a | b", a|b)
	fmt.Println("a ^ b", a^b)
	fmt.Println("a &^ b", a&^b)
	fmt.Println()

	fmt.Println("a << b", a<<b)
	fmt.Println("a >> b", a>>b)
	1010

}
