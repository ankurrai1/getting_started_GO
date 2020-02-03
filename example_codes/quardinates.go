package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("enter the value of x : ")
	fmt.Scanln(&x)
	fmt.Print("enter the value of y : ")
	fmt.Scanln(&y)
	if x > 0 && y > 0 {
		fmt.Print("point is in first quadrant")
	} else if x < 0 && y > 0 {
		fmt.Print("point is in second quadrant")
	} else if x < 0 && y < 0 {
		fmt.Print("point is in third quadrant")
	} else if x > 0 && y < 0 {
		fmt.Print("point is in fourth quadrant")
	} else {
		fmt.Print("point is mean or on a exis")
	}
}
