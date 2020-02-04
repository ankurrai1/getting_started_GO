// defer is a key word which is used before to any line of code to
// it uses before any line of statement to make it delay to end of function

package  main

import(
  "fmt"
)

// NOTE :- just comment following blocks of code to see other block working 

func main()  {

  //========================================================
  // for follwing order should be start middle end
  fmt.Println("start ")
  defer fmt.Println("middle")
  fmt.Println("end")
  // due to defer tag it is  start end then middle

  // =========================================================

  // example will show the last statement which is deferred will execute first
  // for to use case impliment to rapping something and unrapping that
  // closing opened connection in same order

  defer fmt.Println("first ")
  defer fmt.Println("Second")
  defer fmt.Println("last")

  //==========================================================

 // defer will store the value at the time when it is wiritten not the value when it is executed
  a := 10
  defer fmt.Println(a)
  fmt.Println("defer will print after me")
  a = 100

 // in above  example it will store value in defer lile and it will print that
 // it whether it will ececute after value changed it will cached first value


}
