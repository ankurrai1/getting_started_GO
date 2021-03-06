// Channels are the pipes that connect concurrent goroutines. You can send values into channels
// from one goroutine and receive those values into another goroutine.
// Channels are designed to taransfer data between two or more go routines.
package main

import (
  "fmt"
  "sync"
)

// func main() {
// // Create a new channel with make(chan val-type). Channels are typed by the values they convey.
//     messages := make(chan string)
//
// // Send a value into a channel using the channel <- syntax.
// // Here we send "ping" to the messages channel we made above, from a new goroutine.
//     go func() { messages <- "ping" }()
//     msg := <-messages
//     fmt.Println(msg)
// }

// main function is itself a go routine so we are sending data from one go routine to another

// OTHER Example is as following you have to comment out above main function to check this out.

var wg = sync.WaitGroup{}

func main()  {
  cha := make(chan int) // make(type datatype, buffer channel size Number only)
  //its strictly typed it can only send data with predefined type .
  wg.Add(2)

  go func (cha <-chan int)  { // read only channel can only read data from channel
    num := <- cha  // getting value from channel
    fmt.Println(num)
    wg.Done()
  }(cha)

  go func (cha chan<- int)  { // send to only channel will send data to channel only
    cha <- 42 // sending value to channel
    wg.Done()
  }(cha)

  wg.Wait()
}
