/*
= 동시성 프로그래밍을 하게되면 크게 대두되는 문제점 : 동기화 =

*/

package main

import "fmt"

func sigma(start, end int, c chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	var a int
	var b int
	c := make(chan int)
	// 1 부터 100 까지 더하고,
	go sigma(1, 100, c)
	// 거기에 101 부터 200 까지 더한다.
	go sigma(101, 200, c)
	a, b = <-c, <-c
	fmt.Println(a + b)
}
