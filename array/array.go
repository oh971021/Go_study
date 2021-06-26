package main

import "fmt"

func main() {
	var a = [3]int{ //1차원
		1, 2, 3,
	}

	var b = [2][3]int{ //2차원
		{1, 2, 3},
		{4, 5, 6},
	}

	var c = [2][2][3]int{ //3차원
		{{1, 2, 3},
			{4, 5, 6}},
		{{1, 2, 3},
			{4, 5, 6}},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
