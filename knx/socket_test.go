package knx

import (
	"container/list"
	"errors"
	"sync"
)

type dummySocket struct {
	cond    *sync.Cond
	out     *list.List
	in      *list.List
	inbound chan interface{}
}

func (sock *dummySocket) serveOne() bool {
	sock.cond.L.Lock()

	for sock.in != nil && sock.in.Len() < 1 {
		sock.cond.Wait()
	}

	if sock.in == nil {
		sock.cond.L.Unlock()
		return false
	}

	val := sock.in.Remove(sock.in.Front())

	sock.cond.Broadcast()
	sock.cond.L.Unlock()

	sock.inbound <- val

	return true
}

func (sock *dummySocket) serveAll() {
	for sock.serveOne() {}
	close(sock.inbound)
}

func (sock *dummySocket) closeIn() {
	sock.cond.L.Lock()
	defer sock.cond.L.Unlock()

	sock.in = nil

	sock.cond.Broadcast()
}

func (sock *dummySocket) closeOut() {
	sock.cond.L.Lock()
	defer sock.cond.L.Unlock()

	sock.out = nil

	sock.cond.Broadcast()
}

func (sock *dummySocket) Send(payload OutgoingPayload) error {
	return sock.sendAny(payload)
}

func (sock *dummySocket) sendAny(payload interface{}) error {
	sock.cond.L.Lock()
	defer sock.cond.L.Unlock()

	if sock.out == nil {
		return errors.New("Outbound is closed")
	}

	sock.out.PushBack(payload)
	sock.cond.Broadcast()

	return nil
}

func (sock *dummySocket) Close() error {
	sock.cond.L.Lock()
	defer sock.cond.L.Unlock()

	sock.in = nil
	sock.out = nil

	sock.cond.Broadcast()

	return nil
}

func (sock *dummySocket) Inbound() <-chan interface{} {
	return sock.inbound
}

func makeDummySockets() (*dummySocket, *dummySocket) {
	cond := sync.NewCond(&sync.Mutex{})
	forGateway := list.New()
	forClient := list.New()

	client := &dummySocket{cond, forGateway, forClient, make(chan interface{})}
	go client.serveAll()

	gateway := &dummySocket{cond, forClient, forGateway, make(chan interface{})}
	go gateway.serveAll()

	return client, gateway
}
