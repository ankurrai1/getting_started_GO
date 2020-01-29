package main

import (
  "fmt"
)

// iota is basically a counter which increase the value of it when it is called

const (
  a = iota
  b = iota
  c = iota
  d = iota
)

func main() {
  fmt.Printf("%v  %T",a,a)
  fmt.Printf("%v  %T",b,b)
  fmt.Printf("%v  %T",c,c)
  fmt.Printf("%v  %T",d,d)
}
