/*
DB하다가 Go 로 만들어보고 싶어서 만든 예제.

5명의 학생을 만들어서 학생들의 나이를 비교한 후, 나이가 가장 많은 친구와 작은 친구를 뽑아본다.
그 때, 그 학생의 정보를 같이 출력한다.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 학생 타입을 하나 만들어 줌
type student struct {
	name  string
	age   int
	class string
}

func main() {
	// array type을 작성해주고,
	// 학생 Data field를 입력해준다.
	students := []student{
		{
			name:  "A",
			age:   25,
			class: "1반",
		},
		{
			name:  "B",
			age:   30,
			class: "1반",
		},
		{
			name:  "B",
			age:   30,
			class: "1반",
		},
		{
			name:  "C",
			age:   27,
			class: "2반",
		},
		{
			name:  "D",
			age:   23,
			class: "2반",
		},
	}

	// 가장 나이가 어린 친구를 찾기 위해서 변수를 하나 생성해준다.
	// 기준은 A친구로 잡아주는 듯 하다.
	smallest := student{age: students[0].age}
	biggest := student{age: students[0].age}

	// 순환문을 통해 []student 만큼 돌아준다.
	// 돌면서 각각의 변수를 넣어주고, age field 를 출력해준다.
	// 그러면 각 학생의 나이 값을 출력한다.
	for i := 0; i < len(students); i++ {
		name := students[i].name
		age := students[i].age
		class := students[i].class
		fmt.Println("현재 나이:", age)

		// smallest 변수에 나이가 가장 어린 친구가 들어올 때 까지
		// 순환문을 돌면서 기준의 친구는 계속 바뀐다.
		if age <= smallest.age {
			smallest.name = name
			smallest.age = age
			smallest.class = class
		}

		// samllest 와 반대로 가장 나이가 많은 친구가 들어올 때 까지
		// 순환문을 돌리면서 변수 대입을 통해 기준을 바꿔준다.
		if age >= biggest.age {
			biggest.name = name
			biggest.age = age
			biggest.class = class
		}
	}

	// 문단을 나눌 때, 간단하게 중복값을 뽑아주는 strings.Repeat 패키지
	// 파이썬은 * 하나면 되는데.....
	line := strings.Repeat("=", 30)
	fmt.Println(line)

	// 변수를 하나 생성해서 int type의 age 필드를 넣어준다.
	smallestAge1 := smallest.age
	// 아래 Print에 문자열로 사용하기 위해서 strconv.Itoa(int -> string) 으로 타입변환시켜줌.
	smallestAge2 := strconv.Itoa(smallestAge1)
	// 그래서 결국 모든 필드의 type은 string이고, 다 합쳐서 출력시킨다.
	fmt.Println(smallest.class + "의" + smallest.name + "학생은" + smallestAge2 + "의 나이로 나이가 가장 작습니다.")

	fmt.Println(line)

	biggestAge1 := biggest.age
	biggestAge2 := strconv.Itoa(biggestAge1)
	fmt.Println(biggest.class + "의" + biggest.name + "학생은" + biggestAge2 + "의 나이로 나이가 가장 많습니다.")

	fmt.Println(line)
}
