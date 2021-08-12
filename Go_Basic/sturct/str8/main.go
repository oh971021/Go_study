package main

import "fmt"

type triangle struct {
	width, height float32
}

func triArea(s *triangle) float32 {
	return s.height * s.width / 2
}

func main() {
	tri := new(triangle)
	tri.width = 12.5
	tri.height = 5.2

	triA := triArea(tri)
	fmt.Printf("삼각형의 너비:%.2f, 높이:%.2f 일 때, 넓이는 %.2f다.", tri.width, tri.height, triA)
}
