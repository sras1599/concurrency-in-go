// showing how mutexes avoid race conditions by making access to a section
// of code exclusive to a single goroutine (at any point of time)
// in this case, the section of code involves the updation of a variable
// that is shared by multiple goroutines
package main

import (
	"fmt"
	"sync"
)

const million = 1_000_000

func stingy(money *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < million; i++ {
		mu.Lock()
		*money += 10
		mu.Unlock()
	}

	println("Stingy done!")
}

func spendy(money *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < million; i++ {
		mu.Lock()
		*money -= 10
		mu.Unlock()
	}

	println("Spendy done!")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(2)
	go stingy(&money, &mutex, &wg)
	go spendy(&money, &mutex, &wg)
	wg.Wait()

	fmt.Printf("Money: %d\n", money)
}
