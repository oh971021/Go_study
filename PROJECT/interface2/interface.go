package main

import "fmt"

type Knifeofsashimi interface {
	String() string
}

type Sashimi interface {
	ToCutSashimi() Knifeofsashimi
}

type Dish struct { //Dish라는 객체를 만듬
	val string
}

func (d *Dish) OnSashimi(sashimi Sashimi) {
	knife := sashimi.ToCutSashimi()
	d.val += knife.String()
}

func (d *Dish) String() string {
	return "dish" + d.val
}

type FlukeSashimi struct { //광어회
}

func (f *FlukeSashimi) ToCutSashimi() Knifeofsashimi {
	return &KnifeofFlukesashimi{}
}

type KnifeofFlukesashimi struct {
}

func (s *KnifeofFlukesashimi) String() string {
	return "+ 광어회"
}

func main() {
	dish := &Dish{}
	sashimi1 := &FlukeSashimi{}

	dish.OnSashimi(sashimi1)

	fmt.Println("--------------[광안리수변공원횟집]--------------")
	fmt.Println(dish)

}
