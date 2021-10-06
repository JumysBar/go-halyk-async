package main

import (
	"fmt"
	"sync"
	"time"
)

func workerWithWg(wg *sync.WaitGroup, workerNum int, in <-chan int) {
	//wg.Add(1) //страшная ошибка
	defer wg.Done()

	for input := range in {
		fmt.Printf("worker-%d take element %d\n", workerNum, input)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("close:", workerNum)
}

func main() {
	wg := &sync.WaitGroup{}
	poolSize := 10

	PoolEntry := make(chan int)

	wg.Add(poolSize)
	for i := 0; i < poolSize; i++ {
		//wg.Add(1) // как альтернатива для ADD
		go workerWithWg(wg, i, PoolEntry)
	}

	elements := [100]string{}

	for index := range elements {
		PoolEntry <- index
	}
	close(PoolEntry)

	wg.Wait()
	fmt.Println("Bye bitch")
}
