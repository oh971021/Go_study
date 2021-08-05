/*
== 가변 인수 함수 ==
	- 기존 매개변수와 인자의 값을 맞춰줘야 하는데
	- 매개변수 위치에 (...) 을 적어주면 가변 인수가 된다.
*/
package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums type: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(3, 4, 5))
	fmt.Println(sum(0))
}
