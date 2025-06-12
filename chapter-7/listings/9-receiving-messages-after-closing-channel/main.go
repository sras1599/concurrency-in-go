package main

import (
	"fmt"
	"time"
)

func receiver(messages <-chan int) {
	for {
		msg := <-messages
		fmt.Println("Received message: ", msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	messages := make(chan int)
	go receiver(messages)

	for i := 1; i < 4; i++ {
		fmt.Println("Sending message: ", i)
		messages <- i
	}

	close(messages)
	time.Sleep(3 * time.Second)
}
