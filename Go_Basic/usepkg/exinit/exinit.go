package exinit

import (
	"fmt"
	// "tutorial/usepkg/custompkg" import cycle not allowed (서로 동시에 임포트 할수는 없다)
)

var (
	a = c + b // 변수의 값을 정해줘야 한다. 슌서대로 실행해본다. 다 확인하면 a는 9가 된다.
	b = f()
	c = f() // 확인을 위해 f를 실행해본다
	d = 3
)

func init() { // 보조 패키지안에 init 함수가 있으면 실행을 하고 보낸다.
	d++
	fmt.Println("exinit.init function", d)
} // 패키지 초기화의 끝은 init

func f() int { // c = 4 가 된다. 그다음 b가 오는데 d는 c에서 한번 4로 초기화되서 b는 5가 된다.
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	custompkg.PrintCustom()
	fmt.Println("d:", d)
}

// 패키지 초기화는 한번만 호출 된다.
// 지금 usepkg, custompkg 에서 두번 호출 되었지만 실행은 한번만 한다.
