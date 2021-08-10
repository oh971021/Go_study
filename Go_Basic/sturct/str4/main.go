package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Class string
}

func main() {
	student := []Student{
		{
			Name:  "A",
			Age:   25,
			Class: "1반",
		},
		{
			Name:  "B",
			Age:   30,
			Class: "1반",
		},
		{
			Name:  "B",
			Age:   30,
			Class: "1반",
		},
		{
			Name:  "C",
			Age:   27,
			Class: "2반",
		},
		{
			Name:  "D",
			Age:   23,
			Class: "2반",
		},
	}
	fmt.Println(student[2])
}
