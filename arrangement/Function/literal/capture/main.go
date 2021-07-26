/*
	== 캡쳐 ==
	 - call by reference
*/

package main

import "fmt"

// 캡쳐가 어떻게 돌아가는지 확인해보는 예제
func CaptureLoop() {
	// 우선 3개의 슬라이스를 만들어준다.
	f := make([]func(), 3)
	fmt.Println("ValueLoop")

	// for 문을 돌리는데, i 의 값은 3 까지 플러스 되도록 돌린다.
	for i := 0; i < 3; i++ {
		// f[0], f[1], f[2] 에 각각 i의 값과 함수를 넣어줌. (call by reference)
		f[i] = func() {
			fmt.Println(i)
		}
	}

	// 여기서 함수가 실행 되는데, f[i] 값만 이 for문에서 변경되고,
	// 이미 위의 for문에서 i = 3 이 들어가 있기 때문에 print의 i 는 3 으로 출력된다.
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		// 새로운 변수 공간을 만들어준다. (call by value)
		v := i
		// 여기서도 각각 넣어준다.
		f[i] = func() {
			fmt.Println(v)
		}
	}

	// 여기서 각각의 f[i] 를 실행하게 되는데,
	// loop1 과는 다르게 v 라는 새로운 변수를 사용했기 때문에
	// i의 값은 복사되지 않고, 0, 1, 2 라는 값이 출력 된다.
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}

/*
	i := 0
	// 외부의 어떤 변수를 캡쳐해서
	// 내부 상태로 가져와서 사용할 수 있다.
	// 그래서 외부의 값도 변환된다. ( 즉, 포인터처럼 레퍼런스로 가져와서 사용 되고 있다. )
	f := func() { i += 10 }
	f()
	fmt.Println(i)
*/
