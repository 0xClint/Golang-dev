package main

import "fmt"

func main() {
	fmt.Println("Let's study map")

	var languages = make(map[string]string)

	languages["js"] = "Javascript"
	languages["py"] = "Python"
	languages["rb"] = "Rubby"
	languages["go"] = "GOLang"
	fmt.Println(languages)

	for key, values := range languages {
		fmt.Printf("%v => %v \n", key, values)
	}
	delete(languages, "rb")

	fmt.Printf("After deletion of rb \n")

	for key, values := range languages {
		fmt.Printf("%v => %v \n", key, values)
	}
}
