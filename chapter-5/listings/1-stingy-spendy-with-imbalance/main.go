package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func stingy(money *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		mu.Lock()
		*money += 10
		mu.Unlock()
	}

	fmt.Println("Stingy done")
}

func spendy(money *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		mu.Lock()

		for *money < 50 {
			mu.Unlock()
			time.Sleep(10 * time.Millisecond)
			mu.Lock()
		}

		*money -= 50

		if *money <= 0 {
			fmt.Println("Spent all the money!")
			os.Exit(1)
		}
		mu.Unlock()
	}

	fmt.Println("Spendy done")
}

func main() {
	money := 10000
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(2)
	go stingy(&money, &mu, &wg)
	go spendy(&money, &mu, &wg)
	wg.Wait()
}
