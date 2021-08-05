package fedex

import "fmt"

type FedexSender struct {
}

// FedexSender 라는 타입이 있는데, Send라는 메소드를 가지고 있따.
func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex Sends %s parcel\n", parcel)
}
