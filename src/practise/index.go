package main

import (
	"fmt"
)

func main() {
	// ifCond()
	// loop()
	// forLoop()
	// switchCase()
	// array()
	// array1()
	// array2()
	// array3()
	// array4()
}

func array4() {
	a := [...]float64{14.05, 125.30, 145.00}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f \n", i, v)
		sum += v
	}
	fmt.Println("\n sum of all element  of a ", sum)
}
func array3() {
	// a := [5]int{5, 6, 7}
	// var b [5]int
	// b = a
}

func array2() {
	a := [...]int{12, 13, 14, 15, 16}
	fmt.Println(a)
}

func array1() {
	a := [3]int{3}
	fmt.Println(a)
}
func array() {
	var a [4]int
	a[0] = 12
	a[1] = 13
	a[2] = 14
	a[3] = 140

	fmt.Println(a)
}


func switchCase() {
	finder := 110
	switch finder {
	case 1: 
		fmt.Println("1");
	case 10:
		fmt.Println("10")
	}
}

func forLoop() {

	// for no, i:= 10,1;  i > 10 && no 

	// i := 1
	// for ;i < 10; {
	// 	fmt.Printf("%d", i)
	// 	i += 2
	// }
	// fmt.Printf("test")
}

func loop() {
	for i := 1; i <= 10; i++ {

		if(i == 5) {
			continue
		}
		fmt.Printf("%d", i)
		fmt.Printf("\n")
	}
	fmt.Printf("Loop End!")
}

func ifCond() {
	num :=10
	if num % 2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
