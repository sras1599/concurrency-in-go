package main

import (
	"fmt"
	"time"
)

func receiver(messages chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Stopping receiver after sleeping for 2 seconds")
}

func sendMessage(message string, channel chan string) {
	fmt.Println("Sending ", message)
	channel <- message
}

func main() {
	messages := make(chan string)
	go receiver(messages)

	sendMessage("Hello", messages)
	sendMessage("World", messages)
	sendMessage("STOP", messages)
}
