// A goroutine is a lightweight thread of execution it is also known as green thread.
package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

// the usual way, running it synchronously.
// but in through go routine we provide a new thread to execute concurrently.

func main() {
  f("direct")
  go f("goroutine")
  go func(msg string) {
      fmt.Println(msg)
    }("going")
// above two function calls are running asynchronously in separate goroutines now.
    time.Sleep(time.Second)
    fmt.Println("done")
}
