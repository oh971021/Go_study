package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // 0개의 사이즈 채널 생성

	wg.Add(1)

	go square(&wg, ch)

	// 데이터가 하나만 들어가는 ch 공간에
	// 10번 반복하면서 하나씩 넣는다.
	// 들어간 데이터가 처리 될 때까지, 다음 데이터는 대기한다.
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	// 데드락을 피하기 위해 close는 필수!! (무한대기가 발생할 수 있음)
	close(ch)

	// 데이터 삽입이 끝나면 여기서 다른 go routine이 끝나기를 기다린다.
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// go routine으로 실행되며 데이터 처리가 될때까지 기다린다.
	// 데이터 하나가 들어올 때마다 Range로 데이터를 뽑아내면서 처리한다.
	for n := range ch {
		// 결국 데이터가 들어오면 Square : n의 제곱을 출력함으로 데이터를 사용함.
		fmt.Println("Square:", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

// 마지막에 Dead Lock이 뜨는 이유 : Go routine 상에 서로 기다리다가 결국 데드락에러가 발생
// 이러한 현상을 좀비고루틴 이라고 한다.

/*
	좀비 고루틴이란?
	- 채널을 닫아주지 않아서 무한 대기를 하는 고루틴을 좀비 고루틴 또는 고루틴 릭(Leak) 이라고 한다.
*/
