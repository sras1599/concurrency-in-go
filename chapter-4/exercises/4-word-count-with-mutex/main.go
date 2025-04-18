package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

func getResponse(url string) string {
	response, _ := http.Get(url)

	if response.StatusCode != 200 {
		return ""
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	return string(body)
}

func isWord(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+$`)

	return re.MatchString(s)
}

func countWords(url string, frequency *sync.Map, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	content := getResponse(url)
	words := strings.Fields(content)

	for _, word := range words {
		if !isWord(word) {
			continue
		}

		word = strings.ToLower(word)

		mu.Lock()
		count, _ := frequency.LoadOrStore(word, 0)
		frequency.Store(word, count.(int)+1)
		mu.Unlock()
	}
}

func main() {
	var frequencyMap sync.Map
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1000; i <= 1005; i++ {
		wg.Add(1)
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, &frequencyMap, &wg, &mu)
	}

	wg.Wait()

	frequencyMap.Range(func(word, count any) bool {
		if count.(int) > 500 {
			fmt.Printf("Count of %s: %d\n", word, count)
		}

		return true
	})
}
