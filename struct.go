package main

import (
  "fmt"
)


type person struct {
  first_name string
  last_name string
  age int
}

func main()  {
  p := person{first_name:"Ramu",last_name:"kaka",age:100}
  fmt.Println(p)

  fmt.Println(p.age)
}
