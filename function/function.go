package main

import "fmt"

func say(msg *string) {
	println(*msg) // Hello를 출력한다.
	*msg = "Changed"
	// msg 메모리 주소값을 받아서
	// 메모리 주소값의 값을 변환시킨다.
}

func main() {
	msg := "Hello"
	say(&msg)        // say 함수에 msg 메모리 주소값을 보낸다.
	fmt.Println(msg) // 함수에 갔다 온 변환 된 msg를 출력한다.
}
