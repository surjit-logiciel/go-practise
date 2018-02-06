package main

import (
	"fmt"
)

func multiply(a, b int) int {
	var total = a * b

	return total
}

func main(){
	// a, b := 10, 5
	// total := multiply(a, b)
	// fmt.Println(total)
	
	area, _ := rectProps(10.02, 0.05)
	fmt.Printf("Area %f", area)
}

func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = 2.05

	return area, perimeter
}

