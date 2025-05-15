package main

import "sync"

type Semaphore struct {
	permits int // number of concurrent executions available
	cond    *sync.Cond
}

func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{permits: permits, cond: sync.NewCond(&sync.Mutex{})}
}

func (sem *Semaphore) Acquire() {
	sem.cond.L.Lock()

	for sem.permits <= 0 {
		sem.cond.Wait()
	}

	sem.permits--
	sem.cond.L.Unlock()
}

func (sem *Semaphore) Release() {
	sem.cond.L.Lock()

	sem.permits++
	sem.cond.Signal() // send notice to all goroutines waiting to acquire a permit

	sem.cond.L.Unlock()
}
