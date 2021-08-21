/*
자바에서는 toString() 메서드가 모두 상속되어있기 때문에 객체를 print 해도 문제없이 print가 된다.

그러나 golang은 사용자 정의 타입 모두 toString()을 상속하는 것은 아니다.

그렇기 때문에 이와 같은 기능을 하는 인터페이스를 구현하면 문제없이 작동할 수 있다.

그것이 Stringer 인터페이스 이다.
*/

package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffeePot := CoffeePot("cold brew")
	fmt.Println(coffeePot) // cold brew coffee pot
}
