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

// Call by Value, Call by Reference 를 알아보자.
func (p *Point) add(a int) {
	p.x += a
	p.y += a
}

func (p Point) Mul(a int) {
	p.x *= a
	p.y *= a
	// 얘는 이 함수에서만 적용 된다. 즉, 원본엔 영향을 끼치지 못했다.
}

func main() {
	// x = 3 , y = 4 각각 대입시켜주고, 변수에 객체를 넣어준다.
	// 객체를 대입시킨 p 는 인스턴스가 된다.
	p := Point{3, 4}
	// 거기서 Method 를 호출해서 기능을 사용하고,
	// Print 해준다.
	fmt.Println(p.print())

	// 이 친구는 포인터타입의 CbR 라서 원본값에도 영향을 끼친다.
	p.add(10)

	// 이 친구는 변수타입의 CbV 라서 원본값에 영향을 끼치지 않는다.
	p.Mul(10)

	// 그래서 출력 값을 예상해보면 Call by Reference 로 접근한 add() 함수만 적용되서
	// 3 + 10 , 4 + 10 해서 13, 14 가 출력 될 듯하다.
	fmt.Println(p.x, p.y)
}
