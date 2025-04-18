package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLettersConcurrent(url string, frequency []int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	response, _ := http.Get(url)

	if response.StatusCode != 200 {
		panic("Server returned error status code: " + response.Status)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	// it's better to lock the mutex outside the loop as doing it inside will
	// repeat the process for each letter, which will take more time
	mu.Lock()
	for _, byteChar := range body {
		char := strings.ToLower(string(byteChar))
		charIndex := strings.Index(allLetters, char)

		if charIndex >= 0 {
			frequency[charIndex] += 1
		}

	}
	mu.Unlock()

	fmt.Printf("Completed for url: %s\n", url)
}

func main() {
	start := time.Now()
	var frequency = make([]int, 26)
	var wg sync.WaitGroup
	var mu = sync.Mutex{}

	for i := 1000; i <= 1030; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLettersConcurrent(url, frequency, &wg, &mu)
	}

	wg.Wait()

	for i, c := range allLetters {
		fmt.Printf("Count of %c: %d\n", c, frequency[i])
	}

	fmt.Printf("Completed in %f seconds\n", time.Since(start).Seconds())
}
