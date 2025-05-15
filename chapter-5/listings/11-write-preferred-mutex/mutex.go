package main

import "sync"

type WritePrefferingMutex struct {
	readers        int
	writersWaiting int
	writerActive   bool
	cond           *sync.Cond
}

func NewMutex() *WritePrefferingMutex {
	return &WritePrefferingMutex{cond: sync.NewCond(&sync.Mutex{})}
}

func (mu *WritePrefferingMutex) Lock() {
	mu.cond.L.Lock()
}

func (mu *WritePrefferingMutex) Unlock() {
	mu.cond.L.Unlock()
}

func (mu *WritePrefferingMutex) ReadLock() {
	mu.Lock()

	for mu.writersWaiting > 0 || mu.writerActive {
		mu.cond.Wait()
	}

	mu.readers++
	mu.Unlock()
}

func (mu *WritePrefferingMutex) WriteLock() {
	mu.Lock()
	mu.writersWaiting++

	for mu.readers > 0 || mu.writerActive {
		mu.cond.Wait()
	}

	mu.writersWaiting--
	mu.writerActive = true

	mu.Unlock()
}

func (mu *WritePrefferingMutex) ReadUnlock() {
	mu.Lock()
	mu.readers--

	if mu.readers == 0 {
		mu.cond.Broadcast()
	}

	mu.Unlock()
}

func (mu *WritePrefferingMutex) WriteUnlock() {
	mu.Lock()

	mu.writerActive = false
	mu.cond.Broadcast()

	mu.Unlock()
}
