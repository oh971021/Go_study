package main

import "fmt"

func zeroval(ival int) {
	ival = 0 // 함수로 입력되는 값을 그대로 0으로 복사해서 가져온다.
}

func zeroptr(iptr *int) {
	*iptr = 0 // 함수로 입력되는 값을 포인터를 통해 접근한 뒤 역참조 해서 값을 0으로 변환
}

func main() {
	p := fmt.Println // 프린트 명령어를 p 에 대입시킨다.

	i := 1 // i 는 변수, 값은 1이다.
	p("initial:", i)
	// i 를 바로 프린트한 값이 출력된다.

	zeroval(i)
	p("zeroval:", i)
	// i 를 zeroval 함수에 대입시켜서 값을 출력한다.

	zeroptr(&i)
	p("zeroptr:", i)

	p("pointer:", &i)
}
