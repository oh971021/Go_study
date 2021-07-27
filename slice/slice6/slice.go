package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	// 원하는 인덱스 위치
	idx := 2

	// 카피로 한칸씩 땡긴 다음
	copy(slice[idx:], slice[idx+1:])
	// 마지막칸을 삭제해준다.
	slice = slice[:len(slice)-1]

	/*
		// 아래 구문을 한줄로 사용할 수 있따.
		// 두개로 잘린 슬라이스를 하나로 합친다.
		slice = append(slice[:idx], slice[idx+1:]...)
	*/

	/*
		// 한칸씩 값을 앞으로 변환한다.
		for i := idx + 1; i < len(slice); i++ {
			slice[i-1] = slice[i]
		}

		// 그리고 마지막 한칸을 삭제한다.
		slice = slice[:len(slice)-1]
	*/

	fmt.Println("slice1 : ", slice)
}
