/*
  ## 에러핸들링 ? ##
   - 에러를 어떻게 반환하는지 정해주는 방법
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// 파일을 받아서 파일과 에러를 반환한다.
func ReadFile(filename string) (string, error) {
	// Open 이 친구는 file을 열 때, 문제가 생기면 호출자에게 Error 처리를 위임한다.
	// 즉 반환해준다는 뜻.
	file, err := os.Open(filename)
	if err != nil {
		// 에러가 있으면 빈칸과 에러를 반환한다.
		return "", err
	}
	// 한 번 열린 파일은 함수 내에서 꼭 닫아줘야한다.
	defer file.Close()

	// 위에서 파일에 문제가 없다는 것을 알았음.
	// 다시 읽어서 "rd" 라는 buf를 생성해서 buf에 넣는다.
	rd := bufio.NewReader(file)
	// 개행까지 읽는다. (한 줄을 읽는다.)
	// io.EOF 라는 에러는 파일을 다 읽었다 라는 뜻인데, 중요하지 않아서 무시한다.
	// 그 buf에서 한줄을 읽어서 반환하는데, 요 친구는 data, error 를 반환해서 두개를 반환한다.
	line, _ := rd.ReadString('\n')
	// 에러가 없기 때문에 위에서 읽엇던 한줄, 즉 line과 없는에러 nil로 반환한다.
	return line, nil
}

func WriteFile(filename string, line string) error {
	// file이 원래 없었기 때문에 하나 만들어준다.
	// careate는 filename에 해당하는 file을 만들어서 확인하고 반환한다.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// F 라고 하는 건 첫번째 인자로 writer를 정하고(어디에 쓸지), 무엇을 쓸지를 정한다.
	_, err = fmt.Fprintln(file, line)
	return err
}

// 얘는 파일을 처음에 갖고있지않지만,
// WriteFile 에 들어갔다오면서 파일이 생김.
const filename string = "data.txt"

func main() {
	// 파일을 읽어서 한줄을 뽑아내고, 에러가 있으면 에러를 뽑아낸다.
	line, err := ReadFile(filename)

	// 왜 이렇게 될까요??
	// 위에서 첫 Read 에서는 file이 없어서 err가 반환됨.
	if err != nil {
		// 그래서 file을 만든다.
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		// 파일 쓰는 것에 성공햇다면 다시 읽는다.
		line, err = ReadFile(filename)
		// 다시 읽었는데 실패했다면 문제가 있는거...
		if err != nil {
			fmt.Println("파일 읽기에 실패햇습니다.", err)
			return
		}
	}
	fmt.Println("파일 내용:", line)
}
