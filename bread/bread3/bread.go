package main

import "fmt"

type Metarial interface {
	GetOneChopstick() *ChopstickOfKimchi
}

type BoB struct {
	val string
}

type Kimchi struct {
}

type ChopstickOfKimchi struct {
}

func (c *ChopstickOfKimchi) String() string {
	return " + Kimchi"
}

func (k *Kimchi) GetOneChopstick() *ChopstickOfKimchi {
	return &ChopstickOfKimchi{}
}

// 밥은 김치를 통해서 메소드를 호출한다. 여기서 관계를 맺는다.
func (b *BoB) PutKimchi(kimchi Metarial) {
	chopstick := kimchi.GetOneChopstick()
	b.val += chopstick.String()
}

func (b *BoB) String() string {
	return "BoB" + b.val
}

func main() {
	bob := &BoB{}
	kimchi := &Kimchi{}
	bob.PutKimchi(kimchi)

	fmt.Println(bob)
}
