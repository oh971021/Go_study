/*
익명함수

1. 익명함수란??
	- 말그대로 이름이 없는 함수이다.
	- 형태는 기존의 함수와 같지만 이름이 없는 것이 특징
	- 사용하는 이유 ?
		- 함수 선언 자체가 전역으로 초기화 되어 메모리를 많이 잡아먹기 때문에
		- 기능을 수행할 때마다 호출하는게 귀찮아서

*/

package main

import "fmt"

func main() {
	// 그 자리에서 그대로 함수 사용이 가능하다.
	func() {
		fmt.Println("hello")
	}()

	func(a int, b int) {
		result := a + b
		fmt.Println(result)
	}(1, 3)

	// 함수를 만들어서 변수에 대입해서 사용하는 것도 가능하다.
	result := func(a string, b string) string {
		return a + b
	}("hello", " world!")

	// 함수를 가진 변수를 출력하는 코드
	fmt.Println(result)

	// 변수로 값을 만들어주고, 함수에 대입하는데 그 함수는 또 새로운 변수에 대입한다.
	i, j := 10.2, 20.4
	divide := func(a float64, b float64) float64 {
		// i 와 j 를 받아서 계산함.
		return a / b
		// 인자는 함수가 끝나는 } 뒤에다가 () 로 넣어준다.
	}(i, j)

	// 변수만 출력하면 간단스기루~
	fmt.Println(divide)
}
