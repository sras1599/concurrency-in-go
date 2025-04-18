package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
    response, _ := http.Get(url)
    defer response.Body.Close()

    if response.StatusCode != 200 {
        panic("Server returned error status code: " + response.Status)
    }

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

    for i := 1000; i <= 1030; i ++ {
        url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
        countLetters(url, frequency)
    }

    for i, c := range allLetters {
        fmt.Printf("Count of %c: %d\n", c, frequency[i])
    }
}
