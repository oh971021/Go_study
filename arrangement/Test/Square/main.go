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

func main() {
	fmt.Printf("3 * 5 = %d\n", add(3, 5))

	fmt.Printf("9 * 9 = %d\n", square(9))
}
