package main

import "fmt"

func main() {
	// 배열 정렬하기
	arr := [10]int{0, 5, 2, 9, 1, 8, 5, 8, 3, 7}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}
	fmt.Println(arr)
}
