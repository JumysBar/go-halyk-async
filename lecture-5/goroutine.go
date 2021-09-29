package main

import (
	"fmt"
)

func threadFunc() {
	fmt.Println("This code is running in a thread")
}

func main() {
	go threadFunc()
	fmt.Println("This code is running in main thread")

	// little lifehack
	fmt.Scanln()
}
