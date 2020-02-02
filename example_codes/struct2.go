package main

import "fmt"

type student struct {
	rollno int
	name   string
	marks  int
}

func main() {
	var s student
	s.rollno = 188607
	s.name = "prem mishra"
	s.marks = 991 
	fmt.Println(s)
}
