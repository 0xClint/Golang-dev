package main

import "fmt"

func main() {
	var ptr1 *int
	fmt.Println("ptr1 value : ", ptr1)

	myNumber := 45
	var ptr2 = &myNumber
	fmt.Println("ptr2 : ", ptr2)
	fmt.Println("*ptr2 : ", *ptr2)
	*ptr2 *= 2
	fmt.Println("*ptr2 : ", *ptr2)

}
