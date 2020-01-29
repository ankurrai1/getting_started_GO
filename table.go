package main

import "fmt"

func main() {
	var i, table, n int
	fmt.Print("enter a number:")
	fmt.Scanln(&n)
	for i = 1; i <= 10; i++ {

		table = i * n
		fmt.Printf("%d*%d=%d\n", n, i, table)
	}
}
