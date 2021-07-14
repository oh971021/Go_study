package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 햄버거 인터페이스 생성
type MakeBugger interface {
	String() string
}

// 재료 인터페이스 생성 == 왜 인터페이스로 접근할까요?
type Meterial interface {
	GetOnMeterial() MakeBugger
}

// 빵 객체 생성
type Bread struct {
	val string
}

// 메뉴셋팅 인터페이스 생성
type MakeMenu interface {
	Setting() string
}

// 세트 메뉴 인터페이스 생성
type Menus interface {
	PulsMenu() MakeMenu
}

// 셋팅되는 부가 메뉴들
type Set struct {
	me string
}

// 빵을 놓는다.
func (b *Bread) PutMeterial(meterial Meterial) {
	shief := meterial.GetOnMeterial()
	b.val += shief.String()
}

// 빵과 빵 사이에 재료를 넣는다.
func (b *Bread) String() string {
	return "참깨빵" + b.val + " + 참깨빵"
}

// 햄버거를 제외한 셋트메뉴를 추가한다.
func (m *Set) PulsMenus(menus Menus) {
	setmenu := menus.PulsMenu()
	m.me += setmenu.Setting()
}

// 셋트메뉴를 하나씩 셋팅한다.
func (m *Set) Setting() string {
	return m.me
}

// 셋트메뉴에 들어갈 감자튀김 객체 생성
type Poteto struct {
}

// 감자튀김을 꺼낸다.
func (j *Poteto) PulsMenu() MakeMenu {
	return &MakePoteto{}
}

// 아메리카노 객체 생성
type Americano struct {
}

// 아메리카노를 꺼낸다.
func (j *Americano) PulsMenu() MakeMenu {
	return &MakeAmericano{}
}

// 패티 객체 생성
type Patty struct {
}

// 패티를 꺼낸다.
func (j *Patty) GetOnMeterial() MakeBugger {
	return &PutOnPatty{}
}

// 치즈 객체 생성
type Cheese struct {
}

// 치즈를 꺼낸다.
func (j *Cheese) GetOnMeterial() MakeBugger {
	return &PutOnCheese{}
}

// 소스 객체 생성
type Sauce struct {
}

// 소스를 꺼낸다.
func (j *Sauce) GetOnMeterial() MakeBugger {
	return &PutOnSauce{}
}

// 양상추 객체 생성
type Lettuce struct {
}

// 양상추를 꺼낸다.
func (j *Lettuce) GetOnMeterial() MakeBugger {
	return &PutOnLettuce{}
}

// 양파 객체 생성
type Onion struct {
}

// 양파를 꺼낸다.
func (j *Onion) GetOnMeterial() MakeBugger {
	return &PutOnOnion{}
}

// 피클 객체 생성
type Pickle struct {
}

// 피클을 꺼낸다.
func (j *Pickle) GetOnMeterial() MakeBugger {
	return &PutOnPickle{}
}

// 패티 메소드 생성
type PutOnPatty struct {
}

// 패티를 얹는다.
func (s *PutOnPatty) String() string {
	return " + 쇠고기패티"
}

// 치즈 메소드 생성
type PutOnCheese struct {
}

// 치즈를 얹는다.
func (s *PutOnCheese) String() string {
	return " + 치즈"
}

// 소스 메소드 생성
type PutOnSauce struct {
}

// 소스를 뿌린다.
func (s *PutOnSauce) String() string {
	return " + 특별한소스"
}

// 양상추 메소드 생성
type PutOnLettuce struct {
}

// 양상추를 얹는다.
func (s *PutOnLettuce) String() string {
	return " + 양상추"
}

// 피클 메소드 생성
type PutOnPickle struct {
}

// 피클을 얹는다.
func (s *PutOnPickle) String() string {
	return " + 피클"
}

// 양파 메소드 생성
type PutOnOnion struct {
}

// 양파를 얹는다.
func (s *PutOnOnion) String() string {
	return " + 양파"
}

// 감자튀김 메소드 생성
type MakePoteto struct {
}

// 감자튀김을 셋트메뉴에 추가한다.
func (s *MakePoteto) Setting() string {
	return " + 감자튀김"
}

// 아메리카노 메소드 생성
type MakeAmericano struct {
}

// 아메리카노 셋트메뉴에 추가
func (s *MakeAmericano) Setting() string {
	return " + 아메리카노"
}

func main() {
	// 햄버거 재료
	bread := &Bread{}
	cheese := &Cheese{}
	patty := &Patty{}
	sauce := &Sauce{}
	lettuce := &Lettuce{}
	onion := &Onion{}
	pickle := &Pickle{}

	// 셋트 메뉴 : 출력값이 이상합니다... @@@@@@@@@@@@@@@@@ 세트 메뉴 : &{ + 감자튀김 + 아메리카노} // &{} 를 없애고 싶습니다. @@@@@@@@@@@@@@@@
	set := &Set{}
	poteto := &Poteto{}
	americano := &Americano{}

	// 주문시간 측정
	maketime := time.Now()

	// 주문번호 안겹치게 시드 생성
	rand.Seed(maketime.UnixNano())

	print("[ 동대구 맥도날드 ]")
	fmt.Println("주문 시간은", maketime)

	// 주문번호 생성
	r := rand.Intn(100000)
	fmt.Printf("주문 번호는 %v 입니다.\n", r)

	// 햄버거 만드는 시간
	time.Sleep(3 * time.Second)

	// 재료 넣는 중
	bread.PutMeterial(patty)
	bread.PutMeterial(patty)
	bread.PutMeterial(sauce)
	bread.PutMeterial(lettuce)
	bread.PutMeterial(cheese)
	bread.PutMeterial(pickle)
	bread.PutMeterial(onion)
	fmt.Println("햄버거 :", bread)

	// 세트메뉴 준비중
	set.PulsMenus(poteto)
	set.PulsMenus(americano)
	fmt.Println("세트 메뉴 :", set)

	// 완성
	endtime := time.Since(maketime)
	fmt.Println("햄버거가 완성됐습니다.")
	fmt.Printf("햄버거 완성까지 %v 걸렸습니다.", endtime)
}
