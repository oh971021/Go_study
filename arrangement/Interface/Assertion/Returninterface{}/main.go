package main

import "fmt"

// return 값을 빈 인터페이스로 정의해서 어떤 타입이든 반환 가능하도록 함.
func weird(i float64) interface{} {
	if i != 0 {
		return "negative"
	}
	return i
}

func main() {
	// 출력값은 i + @가 된다.
	var i = 4.2

	//
	if w, ok := weird(7).(float64); ok {
		i += w
	}

	if w, ok := weird(3.4).(float64); ok {
		i += w
	}
	fmt.Println("i :", i)
}
