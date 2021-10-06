package main

import (
	"fmt"
	"time"
)

func worker(workerNum int, in <-chan int) {
	for input := range in {
		fmt.Printf("worker-%d take element %d\n", workerNum, input)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("close:", workerNum)
}

func main() {
	poolSize := 10

	PoolEntry := make(chan int)
	for i := 0; i < poolSize; i++ {
		go worker(i, PoolEntry)
	}

	elements := [100]string{}

	for index := range elements {
		PoolEntry <- index
	}

	close(PoolEntry)

	time.Sleep(time.Second)
	fmt.Println("Bye bitch")
}
