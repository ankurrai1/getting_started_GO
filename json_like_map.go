package main

import (
  "fmt"
)

func main(){
  //basically a map is a dic of python object in javascript or
  // in more simple words it is a key value pair of data

  edges := make(map[string]int) // here making map object where key type is string and value is int
  fmt.Println("a map :", edges ,"\n")

  edges["side"] = 5
  edges["len"] = 10
  edges["height"] = 12
  edges["thickness"] = 2

  fmt.Println("map after some data addition : ",edges,"\n")

  fmt.Println("Printing some specific values : ", edges["side"])

  // deleting some thing form map by key name
  delete(edges,"thickness");

  fmt.Println("map after deleting the thickness : ",edges,"\n")
}
