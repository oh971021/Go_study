package main

import "fmt"

type overHeatError float64

// 매개변수 error를 까보면 error 인터페이스로 구성되어있는데,
// 그 친구는 Error() 라는 메소드를 가지고 있다.
// 그래서 요 친구만 살짝 손봐주면 커스텀 에러를 생성할 수 있다!

// 나는 에러에 정답이 나오도록 커스텀해보도록 한다!
func (o overHeatError) Error() string {
	return fmt.Sprintf("over heat is : %0.2f", float64(o))
}

func checkTemperatur(actual float64, criteria float64) error {
	excess := actual - criteria
	if excess > 0 {
		return overHeatError(excess)
	}
	return nil
}

func main() {
	err := checkTemperatur(38.5, 37.5)
	if err != nil {
		fmt.Println(err)
	}
}
