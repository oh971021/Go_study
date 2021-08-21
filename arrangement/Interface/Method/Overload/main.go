/*
go 언어에서는 함수명이 같을 경우 재선언 오류를 발생시킨다.
다른 메소드를 만들고 오버로드 처리하도록 설계해서 구현해야 한다.

오버로드 처리를 위해 type switch를 사용해서 들어오는 타입에 따라 메소드 호출하는 처리를 사용한다.

함수나 메소드 인자를 다양하게 처리하려면 []interface{} 로 정의해서 다양한 인자를 받아서 오버로딩 처리하면된다.
*/

package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) GetName1() string {
	return p.Name
}

func (p *Person) GetName2(s string) string {
	return p.Name
}

func (p *Person) Overload(args interface{}) string {
	switch args.(type) {
	case string:
		return p.GetName2("Moon")
	case nil:
		return p.GetName1()
	default:
		return "No Match"
	}
}

func main() {
	p := Person{"Junseok"}

	var s string = "moon"

	fmt.Println("OverLoading ", p.Overload(s))
	fmt.Println("OverLoading ", p.Overload(nil))
}
