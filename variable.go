package main

import (
  "fmt"
)

func main()  {
  // first type of variable diclaration
  // basically the typed one
  var x int = 5
  var y int = 7
  var sum int = x + y

  // Second type of variable declaration
  // untyped and more clean

  a := 5
  b := 7
  add := a + b

  // Finally printing both result togather to compare

  fmt.Println("with typed variable : ",sum,"\n","With untyped variable : ",add)
}
