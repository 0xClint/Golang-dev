package main

import "fmt"

const Token = "gdkndlbntibn" //Public Variable

func main() {
	var username string = "Clint"
	fmt.Println(username)
	fmt.Printf("Type of username : %T \n", username)

	var username1 string
	fmt.Println(username1)

	var isAdult bool = true
	fmt.Println(isAdult)
	fmt.Printf("Type of isAdult : %T \n", isAdult)

	var isAdult1 bool = true
	fmt.Println("undeclared value of bool : ", isAdult1)

	var smallUint uint8 = 255
	fmt.Println(smallUint)
	fmt.Printf("Type of smallUint : %T \n", smallUint)

	var Integer int
	fmt.Println("undeclared value of Integer : ", Integer)

	var smallFloat float32 = 32.477027859673950683
	fmt.Println(smallFloat)
	fmt.Printf("Type of smallFloat : %T \n", smallFloat)

	var bigFloat float64 = 32.477027859673950683
	fmt.Println(bigFloat)
	fmt.Printf("Type of bigFloat : %T \n", bigFloat)

	// Implicit type
	var website = "omkardarde.vercel.app"
	fmt.Println("website link :", website)
	website = `omkardarde.me`
	fmt.Println("website link :", website)

	age := 21
	fmt.Println("Age :", age)

	fmt.Println("token :", Token)

}
