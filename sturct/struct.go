package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// 스튜던트의 슬라이스 타입인데 별칭이다.
// 다른 곳에서 사용 할 이름을 정해줌.
type Students []Student

// 인터페이스가 가져야 하는 3가지 함수
func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18},
	}

	// 인터페이스 타입을 인자로 받는다.
	sort.Sort(Students(s))
	fmt.Println(s)
}
