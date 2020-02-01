package main

import (
  "fmt"
)

func main() {
  var num int = 5
  fmt.Println(num)

  var num2 int = increment_by_value(num)
  fmt.Println(num,num2) // NO change in value of num but new value is in new variable num2

  increment_by_ref(&num); // passing value by reference

  fmt.Println("value is changed by ref only current num is :",num)

}

func increment_by_value(x int) int {
  x++
  return x
}

func increment_by_ref(x *int)  {
  // x will give the address of given value a * will use for de-referencing
  *x++ // getting x value and incrimenting that
  // there is no need to return any thing the variable is changed at its place only
}
