/*
	Context?
	- 작업을 지시 할 때, 작업 가능시간, 취소 등 조건 지시할 수 있는
	  작업 명세서 역할을 함.
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// Context를 만들때는 기본 Context를 받아서 덮어써서 사용한다.
	// 반환값으로는 변수, 함수를 반환한다.
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)

	// 5초 뒤에 cancel을 실행한다.
	time.Sleep(5 * time.Second)
	// Done()으로 취소 지시를 보낸다.
	cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	Tick := time.Tick(time.Second)
	for {
		select {
		// cancel 에서 이쪽으로 지시가 온다.
		// 작업이 끝날 때 시그널을 받는다.
		case <-ctx.Done():
			// 그래서 작업을 끝낸다.
			wg.Done()
			return
		case <-Tick:
			fmt.Println("tick")
		}
	}
}
