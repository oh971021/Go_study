package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Friend []Person

func (s Friend) Len() int           { return len(s) }
func (s Friend) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Friend) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	p := []Person{
		{"준석", 25},
		{"민영", 28},
		{"수용", 26},
		{"배훈", 23},
		{"성민", 25},
	}

	fmt.Println("Before sort : ", p)
	sort.Sort(Friend(p))
	fmt.Println("After sort : ", p)
}
