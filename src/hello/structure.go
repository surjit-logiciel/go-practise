package main

import (
	"fmt"
)

type Employee struct {
	firstName,  lastName string
	age, salary int
}

func main() {
	emp1 := Employee {
		firstName : "Sam",
		age: 25,
		salary: 500,
		lastName: "Anderson",
	}

	fmt.Println("Employee 1", emp1)

	emp2 := Employee{
		"singh", "surjit", 29, 800,
	}
	fmt.Println("Employee 2", emp2)

	emp3 := struct {
		firstName, lastName string
		age, salary int
	} {
		firstName: "surjit",
		lastName: "singh",
		age : 18,
		salary: 21000,
	}
	fmt.Println("Employee 3", emp3)

	var emp4 Employee
	fmt.Println("Employee 4", emp4)

	emp5 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name", (*emp5).firstName)
}