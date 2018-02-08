package main

import (
	"fmt"
)

type Employee struct {
	name string
	salary int
	currency string
}

type Student struct {
	name string
	salary int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s %d", e.name, e.currency, e.salary)
}

func main() {
	emp := Employee {
		name: "Surjit Singh",
		salary: 14500,
		currency: "$",
	}
	
	// student := Student {
	// 	name: "Manjit Singh",
	// 	salary: 1450,
	// 	currency: "$",
	// }

	// student.displaySalary()
	emp.displaySalary()
}
