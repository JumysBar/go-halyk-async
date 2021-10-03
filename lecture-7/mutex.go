package main

import (
	"fmt"
	"sync"
)

func main() {
	mx := &sync.Mutex{}
	_map := make(map[string]int)

	for i := 0; i < 1000; i++ {
		go func(i int, mx *sync.Mutex) {
			mx.Lock()
			defer mx.Unlock()
			_map["race"] = i
		}(i, mx)
	}

	fmt.Printf("%#v", _map["race"])
	fmt.Scanln()
}
