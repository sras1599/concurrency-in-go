package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func copyEvents(events *[]string) []string {
	ret := make([]string, len(*events))
	copy(*events, ret)

	return ret
}

func clientHandler(events *[]string, mu *sync.RWMutex, start time.Time) {
	for i := 0; i < 100; i++ {
		mu.RLock()
		allEvents := copyEvents(events)
		mu.RUnlock()

		elapsed := time.Since(start)

		fmt.Printf("%d events copied in %s\n", len(allEvents), elapsed)
	}
}

func recordMatch(events *[]string, mu *sync.RWMutex) {
	for i := 0; ; i++ {
		mu.Lock()
		*events = append(*events, "Match event "+strconv.Itoa(i))
		mu.Unlock()

		time.Sleep(200 * time.Millisecond)
		fmt.Println("Added match event")
	}
}

func main() {
	mu := sync.RWMutex{}
	matchEvents := make([]string, 0, 10000)

	for i := 0; i < 10000; i++ {
		matchEvents = append(matchEvents, "Match event")
	}

	go recordMatch(&matchEvents, &mu)

	start := time.Now()

	for i := 0; i < 5000; i++ {
		go clientHandler(&matchEvents, &mu, start)
	}

	time.Sleep(100 * time.Second)
}
