/*
  = 에러 랩핑 =
  - fmt.Errorf()의 %w 포맷터로 에러 랩핑 가능
  - errors.ls() 함수와 As() 함수로 랩핑 된 에러 체크 및 변환

  - 어느 부분에서 에러가 났는지 확인 하기 위해서 그 위치를 래핑해서 상위 에러에게 보낸다.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromstring(str string) (int, error) {
	// NewScanner 는 ioread type을 인자로 받는다.
	// 스캐너는 어떤 데이터를 가지고 올때, 일정한 규칙으로 (한줄씩, 한단어씩) 가져올 때 사용.
	// strings.Newreader 는 인자를 Reader 객체로 만들어서 반환해준다.
	scanner := bufio.NewScanner(strings.NewReader(str))
	// 한 단어 씩 끊어서 읽는다.
	scanner.Split(bufio.ScanWords)

	pos := 0
	// 첫 단어를 읽어서 a 에 받고, n 에 읽은 긁자수, error 는 err
	a, n, err := readNextInt(scanner)
	if err != nil {
		// %w 가 에러를 감싸는 것. ( 래핑함 )
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d, err:%w", pos, err)
	}

	pos += n + 1
	// 첫 단어를 읽어서 b에 받고, n 에 읽은 긁자수, error 는 err
	b, n, err := readNextInt(scanner)
	if err != nil {
		// %w 가 에러를 감싸는 것. ( 래핑함 )
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d, err:%w", pos, err)
	}
	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환하는 함수
// 반환값은 반환된 숫자, 읽은 긁자 수, 에러를 반환한다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	// 한 단어를 읽는 것에 실패했다.
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	// 성공했다면 문자열로 바꿔서 워드라는 변수에 넣어준다.
	word := scanner.Text()

	// 이 문자열을 숫자로 바꿔준다.
	number, err := strconv.Atoi(word) // "24" -> 24 로 바꾸는 함수
	// 변환에 실패했다면
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s, err:%w", word, err)
	}
	// 변환에 성공햇다면 그 숫자와, 길이, 에러는 nil로 반환한다.
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromstring(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		// errors.As 는 래핑된 에러를 변환해주는 함수
		// err 안에 numError라는 객체가 있으면 열어서 어떤 에런지 확인 할 수 있는 As 함수
		// 하위 에러를 꺼내는 함수
		if errors.As(err, &numError) {
			fmt.Println("NumberError", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
