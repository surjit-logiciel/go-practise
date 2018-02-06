package main

import (
	"fmt"
	"math"
)

func main() {

	a,b := 145.8, 145.9

	c := math.Max(a, b)

	fmt.Println("minimum number is", c);

	// a, b := 20, 30
	// fmt.Println("a is", a, "b is", b)
	// a,b := 40,50

	// name, age := "Naveen", 24 //error
	// fmt.Println("My name is",	 name, "my age is", age)
	// var (
	// 	name = "Surjit Singh"
	// 	age  = 20
	// 	height int
	// )

	// fmt.Println("My name is:", name, "my age is:", age,"height is:", height)
	// var age = 26
	// age = 25
	// fmt.Println("My age is", age);
	// fmt.Println("Hellow world!")
	 // var width, height int = 20, 100

	 // fmt.Println("width id", width, "height is", height)
	 
	 //short hand declartion
	 // name, age := "Naveen", 29

	 // fmt.Println("My name is", name, "age is", age)
}

// C:\Projects\Go\src\hello\helloworld.go