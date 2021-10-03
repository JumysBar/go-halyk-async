package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	atomicInt := int32(0)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&atomicInt, 1)
		}()
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println(atomicInt)
}

//fmt.Println(atomic.LoadInt32(&atomicInt))
