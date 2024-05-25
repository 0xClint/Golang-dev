package main

import "fmt"

// DEFER => runs the piece of code at the end of func, if multiple defers than it take place in LIFO manner

func main() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println("Namaste World!")
	myDefer()
}

func myDefer() {
	for i := 0; i <= 4; i++ {
		defer fmt.Println(i)
	}
}
