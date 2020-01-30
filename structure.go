package main

import "fmt"

type employee struct {
	empid   int
	empname string
	salary  int
}

func main() {
	var emp [3]employee
	var i int
	fmt.Println("enter the details of employee")
	for i = 0; i < 3; i++ {
		fmt.Print("enter employee id :")
		fmt.Scan(&emp[i].empid)
		fmt.Println("enter employee name:")
		fmt.Scan(&emp[i].empname)
		fmt.Println("enter employee salary:")
		fmt.Scan(&emp[i].salary)

	}
	fmt.Println("details of employee:")
	for i = 0; i < 3; i++ {
		fmt.Println(emp[i])

	}
}
