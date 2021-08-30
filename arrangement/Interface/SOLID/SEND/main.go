package main

import "fmt"

// 이벤트가 실행되는 interface 생성
type Event interface {
	// 이벤트를 읽는 레지스터 요녀석이 추상화 된 함수
	// 이 친구는 interface를 매개변수로 받는다.
	Register(EventListener)
}

// 이벤트에 상속되어있는 이벤트 리스너
type EventListener interface {
	// 실행되는 함수같죠?
	OnFire()
}

// 메일이라는 스트럭트 생성
type Mail struct {
	// 이 친구는 이벤트 리스너를 통해 이벤트와 이어진다
	listener EventListener
}

// 레지스터는 메일을 갖고있는데, 이벤트리스너를 받는다.
func (m *Mail) Register(listener EventListener) {
	// 리스너는 onfire와 연결되어있는데,
	// 고 친구를 리스너에 넣는다.
	m.listener = listener
}

// 메일 발송
func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

// 알람이라는 이벤트와 메일을 이어준다.
type Alarm struct {
}

// 메일 송신 이벤트
func (a *Alarm) OnFire() {
	fmt.Println("메일이 왔습니다")
}

func main() {
	var mail = &Mail{}
	var listener EventListener

	listener = &Alarm{}

	mail.Register(listener)
	mail.OnRecv()
}
