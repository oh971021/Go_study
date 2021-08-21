// 우리가 원하는 에러 객체를 생성해준다.

package main

import "fmt"

// 우리가 원하는 정보를 담아서 error 를 찾아내기 위해서 사용함.
// 이 type은 메서드로 인해 error로 사용할 수 있다.
type PasswordError struct {
	Lne        int
	RequireLen int
}

// 에러 메세지를 반환하는 메서드 ( 이 친구만 있으면 뭐든 error로 사용할 수 있다 )
func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	// 패스워드 길이가 8보다 작으면 에러 메소드를 반환한다.
	if len(password) < 8 {

		/*
			// 필요한 길이를 알려주는 에러 메세지를 생성.
			// 얘는 특정한 정보를 못 뽑아주기 때문에 아래의 방법이 많이 쓰인다.
			return fmt.Errorf("암호 길이가 짧습니다. 필요한 길이 : %d", 8)
		*/

		// 세 방법 다 에러 객체를 생성하는 것이라 다 똑같다.
		// 그래도 인터페이스는 우리가 원하는 사용자 타입의 에러를 반환할 수 있다는 점이 있다.
		// return fmt.Errorf("...")
		// return errors.New("...")
		return PasswordError{len(password), 8}
	}
	// ... ( 계정 등록 )
	return nil
}

func main() {
	err := RegisterAccount("my ID", "my Pw")
	if err != nil {
		// 패스워드 에러로 타입변환한다.
		// 타입 변환의 성공 유무를 ok에 담는다.
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v, Len:%d, RequireLen:%d\n",
				errInfo, errInfo.Lne, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원가입 됐습니다.")
	}
}
