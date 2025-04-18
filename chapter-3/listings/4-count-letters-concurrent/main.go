// the concurrent implement in this program can be affected by a race condition
// due to which the results might be inaccurate
// refer to listing 5 of chapter-4 which implements the same program using
// mutexes to guarantee exclusive access to a goroutine when updating a variable
// in shared memory
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLettersConcurrent(url string, frequency []int, wg *sync.WaitGroup) {
	defer wg.Done()
	response, _ := http.Get(url)

	if response.StatusCode != 200 {
		panic("Server returned error status code: " + response.Status)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	for _, byteChar := range body {
		char := strings.ToLower(string(byteChar))
		charIndex := strings.Index(allLetters, char)

		if charIndex >= 0 {
			frequency[charIndex] += 1
		}

	}

	fmt.Printf("Completed for url: %s\n", url)
}

func main() {
	var frequency = make([]int, 26)
	var wg sync.WaitGroup

	for i := 1000; i <= 1030; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLettersConcurrent(url, frequency, &wg)
	}

	wg.Wait()

	for i, c := range allLetters {
		fmt.Printf("Count of %c: %d\n", c, frequency[i])
	}
}
