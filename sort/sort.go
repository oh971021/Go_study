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
func (s Friend) Len() int           { return len(s) }
func (s Friend) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Friend) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	p := []Person{
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
	sort.Sort(Friend(p))
	fmt.Println("After sort : ", p)
}
