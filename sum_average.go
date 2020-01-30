//WAP to find Sum & Average of ten number using array
package main

import "fmt"

func main() {
	var x [10]int
	var i, s int
	var avg float64
	s = 0
	fmt.Println("Enter ten Numbers: ")
	for i = 0; i < 10; i++ {
		fmt.Scanln(&x[i])
		s = s + x[i]
	}
	avg = float64(s) / 10
	fmt.Printf("Sum= %d\n", s)
	fmt.Printf("Average= %.2f", avg)

}
