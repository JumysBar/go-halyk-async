package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrSSLInvalid = errors.New("SSL certificate not valid")
)

// ClientStruct - структура имитирующая структуру для сервера
type ClientStruct struct {
	serverName string
	outputCh   chan<- int
	inputCh    <-chan int
}

// ServerImitation - функция для имитации работы некоторого сервера.
// Запросы к "серверу" приходят в канал input
// "Ответы" от сервера посылаются в канал output
func ServerImitation(input <-chan int, output chan<- int) {
	for i := range input {
		fmt.Println("Request received")
		time.Sleep(10 * time.Second)
		output <- i + 1
	}
}

// Имитируем проверку сертификата сервера
func (cl *ClientStruct) CheckSSL(ctx context.Context) error {
	time.Sleep(5 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Context done in CheckSSL function")
		return fmt.Errorf("CheckSSL: %w", ctx.Err())
	default:
		fmt.Println("Context not done yet")
	}
	if cl.serverName != "valid" {
		return fmt.Errorf("CheckSSL: %w", ErrSSLInvalid)
	}
	return nil
}

// Имитируем доставку запроса до сервера
func (cl *ClientStruct) RoundTrip(ctx context.Context, output chan<- int, data int) {
	// "Доставка" данных
	cl.outputCh <- data

	select {
	case <-ctx.Done():
		fmt.Printf("Context done. Error: %s\n", ctx.Err())
		output <- 0
	case response := <-cl.inputCh:
		fmt.Printf("Return response: %d\n", response)
		output <- response
	}
}

// Имитируем отправку запроса
func (cl *ClientStruct) SendRequest(ctx context.Context, data int) (int, error) {
	response := make(chan int, 1)

	cnlCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go cl.RoundTrip(cnlCtx, response, data)

	timeout := 6 * time.Second

	val := ctx.Value("timeout")
	if val != nil {
		timeout = val.(time.Duration) * time.Second
	}

	timeoutCtx, _ := context.WithTimeout(ctx, timeout)

	if err := cl.CheckSSL(timeoutCtx); err != nil {
		return 0, fmt.Errorf("SendRequest: %w\n", err)
	}

	resp := <-response
	return resp, nil
}

func main() {
	connFromClientToServer := make(chan int)
	connFromServerToClient := make(chan int)
	defer close(connFromClientToServer)
	defer close(connFromServerToClient)

	client := &ClientStruct{
		// serverName: "valid",
		serverName: "invalid",
		inputCh:    connFromServerToClient,
		outputCh:   connFromClientToServer,
	}

	go ServerImitation(connFromClientToServer, connFromServerToClient)

	ctx := context.Background()
	// ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	// ctx := context.WithValue(context.Background(), "timeout", time.Duration(1))

	response, err := client.SendRequest(ctx, 1)
	if err != nil {
		if errors.Is(err, ErrSSLInvalid) {
			fmt.Printf("Server SSL invalid. Error: %s\n", err)
		} else {
			fmt.Printf("main: %s\n", err)
		}
		return
	}

	fmt.Printf("Response from server: %d\n", response)
}
