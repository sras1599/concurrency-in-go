package main

import (
	"fmt"
	"sync"
	"time"
)

func receiver(messages <-chan int, wg *sync.WaitGroup) {
	msg := 0

	for msg != -1 {
		msg = <-messages
		fmt.Printf("(%s) Received: %d\n", time.Now().Format("15-04-05"), msg)
	}

	wg.Done()
}

func main() {
	messages := make(chan int, 3)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go receiver(messages, &wg)

	for i := 0; i < 6; i++ {
		bufSize := len(messages)
		fmt.Printf("(%s) Sending %d. Buffer size: %d\n", time.Now().Format("15-04-05"), i, bufSize)
		messages <- i
	}

	messages <- -1
	wg.Wait()
}
