/*
인터페이스를 구현하기 위해서는 해당 타입이 그 인터페이스의 메서드들을 모두 구현하면 되므로,
위의 Shape 인터페이스를 구현하기 위해서는 area(), perimeter() 2개의 메서드만 구현하면 됩니다.
예를 들어 Rect와 Circle이라는 2개의 타입이 있을 때,
Shape 인터페이스를 구현하기 위해서는 아래와 같이 각 타입별로 2개의 메서드를 구현해 주면 됩니다.
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

// Rect 객체 정의
type Rect struct {
	width  float64
	height float64
}

// #### Rect에 맞는 메소드 두개 구현 #####
func (r *Rect) area() float64 {
	return r.width * r.height
}

func (r *Rect) perimeter() float64 {
	return 2 * (r.width * r.height)
}

// ######################################

type Circle struct {
	radius float64
}

// #### Circle에 맞는 메소드 두개 구현 #####
func (r *Circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r *Circle) perimeter() float64 {
	return 2 * math.Pi * r.radius
}

// ######################################

func main() {
	rect := Rect{4.0, 3.0}
	var triangle Shape

	triangle = &rect
	t := triangle
	fmt.Println(t.area())
	fmt.Println(t.perimeter())
}
