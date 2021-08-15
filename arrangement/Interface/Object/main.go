package main

import "fmt"

func main() {
	// 빈 인터페이스란 ???
	/*
		- 어떠한 타입의 매개변수라도 다 받아 주는 Type 이라고 보면된다.
		- 파이썬에서는 Object type 으로 불린다.
		- 결국 interface{} 를 Type으로 사용하는 매개변수는
		- interface{} type을 통해서 매개변수의 Type을 체크하고,
		- 그에 맞는 Type을 반환하기 때문에 어떠한 Type의 변수가 들어오든 상관없다.
		- 굉장히 편한 것 같네요. 그러면 기초 문법에도 적용시켜보기.
	*/

	// x 는 어떤 타입이든 변할 수 있다.
	var x interface{}

	x = "Tom"

	printIt(x)

	fmt.Printf("%T\n", Change(x))

}

// interface{} 타입을 받는 v ... 이 친구도 어떤 타입이든 들어오는대로 변한다.
func printIt(v interface{}) {
	fmt.Printf("%T\n", v) //Tom
}

// 그러면 호출에서도 타입이 변환도 가능할까?
// 가능하네요!!
func Change(v interface{}) interface{} {
	v = 3.14
	return v
}
