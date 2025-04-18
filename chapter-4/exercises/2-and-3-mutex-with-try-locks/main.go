package main

import "sync"

type CustomRWMutex struct {
	readers    int
	rLock      sync.Mutex
	globalLock sync.Mutex
}

func (mu *CustomRWMutex) ReadLock() {
	mu.rLock.Lock()
	mu.readers++

	if mu.readers == 1 {
		mu.globalLock.Lock()
	}

	mu.rLock.Unlock()
}

func (mu *CustomRWMutex) WriteLock() {
	mu.globalLock.Lock()
}

func (mu *CustomRWMutex) ReadUnlock() {
	mu.rLock.Lock()
	mu.readers--

	if mu.readers == 0 {
		mu.globalLock.Unlock()
	}

	mu.rLock.Unlock()
}

func (mu *CustomRWMutex) WriteUnlock() {
	mu.globalLock.Unlock()
}

func (mu *CustomRWMutex) TryLock() bool {
	return mu.globalLock.TryLock()
}

func (mu *CustomRWMutex) TryReadLock() bool {
	if mu.rLock.TryLock() {
		success := true

		if mu.readers == 0 {
			success = mu.TryLock()
		}

		if success {
			mu.readers++
		}

		mu.rLock.Unlock()
		return success
	} else {
		return false
	}
}
