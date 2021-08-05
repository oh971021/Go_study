package koreaPS

import "fmt"

// 한국우체국라는 패키지 안에
// 우편물발송이라는 객체가 있고, 그 객체는 Send라는 메소드를 포함하고 있다.
type PostSender struct {
}

func (p *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %s를 보냅니다.\n", parcel)
}
