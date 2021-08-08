/*
 == TDD 란? ==
	- 설계 및 로직의 한 종류
	- 테스트를 먼저하고, 그 다음에 코딩을 한다. ( 만들어진 코드가 없지만 테스트하는 것 )
	- 즉, 실패를 성공으로 바꿔가면서 로직을 개선( 성공강화 )해나가며 짜가는 작업
*/

package main

import (
	"fmt"
)

var opMap map[string]func(int, int) int

func initOpMap() {
	opMap = make(map[string]func(int, int) int)

	opMap["+"] = add
	opMap["-"] = del
	opMap["*"] = mul
	opMap["/"] = div
	opMap["**"] = pow
}

func add(a, b int) int {
	return a + b
}

func del(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func pow(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

func Calculate(op string, a, b int) int {
	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0
}

func main() {
	initOpMap()
	test()
}

func test() {
	// ! : false
	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	}

	if !testCalculate("Test2", "+", 5, 4, 9) {
		return
	}

	if !testCalculate("Test3", "-", 5, 3, 2) {
		return
	}

	if !testCalculate("Test4", "-", 3, 6, -3) {
		return
	}

	if !testCalculate("Test5", "*", 3, 7, 21) {
		return
	}

	if !testCalculate("Test6", "*", 3, 0, 0) {
		return
	}

	if !testCalculate("Test7", "*", 3, -3, -9) {
		return
	}

	if !testCalculate("Test8", "/", 9, 3, 3) {
		return
	}

	if !testCalculate("Test9", "**", 2, 3, 8) {
		return
	}

	if !testCalculate("Test9", "**", 3, 3, 27) {
		return
	}

	fmt.Println("Success!!")
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Test1 Failed!! expected:%d , output:%d\n", testcase, expected, o)
		return false
	}
	return true
}
