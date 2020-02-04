// A panic typically means something went unexpectedly wrong.
// Mostly we use it to fail on errors that shouldn’t occur during normal operation,
// or that we aren’t prepared to handle . because we don't have Exceptions handling Here

package main

import "os"

func main() {

  panic("a problem")
  //A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
    _, err := os.Create("/file")
    if err != nil {
        panic(err)
    }
}
