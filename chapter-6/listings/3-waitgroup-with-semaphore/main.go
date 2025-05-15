package main

import "fmt"

const workIterations = 4

type WaitGroup struct {
	sem *Semaphore
}

func newWaitGroup(size int) *WaitGroup {
	// Why is the semaphore created with `1 - size` permits?
	//
	// We can only acquire the semaphore and end the wait for a waitgroup
	// when the number of permits available in the semaphore is at least 1.
	// By initializing the semaphore with a negative value of `1 - size`, we
	// can wait for `size` calls to the `Done` function before we can finish
	// waiting and acquire the semaphore. This acts as the signal which tells
	// that all tasks have finished
	return &WaitGroup{sem: NewSemaphore(1 - size)}
}

func (wg WaitGroup) Wait() {
	wg.sem.Acquire()
}

func (wg WaitGroup) Done() {
	wg.sem.Release()
}

func doWork(wg *WaitGroup) {
	fmt.Println("Done working")
	wg.Done()
}

func main() {
	wg := newWaitGroup(workIterations)

	for i := 1; i <= workIterations; i++ {
		go doWork(wg)
	}
	wg.Wait()

	fmt.Println("Finished")
}
