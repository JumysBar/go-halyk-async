package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan string)
	ch3 := make(chan int)

	go func(ch1 chan<- bool) {
		ch1 <- true
	}(ch1)

	go func(ch2 chan<- string) {
		ch2 <- "Hello"
	}(ch2)

	go func(ch3 chan<- int) {
		ch3 <- 1
	}(ch3)

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		select {
		case val1 := <-ch1:
			fmt.Printf("Value from ch1: %v\n", val1)
		case val2 := <-ch2:
			fmt.Printf("Value from ch2: %s\n", val2)
		case val3 := <-ch3:
			fmt.Printf("Value from ch3: %d\n", val3)
		}
	}
	close(ch1)
	close(ch2)
	close(ch3)
	time.Sleep(2 * time.Second)
}
