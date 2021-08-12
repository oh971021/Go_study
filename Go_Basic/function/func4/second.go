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
	var nums = []int{10, 12, 13, 14, 16}

	// 얘는 main function 안에서 익명함수를 이용한다.
	addAnonymous := func(nums ...int) (result int) {
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}

	// 결과를 보면 값은 똑같다.
	// 여기서 다른 점은 메모리를 사용하는 차이점과 호출하는 귀차니즘이 필요없다는 것
	fmt.Println(addAnonymous(nums...))
	fmt.Println(addDeclared(nums...))
}
