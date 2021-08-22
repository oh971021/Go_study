/*
  == 서브 고루틴이 종료 될 때 까지 대기 ==
	- sync 패키지의 WaitGroup을 사용하면 고 루틴이 모두 돌때까지 기다려준다.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	var sum int
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	// 10개의 작업을 기다린다.
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}

	// go routine이 모두 끝날 때 까지 기다린다.
	wg.Wait()
}
