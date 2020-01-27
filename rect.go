package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("enter hight:")

	fmt.Scanln(&a)
	fmt.Print("enter width:")
	fmt.Scanln(&b)
	fmt.Println("area of rectangle:--", (a * b))
	fmt.Println("area of rectangle:--", (a + b + a + b))
}
