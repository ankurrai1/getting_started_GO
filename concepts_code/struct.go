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

  //structs  are passed by values
  // you can use & to pass it by referance
  
  fmt.Println(p)
  fmt.Println(p.age)

  // go will find the way to assign value if given in right orders
  // without keys name as well
  p2 := person{
    "sagar",
    "garg",
    20,
  }

  fmt.Println(p2.age)



  per := struct{name string; age int}{name:"sagar",age:21}
  fmt.Println(per)
}
