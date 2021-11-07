package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"os/signal"
	"syscall"
)


type TwitterProcessor struct{
	demux twitter.SwitchDemux,


}

func main(){
	fmt.Println("Hi")
	config := oauth1.NewConfig("Cbg2XoAeE6ZyF0wp6YxUeTI3T", "Gn8BxOZ2qmpwTzasGWEfeX4PRyU19KsNoEmmq5DXinKEajVDWT")
	token := oauth1.NewToken("1456936943988580358-YigB0BMUs5BjmJnLBxHGo8xjfpzn5a", "KUZh8u5DFmXsc11rVf9asiuw3WaKsnFWrIU5ymJgHxQVA")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	fmt.Println("Starting Stream...")

	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"yoooo","weekend","modi"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()

}





