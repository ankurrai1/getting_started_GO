package main

import (
	"fmt"
	"time"
)

func main() {
	var i, a int
	fmt.Println("enter a number:")
	fmt.Scanln(&a)
	for i = a; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
