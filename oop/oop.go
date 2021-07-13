package main

import (
	"crypto/rand"
	"fmt"
)

type LottoNumbers struct {
	Number int
}

func MakeLottoNumbers(*LottoNumbers) {
	MakeNum = rand.Int(0, 50)
	return MakeNum
}

func main() {
	var LottoNumber []LottoNumbers

	fmt.Println(LottoNumber)
}
