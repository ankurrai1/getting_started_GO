package main

// implimenting interface on integer type in go
import (
  "fmt"
)


func main()  {
  c := intCounter(0)
  var inc Incrementer = &c // allocating intCounter to Interface
  for i:=0; i<10; i++{
    fmt.Println(inc.Increment()) // calling interface method
  }

}

type Incrementer interface{ //declaration of interface
  Increment()(int)
}

type intCounter int // declaration of datatype on which interface will work


// Implimentation of real method
//      data type     methodName  returnType
func (ic *intCounter) Increment() int {
  *ic ++
  return int(*ic)
}
