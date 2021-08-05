package main

import (
	"fmt"
	"os"
)

// 별칭 타입 선언 ( 요건 함수타입 )
type Writer func(string)

// write 메소드를 가진 인터페이스 선언 ( 인터페이스로도 사용 가능하다 )
// type WriterInterface interface {
// 	write(string)
// }

// 함수타입을 가르키고, 그 함수타입의 함수를 실행함.
// 의존성 주입 ( 로직 내에서 주입하는 것 )
func writerHello(writer Writer) {
	writer("Hello World")
}

// 이렇게 인터페이스로 주입할 수도 있다.
// func writerHello2(writer WriterInterface) {
// 	writer.write("Hello World")
// }

func main() {
	// 요건 아까 봤던 os 자원 사용하는 거
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	// 프로그램 끝나기전 꺼주겠다는 뜻.
	defer f.Close()

	// 함수를 호출하는데, 함수 리터럴을 넣어준다.
	// 어떤 메세지가 입력이 되면 Fprintln 을 통해서 file 에 메세지를 옮겨 넣겠다.
	writerHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
