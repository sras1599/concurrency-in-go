package main

import "fmt"

func main() {
	sem := NewSemaphore(0)

	for i := 0; i < 1; i++ {
		go doWork(sem)
		fmt.Println("Waiting for child goroutine")

		sem.Acquire()
		fmt.Println("Child goroutine finished")
	}
}

func doWork(sem *Semaphore) {
	fmt.Println("Work started")
	fmt.Println("Work finished")

	sem.Release()
}
