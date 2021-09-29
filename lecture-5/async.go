package main

import (
	"fmt"
	"time"
)

func main() {
	// Пример с тем, что горутины - асинхронный подход
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Printf("Goroutine %d work begin\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d work end\n", i)
		}(i)
	}

	fmt.Scanln()
}
