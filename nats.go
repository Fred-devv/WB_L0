package main

import (
	"github.com/nats-io/stan.go"
)

// Подключение к NATS Streaming
func connectToNATS() (*stan.Conn, error) {
	//Реализация подключения к NATS Streaming
	sc, err := stan.Connect("test-cluster", "orders-service")
	if err != nil {
		return nil, err
	}
	return sc, nil
}
