package main

import (
	"fmt"
	"sync"
	"time"
)

type WeightedSemaphore struct {
	permits int // number of concurrent executions available
	cond    *sync.Cond
}

func NewSemaphore(permits int) *WeightedSemaphore {
	return &WeightedSemaphore{permits: permits, cond: sync.NewCond(&sync.Mutex{})}
}

func (sem *WeightedSemaphore) Acquire(permits int) {
	sem.cond.L.Lock()

	for sem.permits-permits < 0 {
		sem.cond.Wait()
	}

	sem.permits -= permits
	sem.cond.L.Unlock()
}

func (sem *WeightedSemaphore) Release(permits int) {
	sem.cond.L.Lock()

	sem.permits += permits
	sem.cond.Signal() // send notice to all goroutines waiting to acquire a permit

	sem.cond.L.Unlock()
}

func main() {
	sem := NewSemaphore(3)
	sem.Acquire(2)
	fmt.Println("Parent thread acquired semaphore")
	go func() {
		sem.Acquire(2)
		fmt.Println("Child thread acquired semaphore")
		sem.Release(2)
		fmt.Println("Child thread released semaphore")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Parent thread releasing semaphore")
	sem.Release(2)
	time.Sleep(1 * time.Second)
}
