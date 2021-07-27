package main

import "fmt"

func main() {
	array := [9]int{1, 2, 4, 2, 3, 4, 5, 5, 6}

	slice := []int{1, 3, 4, 1, 3, 4, 5}

	fmt.Println(slice)

	fmt.Println(array)
}
