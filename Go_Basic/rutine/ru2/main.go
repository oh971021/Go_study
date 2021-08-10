package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// i 를 10까지 돌리는 순환문
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Println(strings.Repeat("*", n))
		}(i)
	}
	time.Sleep(time.Second)
}
