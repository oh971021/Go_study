package main

import "fmt"

type DishOfSlice interface { //외부 공개 메서드 = 관계
	String() string //PutOfPatty은 String()이라는 외부 공개 메서드
}

type Slice interface { //오로지 관계만 선언해본다
	GetOneDish() DishOfSlice //Jam은 GetOnsSpoon()이라는 외부 공개 메서드
}

type SaShimi struct { //Bread 객체를 만든다
	val string
}

func (b *SaShimi) NetPutSashimi(slice Slice) { //오렌지잼이든 스토우베리 잼이든 상관없다.
	dish := slice.GetOneDish()
	b.val += dish.String()
}

func (b *SaShimi) String() string { //Bread 메서드2 String
	return "#### DongDeaguStationSashimiWorld ####\nIt's a popular menu at our store : Dish [ " + b.val + " ]"
}

type SeKKoShi struct {
}

func (j *SeKKoShi) GetOneDish() DishOfSlice {
	return &DishOfSeKKoShi{}
}

// 세꼬시 손질
type DishOfSeKKoShi struct {
}

func (s *DishOfSeKKoShi) String() string {
	return " + sekkoshi"
}

type Gwanga struct {
}

func (j *Gwanga) GetOneDish() DishOfSlice {
	return &DishOfGwanga{}
}

// 광어 손질
type DishOfGwanga struct {
}

func (s *DishOfGwanga) String() string {
	return " + gwanga"
}

// 연어 손질
type Salmon struct {
}

func (j *Salmon) GetOneDish() DishOfSlice {
	return &DishOfSalmon{}
}

type DishOfSalmon struct {
}

func (s *DishOfSalmon) String() string {
	return "salmon"
}

func main() {
	SaShimi := &SaShimi{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	slice1 := &Salmon{}
	slice2 := &SeKKoShi{}
	slice3 := &Gwanga{}
	SaShimi.NetPutSashimi(slice1)
	SaShimi.NetPutSashimi(slice2)
	SaShimi.NetPutSashimi(slice3)

	fmt.Println(SaShimi)
}
