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
// the above sleep is needed because main function will run in its own goroutine and as soon as it completed it will come out and
// our go routine didn't execute so we use it to hold main function for some time to make other routine may execute. 
    fmt.Println("done")
}


// NOTE : in other multithreading language they use os threads which Usage a different function stack for new thread
//        but in go goroutine is a lightweight abstraction of threads and mapped with os threads and map them
//        this will use it when processer is free to take with little stack assigned to function
