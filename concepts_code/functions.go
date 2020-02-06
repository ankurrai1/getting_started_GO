package main

import (
  "fmt"
  "errors"
  "math"
)

func main(){
  res := sum(5,7)
  fmt.Println("sum of no :" ,res, "\n")

  result,err := sqrt(36)

  fmt.Println(result,err);

  result2,error := sqrt(-35)

  fmt.Println(result2,error)

  sum("hello every one",1,2,3,4,5,6,7)// it can take more as well and rape them into slices to use
  // msg = "hello every one"
  // values = 1,2,3,4,5,6,7  these are called as "variadic parameter"
}


//defining a new function Sum instade of writing all code in main

// basic structure is this

// func sum(a int, b int)  {
//
// }

// intermediate Structure of function in go
// With return type

func sum(a int, b int) int  {
  return a + b
}

// Advance Structure of function in go
// go language can return multipal values that is an advance level of go


func sqrt(a float64) (float64,error)  {
  if a < 0 {
    return 0,errors.New("Not defined for negative values")
  }
  return math.Sqrt(a), nil
}


//passing multipel values

func sum(msg string,value ...int)  { // this must be last and one argument in function
  fmt.Println(msg,values) // it will create it as Slices
}
