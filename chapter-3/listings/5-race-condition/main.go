// reproduces a race condition by making 2 goroutines modify a variable in shared memory
// Note: The goroutines don't even have to run simultaneously for this to happen
// How is the race condition produced?
// in some iterations, the value of the variable that one goroutine reads will be
// changed by the other goroutine before the gorotuine that initially read it updates
// it. This can happen if the kernel (or go's scheduler), switches the context
// at the "wrong" moment. The goroutine that updates the value of the variable at the
// last would've thus ignored any changes to the variable that the other goroutine might
// have made
package main

import (
	"fmt"
	"sync"
)

func stingy(money *int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 0; i < 1000000; i++ {
        // after compilation, this piece of code is not atomic, which means
        // that this contains multiple instructions, and a context switch can
        // happen after any one of them
        *money += 10;
    }

    println("Stingy done")
}

func spendy(money *int, wg *sync.WaitGroup) {
    defer wg.Done()

    for i := 0; i < 1000000; i++ {
        // after compilation, this piece of code is not atomic, which means
        // that this contains multiple instructions, and a context switch can
        // happen after any one of them
        *money -= 10;
    }

    println("Spendy done")
}

func main() {
    money := 1000
    var wg sync.WaitGroup
    
    
    wg.Add(2)
    go stingy(&money, &wg)
    go spendy(&money, &wg)

    wg.Wait()
    fmt.Printf("Money left: %d\n", money)
}