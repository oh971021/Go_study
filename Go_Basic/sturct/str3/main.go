package main

import "fmt"

type Bank struct {
	Name  string
	Money int
}

func main() {
	bankA := Bank{"Junseok", 50000}
	fmt.Println(bankA)

	bankB := Bank{Money: 20000, Name: "DongChan"}
	fmt.Println(bankB)
}
