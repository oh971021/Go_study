package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("이름을 입력해주세요 : ")
	name, _ := reader.ReadString('\n')
	rs := []rune(name)
	fmt.Println(string(rs[:len(rs)-1]))
}
