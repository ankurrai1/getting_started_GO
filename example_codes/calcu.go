package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Print("choose your option\n1 for addition\n2 for substraction\n3 for multiplication\n4 for devide\n")
	fmt.Scanln(&d)
	fmt.Print("enter two numbers\n")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	switch d {
	case 1:
		c = add(a, b)
		fmt.Print("ans is : ", c)
	case 2:
		c = sub(a, b)
		fmt.Print("ans is : ", c)
	case 3:
		c = mul(a, b)
		fmt.Print("ans is : ", c)
	case 4:
		c = dev(a, b)
		fmt.Print("ans is : ", c)
	}
}

func add(x, y float64) float64 {
	var a float64
	a = x + y
	return (a)
}

func sub(x, y float64) float64 {
	var a float64
	a = x - y
	return (a)
}

func mul(x, y float64) float64 {
	var a float64
	a = x * y
	return (a)
}

func div(x, y float64) float64 {
	var a float64
	a = x / y
	return (a)
}
