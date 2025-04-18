package main

import (
	"fmt"
	"sync"
	"time"
)

func handlePlayer(remainingPlayers *int, cond *sync.Cond, playerId int) {
	fmt.Printf("Player %d connected\n", playerId)

	cond.L.Lock()
	*remainingPlayers--

	if *remainingPlayers == 0 {
		cond.Broadcast()
	}

	for *remainingPlayers > 0 {
		fmt.Printf("Waiting for %d players\n", *remainingPlayers)
		cond.Wait()
	}

	cond.L.Unlock()
}

func main() {
	numPlayers := 6
	cond := sync.NewCond(&sync.Mutex{})

	for i := 1; i < 7; i++ {
		go handlePlayer(&numPlayers, cond, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All players connected!")
}
