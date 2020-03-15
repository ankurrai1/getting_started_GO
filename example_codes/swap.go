package main
import "fmt"

func main() {
	var a, b, c int

	fmt.Println("enter a:")
	fmt.Scanln(&a)
	fmt.Println("enter b:")
	fmt.Scanln(&b)
	c = b
	b = a
	a = c
	fmt.Printf("After Swapping a Is:%d\n", a)
	fmt.Printf("After Swapping b Is:%d", b)
}
