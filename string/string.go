package main

import (
	"fmt"
	"strings"
)

// strings 패키지안에 ToUpper 를 사용하면 간단하게 대문자로 변경가능하다. ( 참고 : https://golang.org/pkg/strings/#ToUpper )

func ToUpper1(str string) string { // 합산을 통해 소문자를 대문자로 바꾸는 함수
	var rst string
	for _, c := range str { // c = Hello world 를 나열해놓음.
		if c >= 'a' && c <= 'z' { // c = 99 , a = 97 , z = 122 / 소문자일 시 if문이 실행 된다.
			rst += string('A' + (c - 'a')) // A(65) + (c(99) - a(97)) A에서 얼마만큼 떨어진 대문자인지를 확인한다.
		} else {
			rst += string(c) // 그대로 반환한다.
		}
	}
	return rst // 하나하나 계산해서 rst로 대입
}

// 합산을 하면 rst += 'H' / rst += 'e' ...... 쭉 계산한다.
// 메모리 공간은 한 글자마다 새로운 공간을 할당해서 값을 복사한다.
// H 	- H 를 담은 메모리가 생긴다.
// HE   - H를 담은 메모리는 사라지고 새로운 공간을 할당해서 H를 복사하고, E를 대입시킨다.
// HEL   - HE를 담은 메모리는 사라지고, 새로운 공간을 할당해서 HE를 복사하고, L을 대입시킨다.

func ToUpper2(str string) string { // WriteRune 모듈을 써서 한글자씩 바꾸는 함수
	var builder strings.Builder
	for _, c := range str { // Hello world 를 한글자씩 c에 대입시킨다.
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a')) // Hello world 를 한글자 씩 써서 빌드해놓은 메모리공간에 대입시킨다.
		} else {
			builder.WriteRune(c) // 한글자씩 그대로 대입
		}
	}
	return builder.String()
}

// strings 패키지에 builder 는 slice 성질을 가진다.
// 미리 공간을 할당해놓고 글자가 대입되더라도 정해진 메모리공간만 사용해서 효율적이다.

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str)) // str로 넣어서 rst로 반환되고 변환된 값이 출력된다.
	fmt.Println(ToUpper2(str))
}
