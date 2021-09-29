package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10000; i++ {
		go func() {
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d work\n", i)
			sum++
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Sum: %d\n", sum)

	// little lifehack
	fmt.Scanln()
}
