package main

// 페덱스 패키지를 통해서 프로그램을 실행한다.
import (
	"Go-study/service/koreaPS"
)

// fedex, koreaPS 에서 메소드를 포함한 하나의 타입의
// 메소드를 인터페이스로 사용가능하게 만든다.
type Sender interface {
	Send(parcel string)
}

// SendBook 함수는 매객변수로 Sender 라는 인터페이스를 받는다.
// 그렇게 하면서 외부의 메소드를 불러올 수 있따.
func SendBook(name string, sender Sender) {
	// 샌드를 통해 책을 보낸다.
	sender.Send(name)
}

// 하나의 변경사항이 생겼을 때,
// 기존 코드의 대부분을 바꿔야하는 것을 "산탄총 코드"라고 한다. = 샷건수술
func main() {
	// Fedex 패키지의 FedexSender 메소드를 통해 책을 발송한다.
	var sender Sender = &koreaPS.PostSender{}
	SendBook("어린왕좌", sender)
	SendBook("잭과콩나물", sender)
}

/*
인터페이스를 쓰는 이유 ?
	- 외부에서 접근 할 때, 구체화된 객체를 가르킬 필요 없이
	  추상화된 영역만을 바꿔주면서 쉽게 변경이 가능하기 때문이다.

추상화 ?
	- 내부동작 (구체화된 객체)를 감춰서 제공자 사용자 모두에게 자유를 주는 방식
	- 구체화와 이어지는 의존성을 끊을 수 있기 때문 ( decoupling )
	- 변경 영역을 최소화 하기 위해 decoupling 을 사용 한다.
*/

/*
### Duck Typing ??? ###
  - 어떤 객체가 있는데, 그 객체의 행동을 보고 객체가 무엇인지 유추하는 것
	- ex : 어떤 새가 있을 때, 오리처럼 걷고, 날고, 소리내면 오리라고 유추한다.

== Why use Duck typing ??? ==
	- 사용자 중심의 코딩 !!
		: 인터페이스 구현 여부를 타입 선언 시 하지않고 (다른언어는 선언과 동시에 사용)
		  인터페이스 사용 될 때, 결정이 된다. 그래서 필요에 따라 정의해서 사용 가능하다.
*/
