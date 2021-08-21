package main

import "fmt"

type Develop interface {
	Coding()
}

// Developer 구조체는  Coding() 이라는 메소드를 구현함으로써 Develop 인터페이스를 만족한다.
// 즉, Develop 타입을 가진 변수에 대해 할당을 해줄 수 있게 되는 것이다.
// 이렇게 인터페이스를 사용하게 되면 구현에 상관없이 오직 기능적인 관점으로 프로그램을 추상화 할 수 있다.
type Developer struct {
	Languages []string
}

func (d *Developer) Coding() {
	for index, values := range d.Languages {
		fmt.Println(index, values)
	}
}

// 인터페이스 타입을 사용하여 함수나 메소드 파라미터의 타입을 지정할 수 있다.
// 해당 인터페이스를 구현하고 있는 구체 타입을 받을 수 있다.
func Work(developer Develop) {
	developer.Coding()
	// 인터페이스 타입으로 하는 경우, 구체 타입에 인터페이스에 선언 된 메소드 이외에 다른 메소드가 있다면
	// 해당 메소드는 함수내에서 일반적으로 호출 할 수 없다.
	// 현재 예시는 Work() 에서 developer 파라미터에 대해 OtherMethod() 메소드를 호출하려 들지만,
	// 인터페이스에는 해당 메소드가 없어서 에러가 난다.
	// developer.OtherMetohd() // -> Error

	// 그 문제를 타입단언으로 해결할 수 있다.
	// 이것은 인터페이스 타입이지만 넘어오는 값이 특정 타입인 경우에 대해 대응할 수 있다.
	// // 인터페이스 타입을 구체 타입으로 변경할 수 있고, ok에 변환여부가 반환된다.
	// if developer, ok := developer.(Developer); ok {
	// 	developer.OtherMethod() // -> Error
	// }
}

func main() {
	var a Developer = Developer{[]string{"Hello", "World"}}
	a.Coding()

	d := &Developer{[]string{"Go", "PHP", "Javascript"}}
	Work(d)
}
