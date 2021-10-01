package main

import (
	"fmt"
	"time"
)

type Service struct {
	HealthCheckChan chan<- struct{}
}

func (s *Service) Run() {
	time.Sleep(5 * time.Second)
	s.HealthCheckChan <- struct{}{}
	fmt.Println("Service shutdown")
}

func main() {
	HealthCheckChan := make(chan struct{}, 1)
	service := &Service{
		HealthCheckChan: HealthCheckChan,
	}
	go service.Run()

	ticker := time.NewTicker(time.Second)

LOOP:
	for range ticker.C {
		select {
		case <-HealthCheckChan:
			fmt.Println("Service healthcheck failed")
			ticker.Stop()
			break LOOP
		default:
			fmt.Println("Everything is allright. Service is up")
		}
	}
}
