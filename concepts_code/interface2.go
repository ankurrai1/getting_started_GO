// here we explained use of interface of Interface
package main

import (
  "fmt"
)

func main()  {

}
//first Interface

type Writer interface{
  Write([]byte)(int, error) // method name then input type output type and error
}

// closer interface

type Closer interface{
  Close()(error)
}

// composing two interfaces together

type WriterCloser interface {
  Writer
  Closer
}

// above is same as embedding structs in sturct to impliment 
