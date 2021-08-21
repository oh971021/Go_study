/*
go언어에서 제너릭 처리를 위해서는 아래처럼  인터페이스를 정의하고 각 타입별로 구현을 한다.
슬라이스에 각 타입으로 컨버전하여 값을 넣고 슬라이스 순번대로 호출하면
자신의 타입에 맞는 메소들를 호출하여 처리한다.
*/

package main

import "fmt"

type integer32 int32

type integer64 int64

type Calculator interface {
	Calculate()
}

func (i integer32) Calculate() {
	fmt.Println("Integer32", i)
}

func (i integer64) Calculate() {
	fmt.Println("Integer64", i)
}

func main() {
	container := []Calculator{integer32(32), integer64(64)}
	fmt.Println("container", container)
	container[0].Calculate()
	container[1].Calculate()
}
