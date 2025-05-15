package main

import (
	"fmt"
	"sync"
	"time"
)

func handlePlayer(remainingPlayers *int, cond *sync.Cond, playerId int) {
	cond.L.Lock()

	*remainingPlayers--
	fmt.Printf("Player %d connected\n", playerId)

	if *remainingPlayers == 0 {
		cond.Broadcast()
	}

	for *remainingPlayers > 0 {
		fmt.Printf("Waiting for %d players\n", *remainingPlayers)
		cond.Wait()
	}

	cond.L.Unlock()
}

func timeoutGame(seconds int, cond *sync.Cond, timedOut *bool) {
	cond.L.Lock()

	for i := 0; i < seconds; i++ {
		fmt.Printf("Will timeout in %v seconds\n", seconds-(i))
		time.Sleep(1 * time.Second)
	}

	*timedOut = true
	cond.Broadcast()

	cond.L.Unlock()
}

func main() {
	playersToConnect := 10
	numPlayers := playersToConnect

	startGameCond := sync.NewCond(&sync.Mutex{})
	timeoutCond := sync.NewCond(&sync.Mutex{})
	hasTimedOut := false

	go timeoutGame(5, timeoutCond, &hasTimedOut)

	for i := 1; i <= numPlayers; i++ {
		if hasTimedOut {
			fmt.Printf("Timed out. Will start game with %d players\n", (numPlayers - playersToConnect))
			break
		}

		go handlePlayer(&playersToConnect, startGameCond, i)
		time.Sleep(1 * time.Second)
	}

	if !hasTimedOut {
		fmt.Println("All players connected!")
	}
}
