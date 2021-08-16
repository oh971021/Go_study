/*
= Go routine 과 동시성 프로그래밍!! =

- 동시성 프로그래밍이란 여러가지 작업(태스크:task)가 동시에 처리되는 것을 의미한다.
병렬 프로그래밍이랑 비슷하긴한데 사실 같은 개념은 아니다.
동시성(concurrency)이란 것은 라운드로빈, 시분할처리를 생각하면 편하다.
병렬성(parallelism)은 정말로 동시에 같이 일어나는 것을 의미한다.
멀티프로세싱이나 멀티쓰레딩을 생각하면 된다.

- go에서는 그럼 os에서 스케줄링안하고 내가(go runtime) 스케줄링할께!하고 만든것이 바로 goroutine이다.
그래서 go에서는 쓰레드를 쓰지 않는다. goroutine을 쓸 뿐이다.
더 정확히 말하면 쓰레드와 프로세스를 os에게 내놔라고 하지 않고
그냥 goroutine을 돌리면 알아서 필요할 경우 쓰레드를 꺼내서 쓴다.(go 1.5버전 이후)
일반적으로는 go runtime에서 자체적으로 시분할 처리를 해서 사용한다.
*/

package main

import (
	"fmt"
	"time"
)

var count int

func first() {
	for i := 0; i < 10; i++ {
		count++
		time.Sleep(time.Second / 10)
		fmt.Printf("first : %d\n", count)
	}
}

func second() {
	for i := 0; i < 10; i++ {
		count++
		time.Sleep(time.Second / 10)
		fmt.Printf("second : %d\n", count)
	}
}

func main() {
	go first()
	go second()
	for i := 0; i < 10; i++ {
		count++
		time.Sleep(time.Second / 10)
		fmt.Printf("main : %d\n", count)
	}

	fmt.Println("3개의 쓰레드로 돌린 count :", count)
}

// 출처 : https://kamang-it.tistory.com/entry/Go19%EA%B3%A0%EB%A3%A8%ED%8B%B4%EA%B3%BC-%EC%B1%84%EB%84%90%EC%96%B4%EB%96%BB%EA%B2%8C-%EB%8F%99%EA%B8%B0%ED%99%94%EB%A5%BC-%EC%99%84%EC%84%B1%ED%95%98%EB%8A%94%EA%B0%80?category=752251
