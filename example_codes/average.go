package main

import ( 
	"fmt"
       )

func main() {
	var x [100]int
	var i, s, n int
	var avg float64
	s = 0
	fmt.Print("Enter the number of element you want to calculate average for : ")
	fmt.Scanln(&n)
	fmt.Print("enter ", n, " numbers in array\n")
	for i = 0; i < n; i++ {
		fmt.Scanln(&x[i])
		s = s + x[i]
	}

	avg = float64(s) / float64(n)
	fmt.Print("sum is = ", s, "\navg is = ", avg)
}
