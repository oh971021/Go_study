package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 5; i++ {
		r := rand.Intn(50)
		var sum int
		for j := 1; j <= r; j++ {
			sum = sum + j
		}
		fmt.Printf("%d의 1~ %d사이의 sum(%d) = %d\n", r, r, r, sum)
	}
}
