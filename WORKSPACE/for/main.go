package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("현재 i의 값은 : %d\n", i)
	// }

	a := 0
	for {
		if a == 12 {
			break
		}
		a++
		fmt.Println(a)
	}

	// j := 0
	// for j < 20 {
	// 	fmt.Printf("j의 값은 : %d\n", j)
	// 	j++
	// }

}
