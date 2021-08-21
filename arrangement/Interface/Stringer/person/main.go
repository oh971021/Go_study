// GoStringer는 구조체 인스턴스를 출력할 때 문자열을 반환하도록하는 인터페이스다.

package main

import "fmt"

type JunStringer interface {
	JunString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) JunString() string {
	return fmt.Sprintf("Who : %v, Age : %v", p.Name, p.Age)
}

func PrintUser(junstringer JunStringer) {
	// 인터페이스 변수에서 접근 가능한 건 정의한 메소드 뿐이다.
	fmt.Println(junstringer.JunString())

	// struct의 변수에 접근하기 위해서 인터페이스 변수를 구체적인 타입으로 변환한다.
	// 여기선 *Person struct로 변환 시킨다.
	// 여기가 바로 구체화 시키는 작업. 인터페이스 -> 구조 변경
	// 타입변환이 성공하면 ok가 Ture를 실패하면 False를 반환한다.
	if s, ok := junstringer.(*Person); ok {
		fmt.Printf("User : %v, %v Year\n", s.Name, s.Age)
		fmt.Println()
	} else {
		fmt.Println("error")
	}
}

func main() {
	A := &Person{"Junseok", 25}
	B := &Person{"Solri", 28}

	PrintUser(A)
	PrintUser(B)
}
