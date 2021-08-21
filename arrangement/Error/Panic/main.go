/*
	# 패닉이란 ? #
	- Error가 발생 했을 때, 빠르게 프로그램을 조기 종료하는 방법.
	- 즉, Error를 만나자마자 프로그램을 즉시 종료해서 호출자에게 알려준다.
	- built in 함수인 듯 하다.
*/

package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) error {
	if b == 0 {
		return errors.New("b 는 0 일 수 없습니다.")

		// 위의 errors.New 는 error 가 발생해도 프로그램을 종료하지 않지만
		// 여기서 문제가 생겼으니까 프로그램을 종료하고 알려준다.
		// panic("b 는 0 일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	return nil
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
