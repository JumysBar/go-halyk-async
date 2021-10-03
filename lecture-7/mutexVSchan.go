package main

import (
	"fmt"
	"sync"
	"time"
)

func increaseWithMutex(x *int64, mx *sync.Mutex) {
	mx.Lock()
	defer mx.Unlock()

	*x++
}

func withMutex() {
	value := int64(0)
	mx := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		go increaseWithMutex(&value, mx)
	}

	time.Sleep(time.Millisecond * 100)

	mx.Lock()
	defer mx.Unlock()
	fmt.Printf("value: %d\n", value)
}

func increaseWithChannelReadWithRace(x *int64, ch chan struct{}) {

LOOP:
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				break LOOP
			}
			*x++
		}
	}
}

func withChannelsReadWithRace() {
	value := int64(0)

	ch := make(chan struct{})
	go increaseWithChannelReadWithRace(&value, ch)

	for i := 0; i < 1000; i++ {
		ch <- struct{}{}
	}
	close(ch)

	time.Sleep(time.Millisecond * 100)
	fmt.Printf("value: %d\n", value)
}

func main() {
	withChannelsReadWithRace()

	withMutex()

	//withChannels()
}

func withChannels() {
	value := int64(0)

	ch := make(chan struct{})
	getCh := make(chan int64)
	go increaseWithChannel(&value, ch, getCh)

	for i := 0; i < 1000; i++ {
		ch <- struct{}{}
	}
	close(ch)

	time.Sleep(time.Millisecond * 100)
	fmt.Printf("value: %d\n", <-getCh)
}

func increaseWithChannel(x *int64, ch chan struct{}, getCh chan int64) {

LOOP:
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				getCh <- *x
				break LOOP
			}
			*x++
		}
	}
}
