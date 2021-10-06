package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	poolSize := int64(10)

	sem := semaphore.NewWeighted(poolSize)
	ctx := context.Background()

	elements := [100]string{}

	for index := range elements {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println(fmt.Errorf("wait for resources: %w", err))
		}
		go workerWithSema(sem, index)
	}
	if err := sem.Acquire(ctx, poolSize); err != nil {
		fmt.Println(fmt.Errorf("wait for resources: %w", err))
	}
	fmt.Println("Bye bitch")
}

func workerWithSema(sem *semaphore.Weighted, workValue int) {
	defer sem.Release(1)

	fmt.Printf("some worker take element %d\n", workValue)
	time.Sleep(time.Millisecond * 100)
}
