/*
함수 선언과 익명 함수의 사용법 차이를 보여주는 예제
*/

package main

import "fmt"

// 전역에 새로운 함수를 초기화 해서 사용한다.
func addDeclared(nums ...int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}

func main() {
	// array 변수를 하나 지정해준다.
	var nums = []int{10, 12, 13, 14, 16}

	// 얘는 main function 안에서 익명함수를 이용한다.
	addAnonymous := func(nums ...int) (result int) {
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}(nums...)

	// 결과를 보면 값은 똑같다.
	// 여기서 다른 점은 메모리를 사용하는 차이점과 호출하는 귀차니즘이 필요없다는 것
	fmt.Println(addDeclared(nums...))

	// 얘는 또 함수호출이 아니라 변수를 사용하는 것처럼 사용하는 간편함이 있다.
	fmt.Println(addAnonymous)
}
