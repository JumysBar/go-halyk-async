package main

import (
	"context"
	"fmt"
	"time"
)

type Request struct {
	data int
	ctx  context.Context
}

type Response struct {
	err  error
	data int
}

func (r *Request) GetContext() context.Context {
	return r.ctx
}

func NewRequest(ctx context.Context, data int) *Request {
	return &Request{
		ctx:  ctx,
		data: data,
	}
}

func ServerImitation(input <-chan *Request, output chan<- *Response) {
	var timer *time.Timer
REQLOOP:
	for req := range input {
		fmt.Println("Request received")

		ctx := req.GetContext()

		timer = time.NewTimer(3 * time.Second)

		select {
		case <-ctx.Done():
			// Возврат "ответа"
			output <- &Response{
				err: fmt.Errorf("Request aborted"),
			}
			continue REQLOOP
		case <-timer.C:
			fmt.Println("Timer event")
		}

		// Возврат "ответа"
		output <- &Response{
			err:  nil,
			data: req.data + 1,
		}
	}
}

func main() {
	connFromClientToServer := make(chan *Request)
	connFromServerToClient := make(chan *Response)
	defer close(connFromClientToServer)
	defer close(connFromServerToClient)

	go ServerImitation(connFromClientToServer, connFromServerToClient)

	data := 1
	// ctx := context.Background()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	req := NewRequest(ctx, data)

	connFromClientToServer <- req

	resp := <-connFromServerToClient

	if resp.err != nil {
		fmt.Printf("Error main: %s\n", resp.err)
		return
	}

	fmt.Printf("Response received: %d\n", resp.data)
}
