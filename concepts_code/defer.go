// defer is a key word which is used before to any line of code to
// it uses before any line of statement to make it delay to end of function

package  main

import(
  "fmt"
)

func main()  {
  // for follwing order should be start middle end
  fmt.Println("start ")
  defer fmt.Println("middle")
  fmt.Println("end")
  // due to defer tag it is  start end then middle 

}
