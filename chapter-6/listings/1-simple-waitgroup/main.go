package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sleepRandom(wg *sync.WaitGroup) {
	sleepDuration := rand.Intn(5)

	time.Sleep(time.Duration(sleepDuration) * time.Second)
	fmt.Printf("Finished after sleeping for %d seconds\n", sleepDuration)

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sleepRandom(&wg)
	}

	wg.Wait()
	fmt.Println("Finished!")
}
