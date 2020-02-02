package main
import "fmt"
func main(){
	var a int 
	var b int
	fmt.Print("enter a:")
	
	fmt.Scanln(&a)
	fmt.Print("enter b:")
	fmt.Scanln(&b)
	fmt.Println("add:--",(a+b))
	fmt.Println("sub:--",(a-b))
	fmt.Println("mul:--",(a*b))
	fmt.Println("div:--",(a/b))
}