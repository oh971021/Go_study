package main

import "context"

type Publisher struct {
	ctx context.Context

	// 채널타입을 받는 채널을 만들 수 있다.
	// chan <- string 은 일방향 채널 : writer only channel
	// <- chan string 도 일방향 채널 : read only channel
	subscribeCh chan chan<- string
	publishCh   chan string

	// 구독자를 받는다.
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
