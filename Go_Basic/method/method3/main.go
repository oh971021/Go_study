package main

import "fmt"

// 유저 객체 생성
type User struct {
	Name string
	Age  int
}

// 유저를 리시버로 받는
func (u User) String() string {
	// Sprint : 스트링으로 변환시켜서 프린트 함.
	return fmt.Sprintf("%s, %d", u.Name, u.Age)
}

type VIPUser struct {
	User     // embedded field
	VIPLevel int
}

func (v VIPUser) vipLevel() int {
	return v.VIPLevel
}

// 요 친구가 있으면 viplevel 을 출력한다.
func (v VIPUser) String() string {
	return fmt.Sprintf("%d", v.VIPLevel)
}

func main() {
	vip := VIPUser{User{"준석", 25}, 10}
	fmt.Println(vip.String()) // 임베디드 필드는 따로 적어주지 않아도 포함된다.
}
