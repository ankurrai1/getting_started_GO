package main

import(
  "fmt"
)

// in go arrays are collection of same data type
// they are of fixed size at compile time
// so you can not take user input and then define a array
// you have to use someting called slices


func main()  {
  // declaration of arrays are of these three types
  a := [3]int{1,2,3} // simple method
  b := [...]{1,2,3}  // using spread oprator
  var c [3]int  // without init
  var d [3]int
  fmt.Println(len(a))

  // Importent : assigned array is a copy of previous one not the referance of that
  c = a;
  fmt.Println(a,c)
  c[2] = 0
  fmt.Println(a,c) // here you can see a is not changed but c is changed so it is a copy of previous array
  d[2] = 8
  fmt.Println(a,d) // passed by referance so value is changed
  // the above thing cause extra memory space

  // SLICES : they are basically dynamic array and they use array under;
  // they can also be declared of three types

  p := [...]{1,2,3,4,5,6,7,8,9,0,11,12,13}
  // 1 is by directly slicing  the array
  q := [:] //sliceing full array
  r := [3:6] // starting index is three and upto 6 excluding 3rd index and including 6th

  fmt.Println(p," ",q," ",r," ")
  fmt.Println("lenght is : ",len(q),"capacity is : ",cap(q))

  // by declaration
  s := []int{1,2,3,4,5}
  // by make function
  t := make([]int, 3) // type of data to make and second is lenght and it will set as capacity as well
  u := make([]int,5,10) // here length is 5 but capacity is set to 10

  fmt.Println("for array t :",len(t),cap(t))
  fmt.Println("for array u :",len(u),cap(u))

  // NOTE : slices are just send by referance

}
