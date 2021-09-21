package main

import (
	"fmt"
)

func square(x int) int {
	return x * x
}

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x / y
}

func main() {

	var a, b int

	a = 30
	b = 50

	fmt.Printf("3 * 5 = %d\n", add(3, 5))

	fmt.Printf("9 * 9 = %d\n", square(9))

	fmt.Printf("%d / %d = %d", a, b, mul(a, b))
}
