package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var n, m int
	n, m = a, b

	for i := 0; i < m; i++ {
		fmt.Print("*")
		for j := 1; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
