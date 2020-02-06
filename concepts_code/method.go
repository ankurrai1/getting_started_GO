// Go supports methods defined on struct as methods defined in any object orianded language
package main

import(
  "fmt"
)

type rect struct {
    width int
    height int
}

// This area method has referance of rect struct.

func (r *rect) area() int { // pass by referance
    return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// following an example of a value receiver.

func (r rect) perim() int { // pass by value
    return 2*r.width + 2*r.height
}

func main() {

  r := rect{width: 10, height: 5} // here we are creating the stucture

// Here we call the 2 methods defined for our struct.

    fmt.Println("area: ", r.area())
    fmt.Println("perimiter:", r.perim())

// Go automatically handles conversion between values and pointers for method calls.
// You may want to use a pointer receiver type to avoid copying on method calls or
// to allow the method to mutate the receiving struct.
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perimiter:", rp.perim())
}
