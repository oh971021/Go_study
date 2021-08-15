package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 리스트 형식을 하나 만들어준다.
	// 리스트 인스턴스를 만듬.
	v := list.New()

	// 리스트 젤 뒤에다가 4를 넣고,
	e4 := v.PushBack(4)

	// 젤 뒤에다가 1을 넣는다.
	e1 := v.PushFront(1)

	// Pushback = 젤 뒤에 넣고, Pushfront = 젤 앞에 넣는다.

	// 얘는 새로만든 원소를 반환한다. e4 앞에.
	v.InsertBefore(3, e4)

	// e1 뒤에다가 원소를 집어넣는다.
	v.InsertAfter(2, e1)

	// 맨 앞의 인스턴스까지 돌린다. e 가 nil이면 탈출하고, 돌면서 한번씩 다음 노드로 간다.
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	// 얘는 맨 뒤에 인스턴스부터 돌리고, e 가 nil이면 탈출하고, 돌면서 이전 노드로 간다.
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}
