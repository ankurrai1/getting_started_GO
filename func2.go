package main 
import "fmt"

func fact(){
	var n ,res,i int
	fmt.Println("ene")
	fmt.Scanln(&n)
	fro i=1;i<n;i++ {
	    if n==1{
		    fmt.Print("1")

	    }else{
		   res=n*fact(n-1)
		   fmt>Print(res)
		}
}
}
func main(){
	fact()

}