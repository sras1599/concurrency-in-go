package main

import "sync"

type Barrier struct {
	cond    *sync.Cond
	size    int
	waiting int
}

func newBarrier(size int) *Barrier {
	cond := sync.NewCond(&sync.Mutex{})

	return &Barrier{cond: cond, size: size}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waiting++

	for b.waiting != b.size {
		b.cond.Wait()
	}

	b.waiting = 0
	b.cond.Broadcast()

	b.cond.L.Unlock()
}
