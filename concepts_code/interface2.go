// here we explained use of interface of Interface
package main

import (
  "fmt"
  "bytes"
)

func main()  {
  var wc WriterCloser = NewBufferedWriterCloser()
  wc.Write([]byte("This line will be print in chunk of 8 and at close it will flush out all left in"))
  wc.Close()
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

//defining type on which interface will work

type BufferedWriterCloser struct{
  buffer *bytes.Buffer
}


func (bwc *BufferedWriterCloser) Write(data []byte)(int, error)  {
  n , err := bwc.buffer.Write(data)
  if err != nil {
    return 0 ,err
  }
  v := make([]byte, 8)
  for bwc.buffer.Len() > 8{
    _,err = bwc.buffer.Read(v)
    if err != nil {
      return 0 ,err
    }
    _, err := fmt.Println(string(v))
    if err != nil {
      return 0 ,err
    }
  }
  return n, nil
}

func (bwc *BufferedWriterCloser) Close()(error)  {
  for bwc.buffer.Len() > 0 {
    data := bwc.buffer.Next(8)
    _,err := fmt.Println(string(data))
    if err != nil {
      return 0 ,err
    }
  }
  return nil
}


func NewBufferedWriterCloser() *BufferedWriterCloser {
  return &BufferedWriterCloser{
    buffer: *bytes.NewBuffer([]bytes{})
  }
}
