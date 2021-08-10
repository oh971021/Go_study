package main

import "fmt"

func main() {
	// 5개의 key를 가진 map을 선언해준다.
	maps := map[int]string{
		1: "First",
		2: "Second",
		3: "Tirth",
		4: "Force",
		5: "Fifth",
	}

	// map 은 range 로 for 문을 돌리면 key 값과 value 값 둘다 출력 가능하다.
	for key, val := range maps {
		fmt.Printf("maps[%d] = %s\n", key, val)
	}
}
