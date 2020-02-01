// shadowing in go variable
// it meanse only local variable will use if there is two variable with same name
// in package level and in local block level

package main

// No , in import if importing more then one package
import (
  "fmt"
  "strconv"
)

//  If somthing starting with capital letter in its name it means it is Exported for other packages as well
// for example
var Status bool = true
// so the above variable is available every where

var name string = "Ramu kaka"
var age string = "54"

func main()  {
  fmt.Printf("age is : %s  %T \n",age,age);
  // the following variable will shadow down the global variable
  var age int = 57
  fmt.Printf("age is : %d  %T \n",age,age);

  // variable coversion form one type to another
  var someNo int = 36;
  fmt.Printf("type of variable : %T %d\n",someNo,someNo)

  // creating new variable in float and casting someNo variable to that
  var newNo float64
  newNo = float64(someNo)

  fmt.Printf("type of variable : %T  %f\n",newNo,newNo)

  // to convert something to string you have to import "strconv" and use its Itoa function as following
  var x int = 59
  fmt.Printf("x is : %d  %T \n",x,x);
  var newX string = strconv.Itoa(x)
  fmt.Printf("now the same above variable is of type is : %T x = %v \n",newX,newX)


  // complex type variable can be used in this language
   var someCom complex64 = 1 + 2i
   fmt.Printf("%v and its type is : %T",someCom,someCom)

}
