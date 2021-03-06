package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond) // 1 second = 1000 millisecond
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i < 5+1; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	// 쓰레드가 끝날 때 까지 기다려주기 위해서 Sleep
	time.Sleep(3 * time.Second)
}
