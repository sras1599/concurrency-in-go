package main

import (
	"fmt"
)

func receiver(messages chan string) {
	msg := ""

	for msg != "STOP" {
		msg = <-messages
		fmt.Println("Received: ", msg)
	}
}

func send(message string, channel chan string) {
	fmt.Println("Sending ", message)
	channel <- message
}

func main() {
	messages := make(chan string)
	go receiver(messages)

	send("Hello", messages)
	send("World", messages)
	send("STOP", messages)
}
