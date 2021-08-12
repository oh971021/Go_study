package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	MakeNumber := func() int {
		rand.Seed(time.Now().UnixNano())
		result := rand.Intn(50)
		return result
	}
	result := MakeNumber()
loop1:
	for {
		var answer int
		fmt.Print("랜덤 숫자를 맞춰보세요 : ")

		fmt.Scan(&answer)
		if answer == result {
			break
		} else if answer > result {
			fmt.Println("DOWN")
			goto loop1
		} else if answer < result {
			fmt.Println("UP")
			goto loop1
		}
	}
	fmt.Println("정답입니다.")
}
