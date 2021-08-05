/* method ??
- 메서드는 타입에 속한 함수다. (타입에 종속적)
- 함수와 같은 형태를 가지고 있지만 리시버를 가지고 있다는 점이 다르다.
- 리시버 : 모든 패키지 지역 타입이 가능하다.(구조체 별칭 등)
- 매개변수를 밖으로 빼낸 모양.
ex . func (a apple) function() int { return r.width * r.height }
	=> 함수는 리시버 타입이다.
*/
package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { // function
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // method 함수와의 차이는 리시버가 있는것.
	a.balance -= amount
}

func main() {
	a := &account{100} // *account

	withdrawFunc(a, 30)

	a.withdrawMethod(30) // a라는 객체를 받는 메소드를 호출.

	fmt.Printf("%d\n", a.balance)
}
