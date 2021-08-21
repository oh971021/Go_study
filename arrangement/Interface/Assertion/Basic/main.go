package main

import "fmt"

type Jun struct {
	x int
	y int
}

var a Jun = Jun{x: 10, y: 20}

var value interface{} = a

func main() {
	// assert type..
	// value 라는 빈 interface 타입이 Jun 타입으로 변경되었다.
	// 그리고 그것을 출력했다.
	fmt.Println("type assertion value -> interface and value : ", value.(Jun))
}
