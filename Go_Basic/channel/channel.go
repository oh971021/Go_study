package main

import "fmt"

func pop(c chan int) {
	fmt.Println("Pop func") // 함수를 사용하면 출력되는 값

	v := <-c // 메인함수의 Chan의 값을 v로 빼낸다. (Pop)
	// v는 main함수로 가서 어떤 값 (이 코딩에서는 c) 이 들어올 때까지 대기하고 있는다.

	// 메인함수에서 c <- 10 이 대입되는 것을 보고, 그 값을 v에 값을 대입시켜서 출력한다.
	fmt.Println(v) // channal을 통해 들어온 v의 값을 출력한다.
}

func main() {
	var c chan int     // 채널선언
	c = make(chan int) // 채널 초기화, Fixed Size
	// 고정된 Fixed Size를 정해주지 않으면 Dead Lcok 버그가 발생한다.
	// 이 버그를 해결하기 위해서 v 라는 쓰레드를 대기 시켜놓고, push 를 통해 쓰레드를 이어준다.

	go pop(c) // go Thread pop / go 선언 뒤 따라오는 함수는 모두 쓰레드가 된다.

	c <- 10 // 10을 넣고, (Push)
	// 대기하고 있던 v 에게 c 값을 밀어 넣는다.
	// v := <-c // 10을 빼준다. (Pop)

	fmt.Println("end of golang Blockchain")
}

// go thread 가 두개다.
