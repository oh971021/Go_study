package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	starttime := time.Now()
	for {
		none := false
		a := int((rand.New(rand.NewSource(time.Now().UnixNano())).Float64() * 100) + 100)
		for i := 2; i < 10; i++ {
			if a%i == 0 {
				none = false
				fmt.Printf("랜덤 난수 중 값은 %d이고, %d의 배수입니다. \n", a, i)
				break
			} else {
				none = true
			}
		}
		if none {
			fmt.Printf("%d는 2~9 사이의 어느 배수도 아닙니다\n", a)
		} else {
			break
		}
	}
	endtime := time.Since(starttime)
	fmt.Println("계산하는데 걸린시간 : ", endtime)
}
