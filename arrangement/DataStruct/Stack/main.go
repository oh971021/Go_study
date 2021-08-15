package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	stack := NewStack()
	books := [5]string{"짱구", "도라에몽", "토토로", "센과 치히로", "짠돌이"}

	fmt.Print("Push : ")
	for i := 0; i < 5; i++ {
		stack.Push(books[i])
		fmt.Printf("%v -> ", books[i])
	}

	fmt.Println()
	val := stack.Pop()

	fmt.Print("Pop : ")
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}
