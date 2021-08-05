/*
  == 함수 리터럴 ==
 	 - literal : 함수 / 다른 언어에서는 lamda 함수라고 부른다.
	 - 함수를 변수에 넣을 때, 한번에 작성해도 된다.
	 - 그 때는 함수의 이름을 정해주지 않고, func 으로 작성해줘도 된다.
	 ex ) f := add 요친구는 add를 만들어놓고 f에 넣어야 하지만
	      f := func(a, b int) int { return a + b } 요렇게 작성해도 된다는 뜻.

	== 특징 ==
	 - literal function 은 "내부상태" 를 가질 수 있다. ( 함수인데 상태를 가지는 것 : first class )
*/

package main

import (
	"fmt"
)

type OpFn func(int, int) int

func getOperator(op string) OpFn {
	if op == "+" {
		//
		return func(a, b int) int { return a + b }
	} else if op == "*" {
		return func(a, b int) int { return a * b }
	} else {
		return nil
	}
}

func main() {
	var operator OpFn
	operator = getOperator("+")
	var result = operator(3, 4)
	fmt.Println(result)
}
