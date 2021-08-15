package main

func main() {
	var a interface{} = 3.14

	j := a

	// a 가 float64 인 경우 return 한다.
	// 결국 a.T == float64 인 경우 i.T == float64 가 된다는 뜻.
	// 결론 : a의 타입을 i에 그대로 대입시킨다.
	i := a.(float64)

	// 만약 a.T != i.T 인 경우 panic이 온다!
	// i := a.(float32)

	println(j, "\n", i)
}
