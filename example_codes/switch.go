package main

import "fmt"

func main() {
	var input int
	fmt.Println("enter 1 for add two num")
	fmt.Println("enter 2 for substract two num")
	fmt.Println("enter 3 for multiply two num")
	fmt.Println("enter 4 for devision ")
	fmt.Println("press (1 or 2 or 3 or 4)")
	fmt.Scanln(&input)
	switch input {
	case 1:
		{
			var a, b int
			fmt.Println("enter a:")
			fmt.Scanln(&a)
			fmt.Println("enter b:")
			fmt.Scanln(&b)
			fmt.Println(a + b)
		}
	case 2:
		{
			var a, b int
			fmt.Println("enter a:")
			fmt.Scanln(&a)
			fmt.Println("enter b:")
			fmt.Scanln(&b)
			fmt.Println(a - b)
		}
	case 3:
		{
			var a, b int
			fmt.Println("enter a:")
			fmt.Scanln(&a)
			fmt.Println("enter b:")
			fmt.Scanln(&b)
			fmt.Println(a * b)
		}
	case 4:
		{
			var a, b int
			fmt.Println("enter a:")
			fmt.Scanln(&a)
			fmt.Println("enter b:")
			fmt.Scanln(&b)
			fmt.Println(a / b)

		}
	default:
		{
			fmt.Println("wrong input")
		}
	}
}
