/*
자판기 프로그램 만들기

=== 로직 ===
1. 프로그램 실행 시 금액 입력
2. 자판기 메뉴 출력
3. 메뉴를 번호로 입력 -> 소지금액에 맞게 감소
4. 메뉴가 나왔다. 그리고 자판기 메뉴로 돌아감.
5. 소지금액이 다 소진되면 집에가기Vending machine
*/

package main

import "fmt"

func main() {
	var cost int

	for {
		fmt.Println("소지금액을 입력해주세요")
		// _, error 작성법 : 들어가는변수, err변수
		// 현재 들어가는 변수는 정의 되어있어서 받는 변수를 미식별자로 만든다.
		// error 는 다른 숫자를 넣었을 때 err를 표시해주기 위해 만들어 준 것.
		_, error := fmt.Scanf("%d\n", &cost)

		// 타입이 맞지 않는 어떤 값이 들어오면 에러가 뜨도록 만든다.
		if error != nil {
			fmt.Println("잘못 입력하셨습니다")
			// 만약 입력한 값이 0보다 작거나, 10000보다 크면 초기화한다.
		} else if (cost < 0) || (cost > 10000) {
			continue
		} else {
			break
		}
	}

	var menu int

	for {
		fmt.Println("----------------------------")
		fmt.Println("----------------------------")
		fmt.Println("----------------------------")
		fmt.Println("---- 동대구역 빅자판기 -----")
		fmt.Println("----------------------------")
		fmt.Println("----------------------------")
		fmt.Println("--- 1. 게토레이 (1000)------")
		fmt.Println("--- 2. 스프라이트 (700)-----")
		fmt.Println("--- 3. 코카콜라 (700)-------")
		fmt.Println("--- 4. 환타 (700)-----------")
		fmt.Println("--- 5. 솔의눈 (600)---------")
		fmt.Println("--- 6. 토레타 (800)---------")
		fmt.Println("--- 7. 쿠우 (1200)----------")
		fmt.Println("--- 8. 파워에이드 (1500)----")
		fmt.Println("--- 9. 물 (500)-------------")
		fmt.Println("--- 10. 조지아 (500)--------")
		fmt.Println("--- 11. 랜덤 (700)----------")
		fmt.Println("----------------------------")
		fmt.Println("-----          (500)   -----")
		fmt.Println()
		fmt.Println()
		_, err := fmt.Scanf("메뉴를 입력하세요 : %d\n", &menu)

		if err != nil {
			fmt.Println("잘못 입력하셨습니다")
		} else if (menu < 0) || (menu > 11) {
			continue
		} else {
			break
		}
	}
}
