package main

import "fmt"

// ####################### 모듬회 준비 ######################## //

// 회뜰때 활용할 나이프.
type MakeSashimi interface {
	Dish() string
}

// 모듬회에 들어갈 물고기를 가져온다.
type GatheringSashimi interface {
	GetOneDish() MakeSashimi
}

// 사시미가 될 객체 생성
type Sashimi struct {
	sashimi string
}

// 셰프가 물고기를 가져와서 회를 뜬다.
func (s *Sashimi) SliceSashimi(gatheringsashimi GatheringSashimi) {
	shef := gatheringsashimi.GetOneDish()
	s.sashimi += shef.Dish()
}

// 접시에 회를 얹는다.
func (s *Sashimi) Dish() string {
	return "plat " + s.sashimi
}

// ------------ 세꼬시 회 추가 ------------ //
type Sekkoshi struct {
}

func (k *Sekkoshi) GetOneDish() MakeSashimi {
	return &PutOnSekkoshi{}
}

type PutOnSekkoshi struct {
}

func (p *PutOnSekkoshi) Dish() string {
	return " + Sekkoshi"
}

// ---------------------------------------- //

// -------------- 연어 회 추가 ------------ //
type Salmon struct {
}

func (k *Salmon) GetOneDish() MakeSashimi {
	return &PutOnSekkoshi{}
}

type PutOnSalmon struct {
}

func (p *PutOnSalmon) Dish() string {
	return " + Salmon"
}

// ---------------------------------------- //

// -------------- 광어 회 추가 ------------ //
type Gwaga struct {
}

func (k *Gwaga) GetOneDish() MakeSashimi {
	return &PutOnGwaga{}
}

type PutOnGwaga struct {
}

func (p *PutOnGwaga) Dish() string {
	return " + Gwaga"
}

// ---------------------------------------- //

// -------------- 우럭 회 추가 ------------ //
type WooLuck struct {
}

func (k *WooLuck) GetOneDish() MakeSashimi {
	return &PutOnWooLuck{}
}

type PutOnWooLuck struct {
}

func (p *PutOnWooLuck) Dish() string {
	return " + WooLuck"
}

// ---------------------------------------- //
// ###################################################### //

// ###################### 지역 선정 ##################### //
type ChoicePlace interface {
	Choice() string
}

type GotoPlace interface {
	PlaceinBusan() ChoicePlace
}

type Place struct {
	place string
}

func (p *Place) SelectPlace(gotoplace GotoPlace) {
	people := gotoplace.PlaceinBusan()
	p.place += people.Choice()
}

func (p *Place) Choice() string {
	return "This is = " + p.place
}

// -------------- 지역 광안리 추가 ------------ //
type Gwanganri struct { // 광안리 객체 생성
}

func (g *Gwanganri) PlaceinBusan() ChoicePlace {
	return &GotoGwanganri{}
}

type GotoGwanganri struct {
}

func (g *GotoGwanganri) Choice() string {
	return "Gwanganri"
}

// -------------------------------------------- //

// -------------- 지역 광안리 추가 ------------ //
type Haeundae struct { // 해운대 객체 생성
}

func (g *Haeundae) PlaceinBusan() ChoicePlace {
	return &GotoHaeundae{}
}

type GotoHaeundae struct {
}

func (g *GotoHaeundae) Choice() string {
	return "Haeundae"
}

// -------------------------------------------- //

// ###################################################### //

func main() {

	// --------- 지역선택 ---------- //
	place := &Place{}

	gwanganri := &Gwanganri{}
	// haeundae := &Haeundae{}
	place.SelectPlace(gwanganri)
	// place.SelectPlace(haeundae)
	fmt.Println(" 어서오세요 ", place, " 횟집입니다.")
	// ----------------------------- //

	// -------- 모듬회 생성 ------- //
	sashimis := &Sashimi{}

	sekkoshi := &Sekkoshi{}
	salmon := &Salmon{}
	gwaga := &Gwaga{}
	wooluck := &WooLuck{}

	sashimis.SliceSashimi(sekkoshi)
	sashimis.SliceSashimi(salmon)
	sashimis.SliceSashimi(gwaga)
	sashimis.SliceSashimi(wooluck)

	fmt.Println(sashimis)
	// ----------------------------- //

}
