// 인터페이스 타입변환
package main

import "fmt"

// 점프하는 메소드를 가진 점퍼 인터페이스
type Jumper interface {
	Jump()
}

// 걷는 메소드를 가진 워커 인터퍼이스
type Walker interface {
	Walk()
}

// 객체 하나 생성해줌
type Person struct{}

// 이 객체로 두 인터페이스를 연결고리를 만들어준다.
func (p *Person) Jump() {
	// 변환에 실패하면 런타임 에러가 뜨지만
	// 비어있길래 그냥 적어놓았따.
	fmt.Println("Failed..")
}

func (p *Person) Walk() {
	// 타입변환에 성공하게 되면 보기좋게 하나 출력해준다.
	fmt.Println("Success")
}

// 점프를 받았지만 출력값은 워크로 나올 수 있도록 변환시켜주는 기능
func JumpingMan(jumper Jumper) {
	// jumper.(Walker) => 여기서 인터페이스 타입변환을 실행한다.
	// if 문에 넣어서 잘 들어갔는지 확인하는 구문을 만들어준다.
	// c : 데이터가 들어가는 공간
	// ok : True / False 가 나타나는 변수
	if v, ok := jumper.(Walker); ok {
		// 만약 변환이 잘 되었다면 Walk 메소드를 실행해본다.
		v.Walk()
	}
}

func main() {
	// Person 객체를 담는 변수 생성
	p := &Person{}
	// Person에 포함 된 메소드 중 Jump를 사용하고
	// Jump 에서 Walk로 타입변환 후 출력한다.
	JumpingMan(p)
}
