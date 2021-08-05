/*
== 인터페이스 ==
	- interface = type (인터페이스 타압의 변수선언이 가능하다)
	- 인터페이스 안에는 구현이 빠진 Method를 적는다 ( ex. func() )
	- 변수가 인터페이스의 메소드를 포함하고 있으면 그 인터페이스를 구현하고 있다.

	== 인터페이스 규칙 ==
		- 메소드는 반드시 메소드이름을 가져야한다.
		- 매개변수 및 반환타입이 달라도 이름이 같은 메소드는 있을 수 없다.
		( 인터페이스 내 오버라이딩 금지)
		- 인터페이스에서 메소드는 구현을 하지 않는다.

*/
package main

import "fmt"

// 인터페이스가 가진 메소드 목록과 객체의 메소드가 현재 겹치고 있다.
// Student 라는 스트럭처가 Stringer 라는 인터페이스를 implement(구현)하고 있따.
type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

// 스트럭처가 String()이라는 메소드를 포함하고 있다.
func (s Student) String() string {
	// Printf과 Sprintf? : 터미널에 출력되는 것과 문자열로 반환되는 차이가 있다.
	return fmt.Sprintf("안녕 나는 %d살 %s라고 해", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	student := Student{"철수", 12}
	// 요 친구가 인터페이스 타입을 지닌 변수
	var stringer Stringer

	// Go 의 두 타입은 같아야 하지만 여기서는 가능하다.
	// 객체의 메소드가 인터페이스내에 포함되고 있기 때문이다.
	// 스트링거는 내가 가르키고 있는 객체가 내가 포함하고있는 메소드를 포함하고 있다.
	// 그래서 인터페이스 정의 된 변수는 인터페이스 내에 있는 메소드만 사용할 수 있다.
	stringer = student
	fmt.Printf("%s\n", stringer.String())
}
