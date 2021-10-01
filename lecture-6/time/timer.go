package main

import (
	"fmt"
	"time"
)

func veryLongFunction(d time.Duration, output chan<- struct{}) {
	time.Sleep(d)
	output <- struct{}{}
	close(output)
}

func main() {
	timer := time.NewTimer(3 * time.Second)

	ch := make(chan struct{})

	go veryLongFunction(time.Second, ch)
	// go veryLongFunction(4*time.Second, ch)

FORLOOP:
	for {
		select {
		case <-timer.C:
			fmt.Println("Timer time end")
			break FORLOOP
		case <-ch:
			fmt.Println("Very long function end")
			break FORLOOP
		default:
			fmt.Println("Timer still working")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
