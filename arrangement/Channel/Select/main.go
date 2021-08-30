/*
	select 문
	: 여러 채널에서 동시에 데이터를 기다릴 때 사용한다.
	: 동시에 기다리고, 데이터가 들어올 때마다 사용한다.

	확인법
	: 일정간격으로 신호를 주는 채널로 반환..
	: 일정 시간 대기 후 한번만 신호를 주는 채널 반환...
*/

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

	// 데이터 삽입이 끝나면 여기서 다른 go routine이 끝나기를 기다린다.
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// 1초에 한번 신호가 오는 변수
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		// Select 문에 여러 case가 대기하는 경우,
		// 그리고 같이 발생할 경우 랜덤하게 실행한다.
		select {

		// tick 이라는 채널에서 데이터를 뽑아낸다면
		// 1초가 지났다는 의미임으로 Tick 출력
		case <-tick:
			fmt.Println("Tick")

		// 10초가 지났음을 나타낸다.
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return

		// 얘는 아까 데이터 삽입부분 요녀식이 되면 요녀석으로 실행
		case n := <-ch:
			fmt.Println("Square :", n*n)
			time.Sleep(time.Second)
		}
	}
}
