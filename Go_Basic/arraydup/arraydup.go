package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}
	temp := [5]int{}

	fmt.Println("array : ", arr)
	fmt.Println("clone : ", clone)
	fmt.Println("temp : ", temp)
	fmt.Println("array2 : ", arr2)

	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
	}
	fmt.Println("dup array clone : ", clone)

	for i := 0; i < 5; i++ { // arr를 역순으로 만들기위해 임시배열을 만들고
		temp[i] = arr[len(arr)-1-i] // 템프의 값을 arr의 역순으로 만든다.
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i] // 역순으로 변환 된 템프의 값을 arr로 복사시킨다.
	}

	fmt.Println("changed temp : ", temp)
	fmt.Println("changed arr : ", arr)

	// 임시배열변수를 만들지 않고, 바로 역순으로 바꾸는 방법 ----
	// for문을 두번 사용해서 돌린다. (반반 나눠서)

	for i := 0; i < len(arr2)/2; i++ { // arr2 배열길이의 반만 돌릴때 사용한다.
		arr2[i], arr2[len(arr2)-1-i] = arr2[len(arr2)-1-i], arr2[i]
		// a, b = 2, 3 : a = 2, b = 3 공식
	}
	fmt.Println("changed arr2 : ", arr2)
}
