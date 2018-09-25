package main

import (
	"fmt"
	"time"
)

type Broker struct {
	stopCh        chan struct{}
	publishCh     chan interface{}
	subscribeCh   chan chan interface{}
	unsubscribeCh chan chan interface{}
}

func NewBroker() *Broker {
	return &Broker{
		stopCh:        make(chan struct{}),
		publishCh:     make(chan interface{}, 1),
		subscribeCh:   make(chan chan interface{}, 1),
		unsubscribeCh: make(chan chan interface{}, 1),
	}
}

func (b *Broker) Start() {
	subs := map[chan interface{}]struct{}{}
	for {
		select {
		case <-b.stopCh:
			return
		case msgCh := <-b.subscribeCh:
			subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubscribeCh:
			delete(subs, msgCh)
		case msg := <-b.publishCh:
			for msgCh := range subs {
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker) Stop() {
	close(b.stopCh)
}

func (b *Broker) Subscribe() chan interface{} {
	msgCh := make(chan interface{}, 5)
	b.subscribeCh <- msgCh
	return msgCh
}

func (b *Broker) Unsubscribe(msgCh chan interface{}) {
	b.unsubscribeCh <- msgCh
}

func (b *Broker) Publish(msg interface{}) {
	b.publishCh <- msg
}

func main() {
	b := NewBroker()
	go b.Start()

	// create and subscribe 3 users:
	userReceiveNotifications := func(id int) {
		msgCh := b.Subscribe()
		for {
			fmt.Printf("User %d got notification: %v\n", (id + 1), <-msgCh)
		}
	}
	for i := 0; i < 3; i++ {
		go userReceiveNotifications(i)
	}

	// publish the notifications:
	go func() {
		for notificationID := 0; ; notificationID++ {
			b.Publish(fmt.Sprintf("number: %d", notificationID))
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
}
