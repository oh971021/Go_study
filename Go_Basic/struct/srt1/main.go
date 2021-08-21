/*
  == 구조체란 ?? ==
	- 하나 이상의 변수를 묶어서 새로운 자료형을 정의하는 Custom Data Type 이다.
	- 각각의 선언 된 변수는 필드가 되고, 필드들의 집합체라고 할 수 있다.
	- 구조체 객체를 생성하는 법 : 객체이름 := 구조체이름{필드에 저장할 값} 을 입력한다.
	- 빈객체로 생성을 한 경우 : 객체이름.필드이름 = 필드에 저장할 값 을 입력한다.
	- 값 입력이 생략된 필드는 " Zero Value "" 로 나온다.
*/

package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Content string
}

func main() {
	p := Person{}
	fmt.Println(p)

	p.Name = "JunBal"
	p.Age = 25
	p.Content = "01022153470"
	fmt.Println(p)

	p = Person{"Junseok", 26, "01022153470"}
	fmt.Println(p)

	p = Person{Age: 30, Name: "JunSeok", Content: "010-2215-3470"}
	fmt.Println(p)
}
