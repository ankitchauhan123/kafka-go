package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "github.com/ankitchauhan123/kafka-go/cmd/utils"



func main(){
	// get the bootstrap server, topic to which data needs to be published

	bootstrapServer := flag.String(utils.BootstrapServers, "localhost:9092", "bootstrap server")
	topic := flag.String(utils.Topic, "test", "topic")
	flag.Parse()


	fmt.Println("Using BootStrap Server:",*bootstrapServer)
	fmt.Println("Using Topic:",*topic)
	fmt.Println("Process ID",os.Getpid())


	signals:=make (chan os.Signal)
	signal.Notify(signals,syscall.SIGINT,syscall.SIGKILL)
	ctx,cancel:=context.WithCancel(context.Background())

	go func(){
		fmt.Println("Starting go func to listen to signals...")
		sig := <-signals
		fmt.Println("Got a signal",sig)
		cancel()
		os.Exit(1)
	}()


	w:= &kafka.Writer{
		Addr:         kafka.TCP(*bootstrapServer),
		Topic:        *topic,
		RequiredAcks: kafka.RequireOne,
		Async:        true, // make the writer asynchronous
		Completion: func(messages []kafka.Message, err error) {
			if err!=nil{
				panic("Error in publishing messages")
				log.Fatal(err)
			}else{
				fmt.Println("Message successfully published", len(messages))
			}

		},
	}

	i:=0
	for {
		msg:=fmt.Sprintf("Message ID : %d",i)
		fmt.Println("Printing message ",msg)
		i++
		w.WriteMessages(ctx,kafka.Message{
			Value:         []byte(msg),
		})
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}

}
