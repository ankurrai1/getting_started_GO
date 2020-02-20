// In the program above, we spawn 1000 Goroutines. If each increments the value of x by 1,
// the final desired value of x should be 1000. In this section,
//  we will fix the race condition.

package main

import (
    "fmt"
    "sync"
    )

var x  = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
    m.Lock()
    x = x + 1
    m.Unlock()
    wg.Done()
}

func main() {
    var w sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < 1000; i++ {
        w.Add(1)
        go increment(&w, &m)
    }
    w.Wait()
    fmt.Println("final value of x", x)
}
