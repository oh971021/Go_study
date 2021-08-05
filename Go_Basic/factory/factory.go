package main

import (
	"fmt"
	"strconv" // 형 변환 패키지
	"time"
)

type Car struct { // Car 타입 만들기
	val string
}

type Plane struct { //	Plane 타입 만들기
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outChan chan Car, outPlaneChan chan Plane) { // 쓰레드 1, 하나의 함수에 채널 4개를 만든다.
	for {
		select { // 두개(car, plane)가 한꺼번에 대기 했다가 값이 둘다 들어오면 같이 작동하도록 select 으로 잡아준다.
		case car := <-carChan: // car 에 특정 값이 들어올 때 까지 대기한다.
			car.val += "번 타이어 설치"
			outChan <- car // Carchan 에 어떤 값이 들어올 때까지는 아래로 내려가지말고 대기한다. 여기서 값을 받아주면 다시 올라가서 car := <- carchan 부터 실행된다.
		case plane := <-planeChan: // plane 에 특정 값이 들어올 때 까지 대기한다.
			plane.val += "번 타이어 설치"
			outPlaneChan <- plane // plane 에 특정 값이 입력 되면 올라가서 다시 실행한다.
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outChan chan Car, outPlaneChan chan Plane) { // 쓰레드 2, 하나의 함수에 채널 4개를 만든다.
	for {
		select {
		case car := <-carChan:
			car.val += ", 엔진 설치"
			outChan <- car
		case plane := <-planeChan:
			plane.val += ", 엔진 설치"
			outPlaneChan <- plane
		}
	}
}

func StartWrok(chan1 chan Car) { // 차체를 만드는 함수 쓰레드 3
	i := 0 // int 형으로 선언하고, 무한루프에서 돌아가도록 만든다. 자동차나 비행기의 숫자를 나타내는 변수
	for {
		time.Sleep(1 * time.Second)                // 1초에 하나씩 만들어라.
		chan1 <- Car{val: "자동차" + strconv.Itoa(i)} // int를 string으로 변환시킨다.
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) { // 비행기를 만드는 함수 쓰레드 4
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "비행기" + strconv.Itoa(i)}
		i++
	}
}

func main() {
	carchan1 := make(chan Car)
	carchan2 := make(chan Car)
	carchan3 := make(chan Car) // 채널은 3개를 사용한다.

	planechan1 := make(chan Plane)
	planechan2 := make(chan Plane)
	planechan3 := make(chan Plane) // 채널은 3개를 사용한다.

	// 총 6개의 채널을 운영한다.

	go StartWrok(carchan1)                                    // go 쓰레드 1
	go StartPlaneWork(planechan1)                             // go 쓰레드 2
	go MakeTire(carchan1, planechan1, carchan2, planechan2)   // go 쓰레드 3
	go MakeEngine(carchan2, planechan2, carchan3, planechan3) // go 쓰레드 4

	for {
		select { // 두개(car, plane)가 한꺼번에 대기 했다가 같이 작동하도록 select 으로 잡아준다.
		case result := <-carchan3:
			fmt.Println(result.val)
		case result := <-planechan3:
			fmt.Println(result.val)
		} // 메인 쓰레드 5
	}
}
