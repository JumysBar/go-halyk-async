package main

import "fmt"

func main() {
	// var ch chan string
	// <-ch

	// var ch chan string
	// ch <- "Hello world"

	// ch := make(chan string)
	// close(ch)
	// ch <- "Hello world"

	ch := make(chan string)
	close(ch)
	val := <-ch
	fmt.Println(val)
}
