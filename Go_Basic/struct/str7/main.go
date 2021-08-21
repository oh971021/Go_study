package main

import "fmt"

// 맵 형태의 data 필드를 가지는 mapstruct를 정의한다.
type mapStruct struct {
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	// 구조체 객체를 생성하고 초기화한다.
	d := mapStruct{}
	// data 필드의 맵을 초기화한다.
	d.data = map[int]string{}
	// 초기화한 포인터 구조체를 반환한다.
	return &d
}

func main() {
	s1 := newStruct() // 생성자 호출
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{}
	s2.data = map[int]string{}
	s2.data[1] = "hello"
	s2.data[10] = "world"

	fmt.Println(s2)
}
