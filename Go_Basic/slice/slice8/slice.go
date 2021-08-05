package main

import (
	"fmt"
	"sort" // 정렬에 관한 함수가 있음.
)

func main() {
	slice := []int{5, 2, 3, 5, 100, 30, 19, 59, 34, 4, 6, 7}
	fmt.Println("Before sort : ", slice)
	sort.Ints(slice)
	fmt.Println("After sort : ", slice)
}
