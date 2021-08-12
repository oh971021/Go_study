package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	MakeNumber := func() int {
		result := rand.Intn(100)
		return result
	}

	for {

		var answer int
		result := MakeNumber()
		fmt.Print("랜덤 숫자를 맞춰보세요 : ")

		fmt.Scan(&answer)
		if answer == result {
			fmt.Println("정답입니다.")
		} else if answer > result {
			fmt.Printf("%d 보다 큽니다.\n", answer)
			fmt.Println("다시 입력해주세요.")
			continue
		} else if answer < result {
			fmt.Printf("%d 보다 작습니다.\n", answer)
			fmt.Println("다시 입력해주세요.")
			continue
		}
		fmt.Println("다시 입력해주세요.")
	}
}
