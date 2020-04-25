package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Threads : ",runtime.GOMAXPROCS(-1) )// will print current threads.
    runtime.GOMAXPROCS(100) // creating 100 threads
    fmt.Println("Threads : ",runtime.GOMAXPROCS(-1) )// will show all threads which i created
}
