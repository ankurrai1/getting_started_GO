// basic example and concept buildig for switch expression
package main

import (
  "fmt"
)

func main()  {
  num := 3
  switch num {
  case 1:
    fmt.Println("num is one")
  case 2:
    fmt.Println("num is two")
  case 3, 4, 5:                  // here we can pass as many as we want values if they belong to same group
    fmt.Println("num is three")
  default:      // default case if no case will run it will fire
    fmt.Println("num is not 1 2 3 4 5")
  }

}
