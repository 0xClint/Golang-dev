package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to Golang!"
	fmt.Println(welcome)

	reader := bufio.NewReader((os.Stdin))

	fmt.Println("Enter the your rating to GOLang")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating!, ", input)
	fmt.Printf("Type of input : %T", input)

}
