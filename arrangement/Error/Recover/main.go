// Panic 객체를 반환하는 Recover : Defer와 함께 사용한다.
// 결국 Recover는 panic을 복구하고, 그 다음 라인을 실행 시키는 것.

package main

import (
	"fmt"
)

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		// 여기서 panic의 전파를 멈춘다. (패닉 복구)
		// recover가 시도 된다는 것은 panic 이 있다는 것
		// r 은 panic이 가진 string을 가진다.
		if r := recover(); r != nil {
			fmt.Println("Panic 복구 - ", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0 일 수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨")
}
