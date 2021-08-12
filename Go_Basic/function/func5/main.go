/*
  = 1. 일급 함수란 ?? =
	- 함수 자체를 type 으로 사용하여 다른 함수의 매개변수로 사용하거나, 반환 값으로 사용이 가능하다.
*/

package main

import "fmt"

// calc 함수는 f 라는 함수타입을 가진 친구와 a, b 를 매개변수로 받는다.
func calc(f func(int, int) int, a int, b int) int {
	// 리턴 하기 전 값을 담는 변수를 만드는데
	// 그 변수에는 매개변수로 들어온 multi 에서 a와 b 를 계산하는 값을 담는다.
	result := f(a, b)
	return result
}

func main() {
	// 얘는 익명함수
	// 곱셈하는 친구
	multi := func(i, j int) int {
		return i * j
	}

	// r1 은 calc 함수를 담는 변수인데,
	// calc 를 변수에 담을 때, multi(함수를 가진 변수)와 10, 20을 인자로 사용한다.
	r1 := calc(multi, 10, 20)
	// 위에서 multi 함수를 통해 10과 20일 계산한 값이 여기서 return 되어 출력한다.
	fmt.Println(r1)

	// 얘는 함수를 호출하는데, f 에 익명함수를 만들어서 넣는다. 그리고 10과 20을 인자로 넣는건 동일
	r2 := calc(func(x, y int) int { return x + y }, 10, 20)
	// 익명함수를 통해 x + y로 만들었기 때문에, 10과 20일 더하는 함수가 된다.
	fmt.Println(r2)
}
