package main

import (
	"fmt"
	"sync"
	"time"
)

var globalMx = &sync.Mutex{}

func workerA(wg *sync.WaitGroup, workValue int, currentValue *int) {
	defer wg.Done()
	defer decreasePoolValue(currentValue)

	fmt.Printf("some worker take element %d\n", workValue)
	time.Sleep(time.Millisecond * 100)
}

func main() {
	wg := &sync.WaitGroup{}

	poolLimitValue := 10
	poolCurrentValue := 0

	elements := [100]string{}

	wg.Add(len(elements))
	for i := range elements {
		for {
			if increasePoolValue(&poolCurrentValue, poolLimitValue) {
				break
			}
		}
		go workerA(wg, i, &poolCurrentValue)
	}

	wg.Wait()
	fmt.Println("Bye bitch")
}

func increasePoolValue(poolValue *int, limit int) (yesYouCan bool) {
	globalMx.Lock()
	defer globalMx.Unlock()

	if *poolValue < limit {
		*poolValue++
		yesYouCan = true
	}

	return yesYouCan
}

func decreasePoolValue(poolValue *int) {
	globalMx.Lock()
	defer globalMx.Unlock()

	*poolValue--
}
