package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in GOlang")
	content := "This is ClintOP!"
	file, err := os.Create("./temp.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is :", length)
	defer file.Close()

	readFile("./temp.txt")

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("content inside the file :\n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
