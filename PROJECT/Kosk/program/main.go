package main

import (
	"fmt"
)

type RawLearning interface {
	String() string
}

type Raw interface {
	MakeRaw() RawLearning
}

type AIprogram struct {
	val string
}

func (d *AIprogram) PlusRaw(raw Raw) {
	AI := raw.MakeRaw()
	d.val += AI.String()
}

func (d *AIprogram) String() string {
	return "AI Learned " + d.val
}

type RawOne struct {
}

func (f *RawOne) MakeRaw() RawLearning {
	return &RawOneLearnig{}
}

type RawOneLearnig struct {
}

func (s *RawOneLearnig) String() string {
	return "RawOne"
}

type RawTwo struct {
}

func (f *RawTwo) MakeRaw() RawLearning {
	return &RawTwoLearning{}
}

type RawTwoLearning struct {
}

func (s *RawTwoLearning) String() string {
	return " + RawTwo"
}

type RawThree struct {
}

func (f *RawThree) MakeRaw() RawLearning {
	return &RawThreeLearning{}
}

type RawThreeLearning struct {
}

func (s *RawThreeLearning) String() string {
	return " + RawThree"
}

func main() {
	fmt.Println("Raw Learning Program")

	fmt.Println()
	fmt.Println("AI is Learning Raw")

	fmt.Println()
	// AI가 법을 배웁니다 //
	AI := &AIprogram{}
	gwanga := &RawOne{}
	salmon := &RawThree{}
	sekkoshi := &RawTwo{}

	AI.PlusRaw(gwanga)
	AI.PlusRaw(salmon)
	AI.PlusRaw(sekkoshi)

	fmt.Println(AI)
	// =================== //
}
