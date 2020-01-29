package main

import(
  "fmt"
)

const a = 34
const b int = 35
const flag bool = true
const str string  = "ram"

func main() {

  const a int32 = 400 // shadowing it will overright the above global form local variable

  fmt.Printf("%v  %T",a,a)
  fmt.Printf("%v  %T",b,b)
  fmt.Printf("%v  %T",flag,flag)
  fmt.Printf("%v  %T",str,str)
}
