package main

import (
	"fmt"
)

type Bread struct {
	val string
}

type Jam struct {
	opened bool
}

type SpoonOfJam struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	bread := make([]*Bread, num) // 포인터 배열타입을 만든다.
	for i := 0; i < num; i++ {
		bread[i] = &Bread{val: "빵"}
	}
	return bread
}

func OpenJam(jam *Jam) {
	jam.opened = true
}

func GetOneSpoon(_ *Jam) *SpoonOfJam {
	return &SpoonOfJam{}
}

func PutJamOnBreadStrawBerry(bread *Bread, jam *SpoonOfJam) {
	bread.val += " + 딸기잼"
}

func PutJamOnBreadOrangeJam(bread *Bread, jam *SpoonOfJam) {
	bread.val += " + 오렌지잼"
}

func PutJamOnBreadPeanutsJam(bread *Bread, jam *SpoonOfJam) {
	bread.val += " + 땅콩잼"
}

func MakeSandwitch(bread []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(bread); i++ {
		sandwitch.val += bread[i].val + " + "
	}
	return sandwitch
}

func main() {
	// 빵을 두개 꺼낸다.
	breads := GetBreads(2)

	// 스트로베리 잼을 바른다.
	jam := &Jam{}

	// 딸기잼 뚜겅을 연다
	OpenJam(jam)

	// 한스푼 뜬다
	spoon := GetOneSpoon(jam)

	/*
		// 딸기잼을 바른다.
		PutJamOnBreadStrawBerry(breads[0], spoon)

		// 오렌지잼을 바른다.
		PutJamOnBreadOrangeJam(breads[0], spoon)

		// 땅콩잼을 바른다.
		PutJamOnBreadPeanutsJam(breads[0], spoon) */

	var rsv int
	fmt.Println("0 = 딸기, 1 = 오렌지, 2 = 땅콩버터, 원하는 잼을 골라보세요")
	fmt.Scanln(&rsv)

	// 숫자에 따른 잼을 발라준다.
	if true { // 잼을 바르게 된다면
		if rsv == 0 {
			PutJamOnBreadStrawBerry(breads[0], spoon)
			sandwitch := MakeSandwitch(breads)
			fmt.Println(sandwitch.val)
		} else if rsv == 1 {
			PutJamOnBreadOrangeJam(breads[0], spoon)
			sandwitch := MakeSandwitch(breads)
			fmt.Println(sandwitch.val)
		} else if rsv == 2 {
			PutJamOnBreadPeanutsJam(breads[0], spoon)
			sandwitch := MakeSandwitch(breads)
			fmt.Println(sandwitch.val)
		} else {
			fmt.Println("잼이 없습니다")
		}
		fmt.Println("샌드위치를 완성했습니다.") // 결과값을 출력한다.
	}

	/*
		// 빵을 덮는다.
		sandwitch := MakeSandwitch(breads)

		// 완성
		fmt.Println(sandwitch.val) */
}
