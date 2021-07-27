package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	// 끝까지 슬라이싱 = : 뒤를 지워줌
	slice2 := slice1[2:]

	// 전체 슬라이싱 = [:] = 요건 보통 배열을 슬라이스로 바꿀 때 사용
	slice3 := slice1[:]

	fmt.Println(slice2)
	fmt.Println(slice3)
}
