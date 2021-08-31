package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
	fmt.Printf("%d", square(5))
}
