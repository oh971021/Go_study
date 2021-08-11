package main

import "fmt"

func main() {
	var (
		name  string
		age   int
		phone string
	)

	fmt.Println("이름, 나이, 전화번호")
	fmt.Scan(&name, &age, &phone)
	fmt.Println(name, age, phone)
}
