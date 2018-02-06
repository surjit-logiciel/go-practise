package main

import (
	"fmt"
	"geometry/rectangle" //importing custom package
)

func main() {
	var rectLen, rectWidth float64 = 6,7
	fmt.Println("Geomatrical shape properties")
	fmt.Printf("Area of rectangles %.2f\n", rectangle.Area(rectLen, rectWidth))
	
	fmt.Printf("diagonal of the rectangle %2.f", rectangle.Diagonal(rectLen, rectWidth))
	// fmt.Println("Geomatrical shape properities")
}
// go run C:/Projects/Go/src/geometry/geometry.go
// https://golangbot.com/packages/