package main

import (
	"fmt"
	"time"
)

// input - read-only channel
func GoroutineWork(input <-chan int) {
	for chValue := range input {
		time.Sleep(time.Second)
		fmt.Printf("Value from channel: %d. Channel len: %d\n", chValue, len(input))
	}
	fmt.Println("End of goroutine")
}

// output - write only channel
func OnlyWriteInChan(output chan<- int) {
	for i := 0; i < 10; i++ {
		output <- i
		fmt.Printf("Value %d was sent in channel\n", i)
	}
}

func main() {
	ch := make(chan int, 3)

	go GoroutineWork(ch)

	OnlyWriteInChan(ch)

	fmt.Println("Message after writing in channel")

	close(ch)

	time.Sleep(5 * time.Second)
}
