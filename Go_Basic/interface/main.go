package main

import "fmt"

// 메소드(기능)들의 집합
// 예시 : 사냥 (오리 닭 뭐든 사냥이라는 기능으로 사냥하는 메소드를 생성)
// type 인터페이스명 interface { 메소드 변환 형() }

// 오리, 물고기도 사냥한다.
// 오리, 물고기 클래스 생성
// 공통 메소드를 찾아서 인터페이스로 묶는 작업

type Duck struct{}

func (d Duck) sound() {
	fmt.Println("나는 오리 꽥꽥")
}

func (d Duck) taste() {
	fmt.Println("맛있는 오리껍질")
}

type Fish struct{}

func (f Fish) sound() {
	fmt.Println("물고기 파닥파닥")
}

func (f Fish) taste() {
	fmt.Println("물고기 지글지글")
}

// 에이미가 사냥을 간다.

type huntiong interface {
	sound()
	taste()
}

func inTheriver(h huntiong) {
	h.sound()
}

func main() {
	var Mcdonald Duck
	var Marine Fish

	inTheriver(Mcdonald)
	inTheriver(Marine)
}
