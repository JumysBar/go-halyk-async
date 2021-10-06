package main

import (
	"fmt"
	"sync"
)

// для возвышения почитайте - https://habr.com/ru/post/338718/
func main() {
	_map := sync.Map{}

	for i := 0; i < 10000; i++ {
		go func(i int) {
			_map.Store("race", i)
		}(i)
	}

	fmt.Println(_map.Load("race"))
	fmt.Scanln()
}
