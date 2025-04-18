package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(seconds *int, mu *sync.RWMutex) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)

		mu.Lock()
		*seconds--
		mu.Unlock()
	}
}

func main() {
	mu := sync.RWMutex{}
	seconds := 5

	go countdown(&seconds, &mu)
	for seconds > 0 {
		time.Sleep(500 * time.Millisecond)

		mu.RLock()
		fmt.Printf("Seconds left: %d\n", seconds)
		mu.RUnlock()
	}
}
