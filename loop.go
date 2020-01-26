package main

import(
  "fmt"
)

func main()  {
  // Array init with 0 if nothing is given
  var arr [5]int
  fmt.Println(arr)

  // getting and changing the value of Array
  arr[3] = 5
  fmt.Println("array after change : ",a,"\n");

  // Second methed of array init With some initial value
  a := [5]int{3,4,5,6,7}
  fmt.Println("Some other array with some initial values : ",a,"\n Dynaminc arrays \n")

  //An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
  //flexible view into the elements of an array. In practice, slices are much more common than arrays.
  primes := [6]int{2, 3, 5, 7, 11, 13} // Fixed sized

	var s []int = primes[1:4] // init without any size and it is a dynamic One
	fmt.Println(s)
  s.append(22)
  s.append(44)
  fmt.Println("after append 22 and 44 to dynaminc array s : ",s)

}
