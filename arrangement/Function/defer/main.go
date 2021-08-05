/*
 == defer 지연 실행 ==
	- 함수 종료 전 실행을 보장한다.
	- 함수 내에 기능을 미리 작성해놓지만, 함수가 끝나는 시점에서 실행되도록 할 때 사용한다.

	- 주로 os 자원을 반납할 때 사용한다. ( os에 자원을 요청해서 사용하면 반드시 반환해야함. )
	- 반환해주지 않으면 프로그램이 끝날 때 까지 가지고 있는다.

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// os 자원에서 create 기능을 가져와서 사용하는데,
	// 이 create는 파일을 반환하거나, error 를 반환한다.
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}

	defer fmt.Println("반드시 호출됩니다.")
	// file 자원이 프로그램 종료 전 반드시 close 를 실행할 수 있도록 defer로 적어줌.
	defer f.Close()
	defer fmt.Println("파일을 닫습니다.")

	fmt.Println("파일에 Hello World를 씁니다.")

	// 생성한 f (텍스트파일) 에 작성 할때는 Fprint 를 사용해준다.
	// Fprint : 첫번째 인자로 io writer 를 받는다.
	fmt.Fprintln(f, "Hello World")

	// => 확인 할 때는 " cat test.txt " 를 터미널에서 작성해준다.
}
