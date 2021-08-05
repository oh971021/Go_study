package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	// copy로 복사하기
	slice2 := make([]int, len(slice1))
	// 첫번째로 받을 값, 두번째론 복사할 값
	copy(slice2, slice1)

	/* append로 슬라이스 복사하기 (... 은 슬라이스, 배열의 전체를 가져옴)
	slice2 := append([]int{}, slice1...) */

	/* 순환문으로 슬라이스 복사하기
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	} */

	slice2[1] = 100

	fmt.Println("slice1 : ", slice1)
	fmt.Println("slice2 : ", slice2)
}
