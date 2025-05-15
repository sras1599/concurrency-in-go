package main

import (
	"fmt"
	"time"
)

func workAndWait(name string, workDuration int, b *Barrier) {
	start := time.Now()

	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(workDuration) * time.Second)

		fmt.Printf("%v %s: waiting for barrier\n", time.Since(start), name)
		b.Wait()
	}
}

func main() {
	b := newBarrier(2)

	go workAndWait("foo", 2, b)
	go workAndWait("bar", 4, b)

	time.Sleep(100 * time.Second)
}
