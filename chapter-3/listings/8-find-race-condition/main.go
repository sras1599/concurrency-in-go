// Question asked in the book: Is there a race condition in this program?
// My answer: Yes
// Explanation: Each iteration of the for loop spawns a new goroutine which
// is supposed to update the `i`th index in the shared `nextNum` array.
// For the program to stop, the last element of the array must be updated
// which means that every goroutine has to update a unique index. The race condition
// can occur when the execution of a goroutine stops after it has read the value of a certain index
// to be 0, but before it can write to `nextNum` and update it. If at this point another
// goroutine reads the value of the first `i` with a 0 value, both of these goroutines
// will end up updating the same index in the `nextNum` array
package main

import (
	"fmt"
	"time"
)

func addNextNumber(nextNum *[101]int) {
	i := 0

	for nextNum[i] != 0 {
		i++
	}

	nextNum[i] = nextNum[i-1] + 1
}

func main() {
	nextNum := [101]int{1}

	for i := 0; i < 100; i++ {
		go addNextNumber(&nextNum)
	}

	for nextNum[100] == 0 {
		println("Waiting for goroutines to complete")
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println(nextNum)
}
