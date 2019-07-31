//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

var canal = make(chan *Tweet)
var final = make(chan bool)

// func producer(stream Stream) (tweets []*Tweet) {
func producer(stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			// return tweets
			break
		}

		// tweets = append(tweets, tweet)
		canal <- tweet
	}
	final <- true
}

// func consumer(tweets []*Tweet) {
func consumer() {
	// for _, t := range tweets {
	for {
		t := <-canal
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	// tweets := producer(stream)
	go producer(stream)
	// Consumer
	// consumer(tweets)
	go consumer()
	<-final
	fmt.Printf("Process took %s\n", time.Since(start))
}
