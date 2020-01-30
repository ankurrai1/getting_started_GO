package main

import (
  "fmt"
)

// iota is basically a counter which increase the value of it when it is called

const (
  a = iota
  b = iota
  c = iota
  d
)
// iota is limited to each const block
// next const block will have its on iota starting from 0

const (
  e = iota
  f = iota
)

func main() {

  // after printing the value of that function we can see every variable is grater then that

  fmt.Printf("%v  %T",a,a)
  fmt.Printf("%v  %T",b,b)
  fmt.Printf("%v  %T",c,c)
  fmt.Printf("%v  %T",d,d)
  fmt.Printf("%v  %T",e,e)
  fmt.Printf("%v  %T",e,e)

}
