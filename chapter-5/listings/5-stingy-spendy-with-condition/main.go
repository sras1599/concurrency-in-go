package main

import (
	"fmt"
	"os"
	"sync"
)

func spendy(money *int, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	timesSubtraced := 0
	for i := 0; i < 200000; i++ {
		cond.L.Lock()

		for *money < 50 {
			cond.Wait()
		}

		*money -= 50
		timesSubtraced++
		if *money < 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		cond.L.Unlock()
	}

	fmt.Printf("Spendy done! Subtracted %d times\n", timesSubtraced)
}

func stingy(money *int, cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	timesAdded := 0
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		timesAdded++

		cond.Signal()
		cond.L.Unlock()
	}

	fmt.Printf("Stingy done! Added %d times\n", timesAdded)
}

func main() {
	money := 100
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go stingy(&money, cond, &wg)
	go spendy(&money, cond, &wg)
	wg.Wait()

	fmt.Printf("Money in account: %d\n", money)
}
