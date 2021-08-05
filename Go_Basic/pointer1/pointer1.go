package main

import "fmt"

func changeValue(str *string) { // string 포인터 타입의 str 변수가 생긴다.
	// 즉, str은 string 타입을 가진 변수의 포인터역할을 한다.
	// 그래서 메인 함수에서 toChange의 메모리 주소값을 가지고 오고, 그 값의 포인터 값이 str이 된다.
	*str = "changed!"
}

func changeValue2(str string) { // 여기서는 str string 함수가 생기
	str = "changed!"
}

func main() {
	var toChange string = "Hello"
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
}

// & = value의 메모리 주소 위치를 가져온다.
// * =
