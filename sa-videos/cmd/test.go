package main

import (
	"github.com/Yangusik/sa_videos/queue"
	"time"
)

func publisher() {
	// the publisher publishes the message "1,1" every 500 milliseconds, perpetually
	for {
		if err := queue.Publish("add_q", []byte("1,1")); err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
