package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	flag := true

	for flag {
		// false 가 나오면 다시 돌리기 위한 시드 생성
		rand.Seed(time.Now().UnixNano())

		// 난수 0~ 200 생성
		r := rand.Intn(200)

		starttime := time.Now()
		for {
			if r > 100 {
				// 만약 2 ~ 9 배수가 맞다면 무한 루프 탈출
				for i := 2; i < 9; i++ {
					if r%i == 0 {
						fmt.Printf("랜덤함수 값: %d , %d의 배수", r, i)
						flag = false
						break
					} else {
						continue
					}
				}
				endtime := time.Since(starttime)
				fmt.Printf("실행시간 : %v\n", endtime)
				break
			} else {
				continue
			}
		}
	}
}
