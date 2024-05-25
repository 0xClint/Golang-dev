package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://omkardarde.vercel.app/"

func main() {
	fmt.Println("WEb srequests in GOlang")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Println("Response = ", response.Body)
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println("Response.Body = ", string(databytes))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
