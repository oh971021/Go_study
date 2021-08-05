/*
 == 함수 타입 변수 ==
	- 함수를 값으로 갖는 변수
	- 간단히 말하면 함수를 변수로 사용할 수 있다.

	- 함수도 시작 주소를 갖고 있따.
	- 주소가 숫자이기 때문에 변수의 값으로 가질 수 있고, 표현 가능하다.

	== 타입은? ==
	 - 함수를 받는 변수의 타입은 무엇인가?
	 - 함수시그니처로 표현한다. ( 함수 이름과 구현을 제외한 함수 시그니처)
*/

package main

import (
	"fmt"
)

// 별칭타입으로 줄여서 사용한다.
// 함수의 주소값을 받고 반환하기 때문에 int를 사용한다.
type OpFn func(int, int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// 현재 func (int,int) int : 요 친구가 하나의 타입을 나타낸다.
// op 를 매개변수로 받고, 반환값은 add, mul의 주소값을 반환해준다.
func getOperator(op string) OpFn {
	// + 를 받았을 때,
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	// operator 라는 변수를 선언했는데
	// 타입이 func(int, int) int 함수 타입이다.
	// 함수의 주소를 나타내는 타입인 듯하다.
	// 타입이름이 너무 길어져서 별칭으로 많이 사용한다.
	var operator OpFn
	// + 라는 op를 보낸다.
	// 결국 add 함수의 주소를 호출받아서
	// operator 는 add 함수의 주소를 가진 함수, 변수가 된다.
	operator = getOperator("+")

	// 사용할 때는 add를 사용하는 것처럼 사용해준다.
	var result = operator(3, 4)
	fmt.Println(result)
}
