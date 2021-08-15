package main

import "fmt"

// 객체 Point { x, y int } 를 생성.
type Point struct {
	x, y int
}

// Method 는 객체를 포함하고 있는 함수라고 보면된다.
// 함수와 기능이라는 역할을 하는 건 같지만 객체가 있다 없다의 차이
func (p Point) print() int {
	return p.x + p.y
}

func main() {
	// x = 3 , y = 4 각각 대입시켜주고, 변수에 객체를 넣어준다.
	// 객체를 대입시킨 p 는 인스턴스가 된다.
	p := Point{3, 4}
	// 거기서 Method 를 호출해서 기능을 사용하고,
	// Print 해준다.
	fmt.Println(p.print())
}
