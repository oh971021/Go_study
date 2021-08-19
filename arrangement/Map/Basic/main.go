package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	// 맵 선언 (make() 함수 사용해서 생성함)
	m := make(map[int]Product)

	// 맵 데이터 생성
	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}

	// 데이터 삭제
	delete(m, 46)
	// 존재하지 않는 요소 삭제 (에러가 뜨지않음.)
	delete(m, 44)

	// 삭제하거나 없는 요소를 삭제하면 { 0} 이 나온다.
	fmt.Println(m[46])
	fmt.Println(m[44])

	// 데이터의 존재 유무를 확인하는 문법 value 와 ok(bool) 을 반환한다.
	v, ok := m[3]
	fmt.Printf("m[3] 는 존재하는가? %v 값은 %v\n", ok, v)

	// Key 와 Value 을 순환문을 통해 모두 출력함.
	for k, v := range m {
		fmt.Println(k, v)
	}
}
