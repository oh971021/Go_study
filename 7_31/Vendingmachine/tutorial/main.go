// 코드훈련

// 코딩을 어떻게 해야할 까?
// 스스로 만든다 ? X
// 정확한 답은 구글링 해서
// 코딩을 똑같이 쳐본다.
// 복사 X 스스로 손코딩
// 기능을 분해하는 능력이 있다면
// 코딩은 이해가 간다.
// 하지만 손에 익지 않는다.
// 구글링을 해보고, 쳐보고, 지우고, 쳐보기를 반복한다.
// 코드가 익숙해진다면 변수 및 값을 변경하며 작성해본다.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 소지금액 입력 기능
	fmt.Println("소지금액을 입력하세요.")

	// 버퍼라는 컴퓨터에 가상의 저장메모리를 만들어 놓음
	// 가장 대표적으로 복사하고, 클립보드에 복사한 데이터가 들어감.
	// 버퍼 안에 기본적인 입출력을 담당하는 패키지가 bufio이고,
	// bufio는 버퍼에 데이터를 입력하고 출력하는 기능
	// bufio.NewReader : 버퍼 안에 데이터를 읽는 기능을 하는 상자를 만드는 것.
	// 이 상자 안에는 운영체제에서 표준의 입출력 기능을 가지도록 입력시킨다.
	reader := bufio.NewReader(os.Stdin)

	// reader 로봇이 사용자가 키보드로 입력을 했을 때, 입력받은 데이터는 문자열인데,
	// 이 문자열의 개행이 되는 문자열을 읽어오겠다는 뜻.
	// _ : 원래는 error 타입으로 반환되는 부분인데, error를 무시하겠다는 의미
	str, _ := reader.ReadString('\n')

	// strings : 문자열을 제어하는 패키지
	// TrimSpace : 문자열에 존재하는 불필요한 데이터, 공백들을 모두 제거해주는 함수
	str = strings.TrimSpace(str)

	// strconv : 문자열을 형변환 시켜주는 패키지
	// Atoi : 문자열 -> 정수 변환 기능
	cost, _ := strconv.Atoi(str)

	// if (cost < 500) || (cost > 3000) {
	// 	fmt.Println("500 ~ 3000 사이의 소지금액만 적어주세요.")
	// 	continue
	// } else {
	// 	money = cost
	// 	break
	// }

	// 메뉴 출력 기능
	fmt.Println("-----------------")
	fmt.Println("-----자판기------")
	fmt.Println("-----------------")
	fmt.Println("--(1) 콜라 _800--")
	fmt.Println("--(2) 환타 _800--")
	fmt.Println("--(3) 제티 _700--")
	fmt.Println("--(4) 우유 _600--")
	fmt.Println("--(5) 생수 _500--")
	fmt.Println("-----------------")

	// 음료수 입력 기능
	res := 0
	for {
		fmt.Println("숫자로 음료를 선택해주세요.")
		reader := bufio.NewReader(os.Stdin)
		drink, _ := reader.ReadString('\n')
		drink = strings.TrimSpace(drink)
		result, _ := strconv.Atoi(drink)
		res = result
		break
	}

	// 음료수 출력 기능
	switch res {
	case 1:
		cost = cost - 800
		fmt.Println("콜라가 나왔습니다")
	case 2:
		cost = cost - 800
		fmt.Println("환타가 나왔습니다")
	case 3:
		cost = cost - 700
		fmt.Println("제티가 나왔습니다")
	case 4:
		cost = cost - 600
		fmt.Println("우유가 나왔습니다")
	case 5:
		cost = cost - 500
		fmt.Println("생수가 나왔습니다")
	default:
		cost = cost - 1000
		fmt.Println("기계가 돈을 먹었습니다")
	}

	fmt.Printf("남은 돈은 %v", cost)

	if cost == 0 {
		fmt.Println("소지 금액이 없습니다.")
	}
	// switch 변수(expression) {} Golang은 expression이 없어도 사용 가능함. 다른언어는 필수
}
