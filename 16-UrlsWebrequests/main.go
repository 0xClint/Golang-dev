package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const Url string = "http://localhost:8000"

func main() {
	fmt.Println("web url reuests GOLang")
	// performGetRequest(Url)
	// performPostRequest(Url + "/post")
	performPostFormRequest(Url + "/postform")

}

func performGetRequest(url string) {

	response, err := http.Get(url)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Response length : ", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	checkNilErr(err)

	// fmt.Println("Get request response :\n", string(content))

	var responseString strings.Builder // Another way to print raw content data
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount of data : ", byteCount)
	fmt.Println("Get request response :\n", responseString.String())

}

func performPostRequest(url string) {

	requestBody := strings.NewReader(`{
		"name":"ClintOP",
		"age":21,
		"website":"https://omkardarde.vercel.app/"
	}`)
	response, err := http.Post(url, "application/json", requestBody)

	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)

	content, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println("Get request response :\n", string(content))
}

func performPostFormRequest(urlstring string) {
	data := url.Values{}
	data.Add("firstName", "John")
	data.Add("lastName", "Wick")
	data.Add("age", "31")

	response, err := http.PostForm(urlstring, data)
	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)

	content, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println("Get request response :\n", string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
