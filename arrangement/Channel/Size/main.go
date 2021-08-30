package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // 0개의 사이즈 채널 생성

	go square()

	// 현재 ch의 사이즈는 0 이므로 데이터 하나가 들어간다.
	// 채널에 값을 집어 넣고, 빠져나가는 부분이 없어서
	// 아래 코드로 진행되지않고, 무한히 대기한다.
	ch <- 9

	// 이 코드가 출력 되도록 하려면, 채널의 사이즈를 키워준다.
	fmt.Println("Naver Print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
