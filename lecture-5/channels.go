package main

import (
	"fmt"
	"time"
)

func GoroutineWork(input chan string) {
	time.Sleep(3 * time.Second)
	inputValue, ok := <-input
	if !ok {
		fmt.Println("Channel closed")
		return
	}
	fmt.Printf("Receive from channel: %s\n", inputValue)
}

func WorkWithChan() {
	var ch chan string

	ch = make(chan string)

	go GoroutineWork(ch)

	ch <- "Hello world!"

	fmt.Println("Message after writing in channel")

	close(ch)

	// "Hello world" -> ch
}

func WorkWithClosedChan() {
	var ch chan string

	ch = make(chan string)

	go GoroutineWork(ch)

	close(ch)

	fmt.Println("Message after closing channel")

	time.Sleep(4 * time.Second)
}

func main() {
	WorkWithChan()
	WorkWithClosedChan()
}
