package main

import (
	"fmt"
	"time"
)

// Эта горутина выполняет мультиплексированный ввод с нескольких каналов
func goroutineWorkInput(ch1 <-chan int, ch2 <-chan string) {
	for {
		select {
		case <-ch1: // Если значение с канала не важно
			fmt.Println("Receive from ch1")
		case val := <-ch2: // Если значение с канала важно
			fmt.Printf("Receive from ch2 %s\n", val)
		default:
			fmt.Println("No one writes to any channel")
			time.Sleep(time.Second)
		}
	}
}

// Функция создает каналы, запускает горутину и с некоторой периодичностью записывает значения
func functionWithOutputChannels() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go goroutineWorkInput(ch1, ch2)

	time.Sleep(time.Second)

	ch1 <- 1

	time.Sleep(time.Second)

	ch2 <- "Hello"

	close(ch1)
	close(ch2)
}

// Горутина, которая выполняет мультиплексированную запись в два канала
func goroutineWorkOutput(ch1 chan<- int, ch2 chan<- string) {
	for {
		select {
		case ch1 <- 1:
			fmt.Println("Value sent to ch1")
		case ch2 <- "Hello world":
			fmt.Println("Value sent to ch2")
		default:
			fmt.Println("No one reads from any channel")
			time.Sleep(time.Second)
		}
	}
}

// Функция создает каналы, запускает горутину и с некоторой периодичностью считывает значения
func functionWithInputChannels() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go goroutineWorkOutput(ch1, ch2)

	time.Sleep(time.Second)

	val1 := <-ch1
	fmt.Printf("Read value from ch1: %d\n", val1)

	time.Sleep(time.Second)

	val2 := <-ch2
	fmt.Printf("Read value from ch2: %s\n", val2)

	close(ch1)
	close(ch2)
}

func main() {
	// functionWithOutputChannels()

	functionWithInputChannels()
	time.Sleep(time.Second)
}
