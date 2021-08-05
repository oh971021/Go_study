package main

import (
	"fmt"
	"tutorial/usepkg/custompkg" // 보조 패키지를 가져와서 사용하겠다.
	"tutorial/usepkg/exinit"

	"github.com/guptarohit/asciigraph"           // 외부 패키지를 가져와서 사용하겠다.
	"github.com/tuckersGo/musthaveGo/ch16/expkg" // go mod tidy 명령어로 외부 패키지를 다운해야 사용가능하다.
)

// 외부, 보조 패키지를 가지고와서 메인에서 실행하여 프로그램으로 만든다.

func main() {
	custompkg.PrintCustom() // custompkg 패키지의 프린트 커스텀 함수를 호출한다.

	s := custompkg.Student{}
	s.Name = "화랑"
	s.Age = 32
	// s.score = 100 외부로 공개되는 스트럭쳐라도 메소드가 비공개면 사용할 수 없다.

	fmt.Println(s.Name, s.Age)
	fmt.Println(custompkg.PI)
	custompkg.Ratio = 10
	fmt.Println(custompkg.Ratio) // 타입, 변수, 상수 등 다 가져와서 사용할 수 있다.

	// custompkg.printcustom2() 외부로 공개되지 않은 패키지를 사용하면 unexported name이라는 에러가 뜬다.
	expkg.PrintSample() // 외부 패키지를 호출한 것

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6, 5, 4, 3, 2, 1, 4, 5, 7, 5, 3, 3, 2}
	graph := asciigraph.Plot(data) // 아스키 패키지의 기능을 활용해서 집어넣은 데이터로 그래프를 출력한다.
	fmt.Println(graph)

	exinit.PrintD() // 패키지를 import하면 변수값들을 초기화하고 init함수가 있으면 실행하고 메인함수로 리턴한다.
}
