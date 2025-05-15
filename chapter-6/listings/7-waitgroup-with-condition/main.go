package main

import "sync"

type WaitGroup struct {
	cond *sync.Cond
	size int
}

func newWaitGroup() *WaitGroup {
	return &WaitGroup{cond: sync.NewCond(&sync.Mutex{})}
}

func (wg *WaitGroup) Add(delta int) {
	wg.cond.L.Lock()

	wg.size += delta

	wg.cond.L.Unlock()
}

func (wg *WaitGroup) Wait() {
	wg.cond.L.Lock()

	for wg.size > 0 {
		wg.cond.Wait()
	}

	wg.cond.L.Unlock()
}

func (wg *WaitGroup) Done() {
	wg.cond.L.Lock()

	wg.size--
	if wg.size == 0 {
		wg.cond.Broadcast()
	}

	wg.cond.L.Unlock()
}
