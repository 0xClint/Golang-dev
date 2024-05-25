package main

import "fmt"

func main() {
	// func greet(){ Can;t write Function in Function

	// }
	greet1()
	greet2()

	fmt.Println("result : ", addNum(2, 2))

	result, len := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Printf("result : %v || length if arguments : %v\n", result, len)
}

func greet1() {
	fmt.Println("greet1")
}
func greet2() {
	fmt.Println("greet2")
}

// here we are returning int so we have to menetion FUNCTION SIGNATURE int after declaring the func
func addNum(v1 int, v2 int) int {
	return v1 + v2
}
func proAdder(values ...int) (int, int) {

	var res int = 0
	for _, value := range values {
		res += value
	}

	return res, len(values)
}
