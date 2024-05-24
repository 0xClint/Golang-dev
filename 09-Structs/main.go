package main

import "fmt"

func main() {
	fmt.Println("Let's study Structs")

	clint := User{"Clint", "clint@abc.xyz", false, 21}
	fmt.Printf("Clint details : %+v\n", clint) //%+v prints whole struct with key and properties
	fmt.Printf("Clint => Name: %v \nEmail: %v\nAdult: %v\nAge: %v", clint.Name, clint.Email, clint.Adult, clint.Age)
}

type User struct {
	Name  string //Capital letter means public
	Email string
	Adult bool
	Age   int
}
