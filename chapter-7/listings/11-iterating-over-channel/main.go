package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func done(who string, wg *sync.WaitGroup) {
	fmt.Printf("%s finished\n", who)
	wg.Done()
}

func receiver(messages <-chan int, wg *sync.WaitGroup) {
	defer done("Receiver", wg)

	for msg := range messages {
		fmt.Println("Message received: ", msg)
	}
}

func sender(messages chan<- int, wg *sync.WaitGroup) {
	defer done("Sender", wg)

	for i := 0; i < 10; i++ {
		msg := rand.Intn(42)
		messages <- msg
		time.Sleep(time.Duration(0.5 * float64(time.Second)))
	}

	close(messages)
}

func main() {
	messages := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go sender(messages, &wg)
	go receiver(messages, &wg)
	wg.Wait()
}
