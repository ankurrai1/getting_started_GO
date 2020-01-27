package main

import "fmt"

func main() {
	var w int
	var h int
	fmt.Print("enter hight:")

	fmt.Scanln(&w)
	fmt.Print("enter width:")
	fmt.Scanln(&h)
	fmt.Println("area of rectangle:--", (w * h))
	fmt.Println("area of rectangle:--", (w + h + w + h))
}
