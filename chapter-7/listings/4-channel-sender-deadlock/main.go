package main

import (
	"fmt"
	"time"
)

func sendMessage(message string, channel chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Sending (but actually not): ", message)
	// channel <- message
}

func main() {
	messages := make(chan string)

	go sendMessage("Hello", messages)

	msg := <-messages
	fmt.Println("Received: ", msg)
}
