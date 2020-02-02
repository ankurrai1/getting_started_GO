package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("enter a number : ")
	fmt.Scanln(&a)
	b = rev(a)
	fmt.Print("reverse is : ", b)
}
func rev(x int) int {
	var b, d int
	for d = x; d > 0; d = d / 10 {
		b = (b + (d % 10)) * 10
	}
	b = b / 10
	return (b)
}
