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

	if b.waiting == b.size {
		b.waiting = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}

	b.cond.L.Unlock()
}
