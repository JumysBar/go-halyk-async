package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		close(c)
		fmt.Println("Oops! I received Ctrl+C signal!")
		os.Exit(1)
	}()
	for {
		// infinity loop
	}
}
