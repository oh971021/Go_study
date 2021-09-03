/*
  - Error 처리 구문
   - 에러를 예외 처리해서 에러가 뜨는 경우 예외값을 출력하는 방식.
   - 에러와 예외의 차이점은 회피할 수 있냐 없냐로 분류 된다.
   - 에러와 예외의 공통점은 발생하면 프로그램 정상진행에 문제가 생겨 프로그램이 종료된다
   - 과거의 Error 처리는 if 문으로 처리했다.
     ex ) if A != nil { return A } else { return err }
   - 그래서 Try Catch 문이 나왔다.
     - 에러가 일어날 가능성이 있으면 try ,
	 - try의 처리는 catch문으로 처리한다.

   - 하지만 Go 는 SEH를 지원하지 않는다.
	- 즉, Try Catch 문을 지원하지 않는다.
	- SEH는 성능문제가 있고, 에러를 먹어버리는 문제가 있다.
*/

// Go 에서는 Error도 항상 반환한다..
package main

import (
	"fmt"
	"time"
)

// go에서는 error를 함께 반환한다. 항상.
// 보통 대부분의 함수들은 로직 뒤에 에러도 반환하게 된다.
// 만약 에러가 안일어난다면 nil을 일어난다면 에러를 반환한다.
func div(num1, num2 float32) (float32, error) {
	if num2 != 0 {
		// 만약 num 이 0 이 아니면 정상진행 하고,
		return num1 / num2, nil
	} else {
		// num 이 0 일 때, Error()메소드를 객체로 접근해서
		// 에러발생 시각과 원인을 return 해준다.
		return 0, &MyError{time.Now(), "You Can't Divide by Zero"}
	}
}

// 에러처리를 커스텀 에러로 만들어서 기록할 수 있다.
type MyError struct {
	When time.Time
	What string
}

func (m *MyError) Error() string {
	// 에러가 발생하면 언제, 어떤 에러가 발생했는지 문자열형태로 반환하는 메소드
	// 에러가 발생하면 해당 Error메소드가 Error 를 만들어준다.
	// 아래 빌트인 설명과 함께 보는 것이 좋다.
	return fmt.Sprintf("When : %v\nWhat : %s\n", m.When, m.What)
}

// Go - Builtin 함수로 Error 인터페이스가 이미 정의 되어있다. 그 모양은
/*
type error interface {
	Error() string
}
*/
// 그래서 커스텀에러를 만들고, Error 메소드를 만들어서 접근하면 인터페이스처럼 활용할 수 있다.

func main() {
	var num1 float32 = 4
	var num2 float32 = 0

	result, err := div(num1, num2)
	// 만약 에러가 나면 에러를 출력한다.
	if err != nil {
		fmt.Println(err)
	} else {
		// 에러가 나지 않고 div 함수에서 잘 처리되어 돌아오면
		// 값을 출력한다.
		fmt.Println(result)
	}
}
