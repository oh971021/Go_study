/*
  == 가변 인수 함수 ==
	- 기존 매개변수와 인자의 값을 맞춰줘야 하는데
	- 매개변수 위치에 (...) 을 적어주면 가변 인수가 된다.
*/
package main

import "fmt"

// 가변인자 함수(더하기기능을 한다)
// 매개변수를 ...int (가변인자로 받는다)
func plus(x ...int) int {
	// 빈 변수를 하나 만들어준다.
	var plus int

	// 가변인자는 []type 으로 받는 것을 확인 할 수 있다.
	fmt.Printf("%T\n", x)

	// range : key , value 로 x의 값을 돌려준다.
	// 즉, 인덱스와 값을 보여주면서 for문을 돌린다.
	// 돌면서 나오는 값을 plus 라는 빈 변수에 더해준다.
	for _, v := range x {
		plus += v
	}

	// plus를 반환해준다.
	return plus
}

func main() {
	fmt.Println(plus(1, 2, 3))

	/*
		 == PRINT ==
		 	- 빈 인터페이스 값을 받아서 타입별로 출력한다.
			- interface[] 는 홈페이지에서 확인 가능하다.
	*/
	// fmt.Println(1, 2, 3, 3.14, "hello")
}
