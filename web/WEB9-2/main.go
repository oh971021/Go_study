package main

import (
	"app/Go-study/web/WEB9-2/cipher"
	"app/Go-study/web/WEB9-2/lzw"
	"fmt"
)

type Send interface {
	SendData(string)
}

type Receive interface {
	ReceiveData(string)
}

var sentData string
var recvData string

// -------- 전송기능 -------- //
type SendComponent struct{}

func (self *SendComponent) SendData(data string) {
	//Send Data
	sentData = data
}

// -------------------------- //

// -------- 압축기능 --------- //
type ZipComponet struct {
	com Send
}

func (self *ZipComponet) SendData(data string) {
	// 미리 만들어놓은 lzw패키지의 write 함수를 통해 압축한다.
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.SendData(string(zipData))
}

// --------------------------- //

// 암호화 시킨다.
// 키와 컴퍼넌트를 가지는 구조체
type EncryptComponent struct {
	key string
	com Send
}

func (self *EncryptComponent) SendData(data string) {
	// cipher.go 의 encrypt 를 불러와서 키값을 넣고, 암호화 시킴.
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.SendData(string(encryptData))
}

// 암호를 풀어준다.
type DecryComponent struct {
	key string
	com Receive
}

func (self *DecryComponent) ReceiveData(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.ReceiveData(string(decryptData))
}

// -------- 압축해제 -------- //
type UnzipComponent struct {
	com Receive
}

func (self *UnzipComponent) ReceiveData(data string) {
	// 압축을 푸는 패키지
	// read 함수를 통해서 압축된 값을 해제해준다.
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	// 에러가 없으면 오퍼레이터가 호출되어 처리한다.
	self.com.ReceiveData(string(unzipData))
}

// ------------------------- //

// ----- 최종 응축 단계 ----- //
type ReadComponent struct{}

// 암호화가 끝나고 복호화된 데이터를 응축시켜 넘겨준다.
func (self *ReadComponent) ReceiveData(data string) {
	recvData = data
	// 요 데이터는 Hello World 를 나타냄
}

// ------------------------- //

func main() {

	// 키 값과 컴퍼넌트를 넣어서 데이터를 하나 만든다..
	sender := &ZipComponet{
		// 집 컴퍼넌트 타입의 컴퍼넌트.:
		// 그 안은 샌드 컴퍼넌트 타입의 컴퍼넌트.
		com: &SendComponent{},
	}

	// 오퍼레이터들에 값을 보내준다.
	sender.SendData("success")

	// 암호화 + 압축 데이터는 최종 sentData에 들어있다.
	fmt.Println(sentData)

	// 압축 풀린 값을 리시버에 넣는다.
	receiver := &UnzipComponent{
		// 키 값과 컴퍼넌트를 넣어서 암호 풀어줌
		com: &ReadComponent{},
	}

	receiver.ReceiveData(sentData)
	fmt.Println(recvData)
}

// 암호화와 압축
