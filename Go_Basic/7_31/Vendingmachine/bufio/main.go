package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("소지금액을 적어주세요 : ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	coin, _ := strconv.Atoi(str)

	// if (cost < 0) || (cost > 3000) {
	// 	fmt.Println("0 ~ 3000 사이의 값을 입력해주세요")
	// 	continue
	// } else {
	// 	break
	// }

	var menu int

	fmt.Println("-----------------")
	fmt.Println("-----자판기------")
	fmt.Println("-----------------")
	fmt.Println("--(1) 콜라 _800--")
	fmt.Println("--(2) 환타 _800--")
	fmt.Println("--(3) 제티 _700--")
	fmt.Println("--(4) 우유 _600--")
	fmt.Println("--(5) 생수 _500--")
	fmt.Println("-----------------")
	fmt.Println("---      (500)---")
	fmt.Println("---           ---")

	for {
		fmt.Println("숫자로 선택해주세요.")
		reader := bufio.NewReader(os.Stdin)
		drink, _ := reader.ReadString('\n')
		drink = strings.TrimSpace(drink)
		result, _ := strconv.Atoi(drink)
		menu = result
		break
	}

	switch menu {
	case 1:
		coin = coin - 1000
		fmt.Println("게토레이가 나왔습니다.")
	case 2:
		coin = coin - 700
		fmt.Println("스프라이트가 나왔습니다.")
	case 3:
		coin = coin - 700
		fmt.Println("코카콜라가 나왔습니다.")
	case 4:
		coin = coin - 800
		fmt.Println("환타가 나왔습니다.")
	case 5:
		coin = coin - 600
		fmt.Println("솔의눈이 나왔습니다.")
	default:
		coin = coin + 500
		fmt.Println("바닥의 돈을 주웠습니다.")
	}

	fmt.Printf("남은 돈은 %v 입니다.", coin)

	if coin == 0 {
		fmt.Println("소지금액이 없습니다.")
	}
}
