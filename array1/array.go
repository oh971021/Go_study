package main

import "fmt"

func main() {
	var A [10]int

	for i := 0; i < 10; i++ {
		A[i] = i * i
	}
	fmt.Println(A)
	fmt.Println(len(A))

	s := "Hello world"
	v := "Hello 월드"               // 숫자는 3byte라서 rune배열로 바꿔주면 출력이 가능하ㅏㄷ.
	for i := 0; i < len(s); i++ { // 0부터 시작해서 s배열의 길이만큼 1씩 더하면서 반복된다.
		fmt.Print(s[i], " : ")
		fmt.Print(string(s[i]), "\n") // 문자 코드들이 합쳐져서 문자열을 만든다.
	}
	v2 := []rune(v)
	fmt.Println("len(v2) = ", len(v2))
	for i := 0; i < len(v2); i++ {
		fmt.Print(v2[i], " : ")
		fmt.Print(string(v2[i]), "\n")
	}
}

// 배열 = 메모리주소, 길이를 가지고 있다.
// 길이 = Type[?] x 갯수
// var A [10]int32 = int32 : 4byte, 즉 A의 길이는 40byte (메모리칸에 40byte 크기가 저장된다)
// string(문자열)도 배열이다. 예를들어, s := "Hello World"는 11byte의 배열이다.
// Go = UTF-8을 사용한다 - 문자의 길이는 1~3byte
// 즉, s := "hello world" 는 s := [10]string 과 같다.
