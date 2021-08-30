package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// 채널 인스턴스 생성
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	// 채널에 9라는 데이터를 삽입함
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// 채널에서 데이터를 빼고, n이라는 변수에다가 넣는다.
	n := <-ch

	time.Sleep(time.Second)
	fmt.Println("Square :", n*n)
	wg.Done()
}
