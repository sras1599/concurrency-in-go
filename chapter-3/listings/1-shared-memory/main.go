package main

import (
	"fmt"
	"time"
)

func countdown(count *int) {
    for *count > 0 {
        time.Sleep(1 * time.Second)
        *count -= 1
    }
}

func main() {
    count := 5

    go countdown(&count)

    for count > 0 {
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("Count: %d\n", count)
    }
}