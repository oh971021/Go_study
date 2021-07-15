package main

import "fmt"

//=== 사시미 제작 과정에 필요한 인터페이스 및 객체 ====//
type SliceSashimi interface {
	String() string
}

type Sashimi interface {
	GetOneDish() SliceSashimi
}

type BusanSashimi struct { //Dish라는 객체를 만듬
	val string
}

//=====================================================//

// 접시에 사시미를 얹는 작업
func (d *BusanSashimi) OnSashimi(sashimi Sashimi) {
	dish := sashimi.GetOneDish()
	d.val += dish.String()
}

// 접시에 썰어놓은 사시미를 하나씩 얹는 작업
func (d *BusanSashimi) String() string {
	return "GatheringDshi = " + d.val
}

// ============ 광어회떠서 접시에 얹는 작업 ============ //
type GwangaSashimi struct { // 광어회 객체 생성
}

func (f *GwangaSashimi) GetOneDish() SliceSashimi { // 광어회 객체를 접시에 얹는 작업
	return &SliceGwangaSashimi{}
}

type SliceGwangaSashimi struct { // 광어회 뜨는 객체 생성
}

func (s *SliceGwangaSashimi) String() string { // 접시에 얹기전 이 메소드를 통해 썰어서 얹는다.
	return "GwangaSashimi"
}

// ====================================================  //

// ============ 세꼬시떠서 접시에 얹는 작업 ============ //
type SekkoshiSashimi struct { // 세꼬시 객체 생성
}

func (f *SekkoshiSashimi) GetOneDish() SliceSashimi { // 세꼬시 객체를 접시에 얹는 작업
	return &SliceSekkoshiSashimi{}
}

type SliceSekkoshiSashimi struct { // 세꼬시 뜨는 객체 생성
}

func (s *SliceSekkoshiSashimi) String() string { // 접시에 얹기전 이 메소드를 통해 썰어서 얹는다.
	return " + SekkoshiSashimi"
}

// ====================================================  //

// ============ 연어회떠서 접시에 얹는 작업 ============ //
type SalmonSashimi struct { // 세꼬시 객체 생성
}

func (f *SalmonSashimi) GetOneDish() SliceSashimi { // 연어회 객체를 접시에 얹는 작업
	return &SliceSalmonSashimi{}
}

type SliceSalmonSashimi struct { // 연어회 뜨는 객체 생성
}

func (s *SliceSalmonSashimi) String() string { // 접시에 얹기전 이 메소드를 통해 썰어서 얹는다.
	return " + SalmonSashimi"
}

// ====================================================  //

//=== 지역 선정 과정에 필요한 인터페이스 및 객체 ====//
type GoToPlace interface {
	String() string
}

type Place interface {
	ChoicePlace() GoToPlace
}

type PlaceinBusan struct { // 부산 내에 있는 장소 객체 생성
	van string
}

//=====================================================//

// 선정한 지역으로 이동한다.
func (d *PlaceinBusan) GoingPlace(place Place) {
	user := place.ChoicePlace()
	d.van += user.String()
}

// 지역 선정
func (d *PlaceinBusan) String() string {
	return "This is " + d.van
}

// ================== 광안리 지역 선정 및 이동 ================= //

type Gwanganri struct { // 광안리 객체 생성
}

func (f *Gwanganri) ChoicePlace() GoToPlace { // 광안리 객체를 선정해서 이동하는 작업
	return &GotoGwanganri{}
}

type GotoGwanganri struct { // 광안리로 이동하는 객체
}

func (s *GotoGwanganri) String() string { // 이동 전 선정해서 이동한다.
	return "Gwanganri SuSanSiJang"
}

// ============================================================== //

// ================== 해운대 지역 선정 및 이동 ================= //

type Haeundae struct { // 해운대 객체 생성
}

func (f *Haeundae) ChoicePlace() GoToPlace { // 해운대 객체를 선정해서 이동하는 작업
	return &GotoHaeundae{}
}

type GotoHaeundae struct { // 해운대로 이동하는 객체
}

func (s *GotoHaeundae) String() string { // 이동 전 선정해서 이동한다.
	return "Haeundae SuSanSiJang"
}

// ============================================================== //

func main() {
	fmt.Println("Welcome To Busan")
	fmt.Println()
	// 지역 선정 빌드업 //
	place := &PlaceinBusan{}
	gwanganri := &Gwanganri{}
	// haeundae := &Haeundae{}

	place.GoingPlace(gwanganri)
	// place.GoingPlace(haeundae)

	fmt.Println(place)
	// ================ //
	fmt.Println()
	fmt.Println("You ordered GatheringDshi")
	// 모듬회 생성 빌드 업 //
	dish := &BusanSashimi{}
	gwanga := &GwangaSashimi{}
	salmon := &SalmonSashimi{}
	sekkoshi := &SekkoshiSashimi{}

	dish.OnSashimi(gwanga)
	dish.OnSashimi(salmon)
	dish.OnSashimi(sekkoshi)

	fmt.Println(dish)
	// =================== //
}
