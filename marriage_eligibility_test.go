//Ladder if-else
//Marriage Eligibility Test
package main

import "fmt"

func main() {
	var name, gender string
	var age int
	fmt.Println("Enter your Name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your Age:")
	fmt.Scanln(&age)
	fmt.Println("Enter your Gender:")
	fmt.Scanln(&gender)
	if gender == "m" && age >= 21 {
		fmt.Printf("Hello! %s Congratulations! You are eligible for marriage", name)
	} else if gender == "f" && age >= 18 {
		fmt.Printf("Hello! %s Congratulations! You are eligible for marriage", name)
	} else {
		fmt.Printf("Hello! %s Sorry! You are not eligible for marriage", name)
	}
}
