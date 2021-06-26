package custompkg // main이 아닌 패키지는 프로그램이 아니라 프로그램을 보조해주는 보조패키지이다.

import (
	"fmt"
	"tutorial/usepkg/exinit"
)

type Student struct { // 타입도 똑같이 앞글자가 대문자면 다른 패키지에서 사용이 가능하다.
	Name  string
	Age   int
	score int
}

var Ratio int // 전역변수도 앞글자가 대문자면 다른 패키지에서 사용이 가능하다.
var radio int

const PI = 3.14 // 상수도 앞글자가 대문자면 다른 패키지에서 사용이 가능하다.
const pI2 = 3.1415

func PrintCustom() {
	fmt.Println("This is custom package!")
	exinit.PrintD() // 메인 패키지가 아니여도 외부 패키지를 임폿가능하다.
}

// 앞글자가 대문자면 다른 패키지에서 사용이 가능하다. ( 공개형 )

func printcustom2() {
	fmt.Println("This is custom package222!")
}

// 앞글자가 소문자면 다르패키지에서 사용이 불가능하다. ( 비공개형 )

// 함수명을 한글로 쓰면 안되는 이유 = 공개인지 비공개인지 알 수가 없어서 영어로 지정해주고,
// 공개 비공개를 확실하게 적어준다.

// 값을 출력하는 함수를 가진 패키지가 만들어졌다.
