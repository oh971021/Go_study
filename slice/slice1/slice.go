package main

import "fmt"

func changeArray(array2 [5]int) [5]int {
	array2[2] = 200
	return array2
}

func changeSlice(slice2 []int) {
	slice2[2] = 300
}

func main() {
	array2 := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	array2 = changeArray(array2)
	changeSlice(slice)

	fmt.Println(array2)
	fmt.Println(slice)
}
