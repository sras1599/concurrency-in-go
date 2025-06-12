package main

import (
	"fmt"
	"math/rand"
	"time"
)

func receiver(messages <-chan int) {
	for {
		msg, more := <-messages
		fmt.Printf("Received message: %d\nSender closed: %t\n\n", msg, !more)
		time.Sleep(1 * time.Second)
		if !more {
			fmt.Println("No more messages to be received")
			return
		}
	}
}

func main() {
	messages := make(chan int)
	go receiver(messages)

	for i := 0; i < 5; i++ {
		msg := rand.Intn(42)
		messages <- msg
	}
	close(messages)
	time.Sleep(4 * time.Second)
}
