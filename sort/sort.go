package main

import (
	"fmt"
	"sort"
m)

type Person struct {
	Name string
	Age  int
}

// 스튜던트의 슬라이스 타입인데 별칭이다.
// 다른 곳에서 사용 할 이름을 정해줌.
type Friend []Person

// 인터페이스가 가져야 하는 3가지 함수
func (s Friend) Len() int           { return len(s) } = 2
func (s Friend) Less(i, j int) bool { return s[i].Age < s[j].Age } = 공부 필요
// 선언해놓은 구조체가 담겨 있는 배열을 건드림.
func (s Friend) Swap(i, j int)      { s[i], s[j] = s[j], s[i] } = 공부 필요

func main() {
	p := []Person{
		// 구조체
		{"준석", 25},
		{"호민", 55},
		{"민영", 28},
		{"수용", 26},
		{"예리", 42},
		{"배훈", 23},
		{"성민", 25},
		{"문성", 31},
	}

	// 인터페이스 타입을 인자로 받는다.
	fmt.Println("Before sort : ", p)
	// sort 패키지 안의 Sort 함수를 사용함
	// Sort() = 매개변수로 인터페이스로 받는다.
	// 그 받은 인터페이스에 포함되어있는 객체(현재는 배열타입의 Person)를 지정 된 정렬(현재는 Frined[i].Age)을 순서대로 정렬해준다.
	sort.Sort(Friend(p))
	fmt.Println("After sort : ", p)
}
