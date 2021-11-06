package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
)

func main(){

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()
	consumer:=kafka.NewReader(kafka.ReaderConfig{
		Brokers:                []string{"localhost:9092"},
		GroupID:                "g1",
		Topic:                  "first_topic",

	})
	defer func() {
		err := consumer.Close()
		if err != nil {
			fmt.Println("Error closing consumer: ", err)
			return
		}
		fmt.Println("Consumer closed")
	}()

	for {
		m, err := consumer.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}
		fmt.Printf("Received message from %s-%d [%d]: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

}
