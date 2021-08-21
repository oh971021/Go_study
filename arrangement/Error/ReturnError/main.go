/*
  ## 사용자 에러 반환 ##
   - 에러가 만드는 함수를 사용해서 사용자의 커스텀 에러를 생성해준다.
   - Errorf(formatter string, ...interface{}) error {}
   - errors.New(errormessage string) error {}
*/

package main

import (
	"fmt"
	"math"

	// 에러를 생성하는 패키지
	"errors"
)

// 제곱근 함수
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// 제곱근이 음수면 처리할 수 없다.
		// 새로운 에러를 생성해준다.
		return 0, errors.New("제곱근은 양수여야 합니다")
	}
	// 제곱근을 구할 값이 양수면
	// Math 패키지에 Sqrt(제곱근)으로 매개변수를 계산해준다.
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
