// In this section,I have written a program which has race condition

// RACE CONDITION :
// A race condition is a situation that may occur inside a critical section.
// This happens when the result of multiple thread execution in critical section differs
// according to the order in which the threads execute

package main
import (
    "fmt"
    "sync"
    )
var x  = 0
func increment(wg *sync.WaitGroup) {
    x = x + 1
    wg.Done()
}
func main() {
    var w sync.WaitGroup
    for i := 0; i < 1000; i++ {
        w.Add(1)
        go increment(&w)
    }
    w.Wait()
    fmt.Println("final value of x", x)
}


// Oprating System CONCEPTS : critical section
// When a program runs concurrently, the parts of code which modify
// shared resources should not be accessed by multiple Goroutines at the same time.
// This section of code which modifies shared resources is called critical section
