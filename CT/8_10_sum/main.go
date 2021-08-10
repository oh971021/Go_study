package main

import "fmt"

func solution(a, b int) int {
	var res int
	if a < b {
		for i := a; i <= b; i++ {
			res += i
		}
	} else if a >= b {
		for i := b; i <= a; i++ {
			res += i
		}
	}
	return res
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(solution(a, b))
}
