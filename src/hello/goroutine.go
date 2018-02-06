package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("Hello world goroutine.")
}

func numbers() {
	for i:=1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}
func main() {
	go hello()
	// time.Sleep(1 * time.Second)
	fmt.Println("Main Function")
}