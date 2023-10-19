package channel

import (
	"log"
	"sync"
	"time"
)

type Messenger[T any] struct {
	Name         string
	Sender       chan T
	SenderClosed bool
}

func (m *Messenger[T]) loop() {
	for v := range m.Sender {
		message := v
		log.Println("Message: ", message)
	}
	log.Println("already closed")
	m.SenderClosed = true
}

func (m *Messenger[T]) close() {
	close(m.Sender)
}

func (m *Messenger[T]) Exec(v T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			m.SenderClosed = true
		}
	}()
	m.Sender <- v
}

func Sample1() {
	msg := Messenger[string]{
		Name:   "main messenger",
		Sender: make(chan string),
	}
	go msg.loop()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		msg.Sender <- "is opened"
		msg.close()
		time.Sleep(1 * time.Second)
		if msg.SenderClosed {
			log.Println("closed")
		} else {
			log.Println("not yet")
			msg.Sender <- "help "
		}
		wg.Done()
	}()
	wg.Wait()
	log.Println("all done")
}

func Sample2() {
	msg := Messenger[string]{
		Name:   "main messenger",
		Sender: make(chan string),
	}
	go msg.loop()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		msg.Sender <- "is opened"
		msg.close()
		time.Sleep(1 * time.Second)
		msg.Exec("test ok")
		wg.Done()
	}()
	wg.Wait()
	log.Println("all done")
}
