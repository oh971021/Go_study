package main

import "fmt"

type person struct {
	name    string
	age     int
	content string
}

func addAgeRef(a *person) {
	a.age += 4
}

func addAgeVal(a person) {
	a.age += 4
}

func main() {
	var p1 = new(person)
	p2 := person{}

	fmt.Println(p1, p2)

	p1.age = 25
	p1.age = 25

	addAgeRef(p1)
	addAgeVal(p2)

	fmt.Println(p1.age)
	fmt.Println(p2.age)

	addAgeRef(&p2)
	fmt.Println(p2.age)
}
