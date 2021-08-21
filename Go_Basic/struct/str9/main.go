/*
	8월 12일 -- 성적 저장 프로그램 예제 (구름에듀)
*/

package main

import "fmt"

// 학생 정보를 담은 struct 선언
type student struct {
	name   string
	gender string

	// key = string / value = int
	data map[string]int
}

// student 의 data를 초기화하는 함수 생성
func newStudent() *student {
	// 변수로 선언해준다.
	d := student{}

	// map 을 초기화 시킨다.
	d.data = map[string]int{}
	return &d
}

func main() {
	// 정수에는 학생번호, 과목번호, 점수가 있다.
	var studentNumber, subjectNumber, score int

	// 문자열에는 이름, 성별, 과목이 있다.
	var name, gender, subject string

	// 학생 번호와 과목 번호를 입력한다.
	fmt.Scanln(&studentNumber, &subjectNumber)

	// studne type의 배열을 생성한다. 길이는 학생번호만큼.
	s := make([]student, studentNumber)

	// 순환문을 돌린다.
	for i := 0; i < studentNumber; i++ {
		// 돌면서 data 초기화함수를 부르고,
		object := newStudent()

		// 요거는 학생 정보랑 성별을 각각 입력해주는 친구
		fmt.Scanln(&name, &gender)

		// 입력받은 이름과 성별을 초기화된 변수에 대입해준다.
		object.name = name
		object.gender = gender

		// 두번째 순환문을 돌리는데,
		for j := 0; j < subjectNumber; j++ {
			// 이땐 과목이름과 점수를 적는다.
			fmt.Scanln(&subject, &score)
			// 그리고 map type에 key 값으로 과목이름, value 값으로 점수를 대입시킨다.
			object.data[subject] = score
		}
		// 과목 번호만큼 돌면서 과목 번호만큼 데이터들을 넣어준다.
		s[i] = *object
	}

	// 출력을위한 순환문
	for i := 0; i < studentNumber; i++ {
		fmt.Println("----------")
		// 과목번호에 맞는 학생번호를 맞춰서 출력해준다.
		// 이름과 성별 출력
		fmt.Println(s[i].name, s[i].gender)

		// range 문을 통해 key, value 를 쌍으로 출력한다.
		for index, val := range s[i].data {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
