package main

import (
	"fmt"
)

func main() {
	_map := make(map[string]int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			_map["race"] = i
		}(i)
	}

	fmt.Printf("%#v", _map)
	fmt.Scanln()
}
