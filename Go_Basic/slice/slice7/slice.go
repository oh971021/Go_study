package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	// 삽입하고 싶은 위치
	idx := 2

	slice = append(slice, 0)
	// 4,5,6,0 에 3,4,5,6 을 대입 시킨다.
	copy(slice[idx+1:], slice[idx:])
	slice[2] = 100

	/*
		// 한줄로 쓰면 요렇게
		// 1,2 + 100,3,4,5,6 의 형태를 가진다
		// 요건 임시 버퍼(새로운 메모리) 를 사용함으로 비효율
		slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	*/

	/*
		slice = append(slice, 0)
		for i := len(slice) - 2; i >= idx; i-- {
			slice[i+1] = slice[i]
		}
		slice[2] = 100
	*/

	fmt.Println(slice)
}
