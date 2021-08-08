package main

import (
	"banker"
	"fmt"
)

type Bank struct {
	Name  string
	Money int
}

// New returns Bank struct
// Bank 타입으로 리턴하는데, 참조할 수 있도록 *을 붙혀준다.
func New(name string) *Bank {
	return &Bank{name: name}
}

func (bank *Bank) Deposit(money int) {
	bank.Money += money
}

func main() {
	bank := banker.New("someone")
	bank.Deposit(3000)
	fmt.Println(bank)
}
