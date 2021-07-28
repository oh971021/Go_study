/*
1. 객체 : 데이터와 기능을 묶은 것. = Typpe + Func
	= 얘네를 메소드로 묶는다.
	= 객체와 객체 간의 관계를 맺는 것
	= 리시버와 매개변수의 객체의 관계를 맺어준다.

2. 객체지향 : 의존성은 떨어뜨리고, 결합도는 높인다.

3. OOP (Object Oriented Programing)
4. 객체가 생기면서 관계 중심으로 프로그램이 변화 되었다.
	- 클래스 다이어그램 (객체간 어떤 관계를 맺고 있는지 확인하는 다이어그램) */

package main

import "fmt"

type account struct {
	balance   int
	firstname string
	lastname  string
}

// 포인터 타입에 포함된 메소드
// mainFunc에서 메소드를 호출 할 때, 30이 인자로 들어온다.
func (a1 *account) withdrawPointer(amount int) {
	// 현재 account 에 들어있는 { 100, 조, 박 } 중 100 에서
	// amount 인자로 들어온 30 을 뺀다.
	a1.balance -= amount
	// 리턴은 되지 않지만 포인터로 값을 변환했기때문에
	// call by reference로 주소값의 값이 변하게 된다.
}

// 값 타입에 포함된 메소드
// 요친구는 mainA 와 다른 인스턴스가 생성된다.
// 요친구를 동작 시키려면 리턴을 시켜주고,
// 메인함수에서 그 값을 대입 받아야 사용 가능하다.
func (a2 account) withdrawValue(amount int) {
	// 값이 복사된 인스턴스에 값을 빼는 것이기 때문에 메인함수의 mainA에는 적용 되지 않는다.
	a2.balance -= amount
}

// 요렇게 값을 리턴해줘야 사용가능하다.
func (a2 account) withdrawValue2(amount int) account {
	a2.balance -= amount
	return a2
}

func main() {
	// 메인이라는 포인터 객체를 하나 만들고, 그 주소에 값을 넣음
	// 밸런스 = 100 / 첫이름 = 조 / 마지막이름 = 박
	var mainA *account = &account{100, "joe", "park"}

	// 메인A의 타입에 속해있는 메소드를 호출
	// 메인A = *account의 주소값
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	*mainA = mainA.withdrawValue2(50)
	fmt.Println(mainA.balance)
}
