/*
  = Closer 란 ?? =
   - 함수 안에서 익명 함수를 정의하고, 바깥 함수에서 선언한 변수에도 접근 할 수 있는 함수를 말한다.
   - 원래는 pass by value 형식이나 pass by reference 로 접근해야하는데, 익명함수는 클로저라서 그냥 된다.
*/

package main

import "fmt"

var float float64 = 5.4

func main() {
	a, b := 10, 20
	str := "Hello Junseok"

	// 매개변수에 아무것도 안넣고, 인자로 넣지 않아도
	// 익명함수는 클로저 속성을 가져서 함수 외부에 있는 변수에도 접근할 수있다. (스코핑)
	// 내 생각으론 파이썬에서는 global 변수에 접근하는 느낌인 듯 하다.
	result := func() int {
		return a + b
	}()

	// 얘도 바깥에 있는 것도 사용 가능
	func() {
		fmt.Println(str)
	}()

	// 전역변수도 접근이 가능하네요 신기하다..
	func() {
		fmt.Println(float)
	}()

	fmt.Println(result)
}
