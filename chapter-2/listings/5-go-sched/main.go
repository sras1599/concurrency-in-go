package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
    fmt.Println("Hello, World!")
}

func main() {
    go sayHello()
    runtime.Gosched()
    
    println("Finished!")
}