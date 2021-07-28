package main

import "fmt"

type myInt int

/*
func (a int) Del(b int) int { // 요 친구는 현재 a가 로컬타입이 아니라서 사용할 수 없다.
	return a + b
} */

func (m myInt) Add(a int) myInt { // 메소드
	rst := int(m) + a
	return myInt(rst)
}

func Add(m myInt, a int) myInt { // 함수
	rst := int(m) + a
	return myInt(rst)
}

func main() {
	var a myInt
	a = a.Add(10)
	fmt.Println(a)
}
