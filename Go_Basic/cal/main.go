package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("첫 번째 숫자를 입력하세요.")

	// 버퍼 공간을 만들어준다.
	reader := bufio.NewReader(os.Stdin)
	// 엔터친 부분을 읽겠다.
	line, _ := reader.ReadString('\n')
	// 공백값을 없애주는 트림 스페이스
	line = strings.TrimSpace(line)
	// 문자열 변환기 (문자열을 인트로 바꿔준다)
	num, _ := strconv.Atoi(line)

	fmt.Println("두 번째 숫자를 입력하세요.")

	reader01 := bufio.NewReader(os.Stdin)
	line01, _ := reader01.ReadString('\n')
	line01 = strings.TrimSpace(line01)
	num01, _ := strconv.Atoi(line01)

	fmt.Println("연산자를 입력하세요.")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)
	result := 0

	// 	if op == "+" {
	// 		result = num + num01
	// 		fmt.Printf("%v + %v = %v", num, num01, result)
	// 	} else if op == "-" {
	// 		result = num - num01
	// 		fmt.Printf("%v - %v = %v", num, num01, result)
	// 	} else if op == "*" {
	// 		result = num * num01
	// 		fmt.Printf("%v * %v = %v", num, num01, result)
	// 	} else if op == "/" {
	// 		result = num / num01
	// 		fmt.Printf("%v / %v = %v", num, num01, result)
	// 	} else {
	// 		fmt.Println("연산자가 아닙니다.")
	// 	}
	switch op {
	case "+":
		result = num + num01
		fmt.Printf("%v + %v = %v", num, num01, result)
	case "-":
		result = num - num01
		fmt.Printf("%v - %v = %v", num, num01, result)
	case "*":
		result = num * num01
		fmt.Printf("%v * %v = %v", num, num01, result)
	case "/":
		result = num / num01
		fmt.Printf("%v / %v = %v", num, num01, result)
	default:
		fmt.Println("연산자가 아닙니다.")
	}
}
