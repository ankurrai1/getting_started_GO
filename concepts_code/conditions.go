package main

import (
	"fmt"
)

func main() {
	x := 4

	// following is showing nested if statement
	// NOTE a: there is no () around condition
	if x > 6 {
		if x == 7 {
			fmt.Println("X is 7 \n")
		}
	} else if x == 5 {
		fmt.Println("X is 5")
	} else {
		fmt.Println("X is less then 5")
	}
}
