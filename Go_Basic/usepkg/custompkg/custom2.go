package custompkg // 하나의 파일에 들어있는 패키지들은 같은 이름을 가져야한다.

func print2() {
	PrintCustom() // 같은 패키지 내에서는 import로 호출을 하지않아도 바로 사용이 가능하다.
}
