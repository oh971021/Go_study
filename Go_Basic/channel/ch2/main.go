package main

import (
	"fmt"
	"strings"
)

func main() {
	// 채널을 선언해주는 함수
	ch := make(chan string)

	// for 순환문을 돌리는데,
	for i := 1; i <= 10; i++ {
		go func(n int) {
			s := strings.Repeat("*", n)
			ch <- s
		}(i)
	}

	star := []string{}
	for i := 1; i <= 10; i++ {
		star = append(star, <-ch)
	}
	fmt.Println(strings.Join(star, "\n"))
}
