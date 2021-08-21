package main

import "fmt"

func main() {
	one := "Golang"
	two := "is Awsome"

	strSlice := []string{}
	strSlice = append(strSlice, one)
	fmt.Println(strSlice)
	strSlice = append(strSlice, two)
	fmt.Println(strSlice)
}
