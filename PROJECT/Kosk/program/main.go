package main

import "fmt"

// ======================== 전송 인터페이스 ======================= //
// 인터페이스 생성 (메소드를 가진 친구)
type OneMethod interface { //
	String() string
}

// 메소드를 가진 친구 (가장 먼저 시작되는 친구)
type UseMethod_1 interface {
	Transmission() OneMethod // 전송 메소드
}

type Program_1 struct {
	programs string
}

// 프로그램을 실행하면 전송 메소드를 실행하겠다.
func (p *Program_1) StartProgram(usemethod UseMethod_1) {
	start := usemethod.Transmission()
	p.programs += start.String() // 전송된 파일 저장
}

func (p *Program_1) String() string {
	return p.programs
}

// 전송자
type Trans struct {
}

func (t1 *Trans) Transmission() OneMethod {
	return &UseTransmission{}
}

type UseTransmission struct {
}

func (t2 *UseTransmission) String() string {
	return "전송 된 데이터 : "
}

// ============================================================== //