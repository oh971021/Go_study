package main

import "fmt"

// 스택에 사용할 노드를 선언한다.
type Node struct {
	Value int
}

// 스택 구조의 본체 선언
type Stack struct {
	nodes []*Node
	count int
}

// 스택(LIFO)
func NewStack() *Stack {
	return &Stack{}
}

// 데이터 삽입 (Push)
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// 데이터 빼내기 (Pop)
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// 값 출력을 위한 String 함수 선언
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func main() {
	// 스택 자료구조 리스트를 배열로 만든다.
	// 스택 선언
	s := NewStack()

	// 5개 데이터 push
	s.Push(&Node{1})
	s.Push(&Node{2})
	s.Push(&Node{3})
	s.Push(&Node{4})
	s.Push(&Node{5})

	fmt.Println(s)
	fmt.Println(s)

	// 5개 데이터 pop
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// 없는 데이터를 빼려고 하면 nil 표시
	fmt.Println(s.Pop())

	fmt.Println(s)
}
