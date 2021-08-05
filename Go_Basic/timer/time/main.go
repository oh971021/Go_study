// time = struct 안에 종속 된 메소드를 가지고 있다.
// 대부분 값을 리시버로 받는 함수가 대다수
/*
1. 언제 값타입을 쓰고, 포인터타입을 쓸까?
	- 필드값이 바껴도 개념적 실체는 변하지 않는 경우 포인터를 주로 사용한다.
	- 필드값이 변할 때 개념적 실체가 변할 경우 값을 리시버로 주로 사용한다.

2. Time 과 Timer 의 차이 ????
	- Time : 시각을 나타냄 (필드값이 변할 때 시각이라는 개념적 실체가 변함 => 값을 사용)
	- Timer : ex) 알람이라는 실체에 10분 알람, 20분 알람 등등 필드값만 변함 => 포인터 사용)
*/

// Go = 생성자와 소멸자가 없음. 그래서 struct를 만들면 필드만 정의한다.
package main

import (
	"fmt"
	"time"
)

func main() {
	var t1 = time.Timer{}
	var t2 = time.NewTimer(time.Second)
	fmt.Println(t1)
	fmt.Println(t2)
}
