package main

import(
  "fmt"
)

func main()  {
  a:= 10 // 1010
  b:= 3  // 0011
  fmt.Println(a + b) // addition
  fmt.Println(a - b) // substraction
  fmt.Println(a * b) // multiplication
  fmt.Println(a / b) // division  Note: int divesion will return only int not a float so ans will be 3 not 3.33333
  fmt.Println(a & b) // and oprator
  fmt.Println(a | b) // or oprator
  fmt.Println(a ^ b) // exclusive Or oprator

  // Bit shifting

  x:=8 // 2^3
  fmt.Println(x << 3) // bit shift left 3 places 2^3 * 2^3 = 2^6 = 64
  fmt.Println(a >> 3) // bit shift right 3 places 2^3 / 2^3 = 2^0 = 1

}
