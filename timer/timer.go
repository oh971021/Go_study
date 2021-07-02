package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++ // 2초 간격으로 0~ 1씩 증가하는 정수
	}
}

func main() {
	c := make(chan int) // 채널 생성

	go push(c) // 채널1을 push 함수에 c 인자로 넘기고 go thread 실행

	timerChan := time.After(10 * time.Second)   // 특정시간을 주기적으로 반환하는 타이머 채널을 만든다.
	tickTimerChan := time.Tick(2 * time.Second) // 증가값을 2초 단위로 tick 이 알려준다.

	for { // loop문
		select {
		case v := <-c: // c에 입력값이 있으면 아래 v 출력
			fmt.Println(v)
		case <-timerChan: // 대기는 하되, 10초 뒤 실행된다.
			fmt.Println("timed out")
			return
		case <-tickTimerChan:
			fmt.Println("Tick") // tick 은 주기적 타이머

			/* default: // c에 입력값이 없거나 일반적일 때는 아래 Idle 출력
			fmt.Println("Idle")
			time.Sleep(1 * time.Second) // 1초 단위로 Idle 출력 */
		}
	}
}
