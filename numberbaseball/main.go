// 룰 = 컴퓨터가 무작위로 생각한 3가지의 숫자를 맞추는 것

package main

import (
	"fmt"
	"math/rand" // 랜덤 난수를 발생시키는 패키지
	"time"
)

type Result struct { // func CompareNumbers 함수에서 두가지를 내보내기 위해서 타입설정으로 한번에 보낼 수 있게 한다.
	strikes int
	balls   int
}

func main() {
	// 실행할때 마다 다른 값이 나오도록 시드(실행시간을 기준)값을 지정해준다.
	// UnixNano는 1970 1월 1일 부터 현재까지의 시간을 나노초로 나눠서 나타내는 기능이다.
	rand.Seed(time.Now().UnixNano())

	// 무작위 숫자 3개를 만든다.
	numbers := MakeNumbers()

	cnt := 0
	for { // 게임이 끝나지 않으면 무한루프로 돌아간다.
		cnt++ // 게임을 끝내지 못하면 횟수가 추가된다.
		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3S인가? - 게임이 끝나는가?
		if IsGameEnd(result) {
			// 만약 결과를 비교해서 두개가 맞으면 게임이 끝난다.
			break
		}
	}

	// 게임이 끝나면 몇번만에 맞췄는지 출력한다.
	fmt.Printf("%d번만에 맞췄습니다. \n", cnt)
}

// 무작위 숫자 3개를 만드는 함수
func MakeNumbers() [3]int {
	// 0~9 사이에 겹치지 않는 무작위 숫자 3개를 반환한다.
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			// 겹치지 않는 0~9사이의 무작위 숫자를 만들어낸다. [0~10)
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n { // 겹치게 되면 다시 뽑아야 하는 함수를 걸어놓는다.
					duplicated = true
					break // 겹치지 않게 되면 for문을 빠져나간다.
				}
			}
			if duplicated == false {
				rst[i] = n // 만약 값이 겹치지 않는다면 n에 값을 대입시킨다.
				break
			}
		}
	}

	// fmt.Println(rst) // ============================= 컴퓨터 랜덤 지정 난수 3개
	return rst
}

// 사용자 숫자 3개를 입력받는 함수
func InputNumbers() [3]int {
	// 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요.")
		// 키보드로부터 입력받는 기능
		var no int
		_, err := fmt.Scanf("%d\n", &no) // 키보드 버퍼에 엔터가 남아있어서 에러가 뜬다.
		// 해결을 위해 엔터값을 없애주려면 \n을 추가해주면 된다.
		if err != nil { // 에러가 나게 된다면(에러가 아무것도 아닌게 아니라면)
			fmt.Println("잘목 입력하셨습니다.")
			continue // 에러가 뜨면 컨티뉴 기능을 써서 for문에 다시 돌아가게 한다.
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10 // no 가 123일 때, n = no%10 하면 3이 남는다. 그리고 no = 123/10 이니까 12
			// 12는 다시 for로 돌아가고, n = 12%10이 되고, 12 = no/10 이 되고, 2니까 돌아간다.
			// 거기서 다시 2를 넣고 다시 for문을 실행하면 1의 몫이 출력된다.

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n { // 겹치게 되면 다시 뽑아야 하는 함수를 걸어놓는다.
					duplicated = true
					break // 겹치지 않게 되면 for문을 빠져나간다.
				}
			}
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break // 이 브레이크는 for no > 0 문에 브레이크가 걸린다.
			}

			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n // 4개를 입력하면 여기서 배열의 범위가 벗어났다는 에러가 나온다
			// 이 배열은 rst[3]int 인데 그거보다 크게 넣었기 때문에 에러가 발생한다.
			idx++
		}
		if success == true && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		} //

		if !success {
			continue // 여기서 컨티뉴를 하면 아래의 브레이크를 건너뛰고 다시 for문을 돌린다.
		}
		break // 이건 무한루피 for {} 를 빠져나오는 브레이크다.
	}
	rst[0], rst[2] = rst[2], rst[0] // 위의 for문에서 뒤섞인 숫자를 재정렬 해준다.
	// fmt.Println(rst) // =================== 사용자가 입력해서 나온 출력값을 지워준다.
	return rst
}

// 무작위 숫자 3개와 사용자 입력 숫자 3개를 비교하는 함수
func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// 두 개의 숫자 3개씩을 비교해서 두개가 같은지 결과를 반환한다.
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ { // 입력한 숫자를 앞자리 부터 같은 숫자가 있는지, 같은 위치의 인덱스에 있는지 확인한다.
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	return Result{strikes, balls}
}

// 결과를 출력하는 함수
func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}

// 결과가 참인지 확인하는 함수
func IsGameEnd(result Result) bool {
	// 비교 결과가 3S 인지 확인한다.
	return result.strikes == 3
}
