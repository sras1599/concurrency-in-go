package main

import (
	"fmt"
	"time"
)

func main() {
	mu := NewMutex()

	for i := 0; i < 2; i++ {
		go func() {
			for {
				mu.ReadLock()
				time.Sleep(1 * time.Second)

				fmt.Println("Read done")
				mu.ReadUnlock()
			}
		}()
	}

	time.Sleep(1 * time.Second)
	mu.WriteLock()
	fmt.Println("Write finished")
}
